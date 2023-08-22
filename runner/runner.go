package runner

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

// ------------------------- ServiceQ -------------------------
type Runner struct {
	id                 string         // id of the runner. md5(epoch+sysCmd)
	stream             bool           // if true, then streams out the stdout's content through an unbuffered output chanel. default=false
	loopCount          int            // 1 => run once; n(n>1) => run n times; 0 => loop forever. default=1
	iterationStatus    map[int]string // records iteration number and the corrosponding state. e.g. [1] => TIMEDOUT, [2] => SUCCESSFUL [3] => SUCCESSFUL
	timeout            int            // maximum duration (in seconds) the system call has before declearing failed. default=0
	waitingPeriod      int            // number of seconds the runner waits before executing the system call. default=0
	restingPeriod      int            // number of seconds the runner waits before the next iteration. only when in loopForever=true. default=0
	scheduledAt        int64          // future time when the command is scheduled to execute. default=0
	filePath           string         // specifiy a file to write logs. if  set to "" then no log files will be written. default=""
	execuitedAt        uint64         // epoch then the system call was execuited
	executionTimeNano  uint64         // units of nano seconds ellapsed since the system call has started. default=0
	verificationPhrase string         // if this string is found the in the log then that means the output is vefified, and now the status becomes SUCCESSFUL. default="" that means verification is disabled and now the status can become FINISHED
	status             string         // PENDING | RUNNING | SUCCESSFUL | CRASHED | FINISHED | TIMEDOUT. default=PENDING
	sysCmd             string         // the system command to be execuited. e.g. "ls -al | grep foo". Required
	log                string         // stdout of the system call is kept here. default="", max=1MB
	mu                 sync.RWMutex   // for making resources thread safe
}

// ----------------------- instantiation -----------------------
func NewRunner(cmd string) *Runner {
	// TODO take a second optional string parameter. if that is set to "RUN_ONCE" then this function's behaviour will be modified and

	// initialize a runner with default settings and a system command
	r := &Runner{}
	r.id = calculateMD5([]byte(fmt.Sprintf("%v", time.Now().UnixNano()) + cmd)[:]) // md5 (current_epoch + name)
	r.sysCmd = cmd
	r.loopCount = 1
	r.status = "PENDING"

	return r
}

// if enabled, stdout is piped out using `stdout` channel
func (r *Runner) EnableStream() {
	r.stream = true
}

func (r *Runner) DisableStream() {
	r.stream = false
}

// set number of times the the shell command will run
func (r *Runner) SetLoopCount(count int) {
	if count < 0 {
		return
	}
	r.loopCount = count
}

func (r *Runner) SetTimeout(timeout int) {
	r.timeout = timeout
}

func (r *Runner) SetWaitingPeriod(wp int) {
	r.waitingPeriod = wp
}

func (r *Runner) SetRestingPeriod(rp int) {
	r.restingPeriod = rp
}

func (r *Runner) Schedule(sc int64) {
	r.scheduledAt = sc
}

func (r *Runner) SetLogPath(log string) {
	r.log = log
}

func (r *Runner) SetVerificationPhrase(phrase string) {
	r.verificationPhrase = phrase
}

func (r *Runner) GetState() map[string]interface{} {
	var logSize = len(r.log)
	return map[string]interface{}{
		"id":                   r.id,
		"stream":               r.stream,
		"loop_count":           r.loopCount,
		"iterationwise_status": r.iterationStatus,
		"timeout":              r.timeout,
		"waiting_period":       r.waitingPeriod,
		"resting_period":       r.restingPeriod,
		"scheduled_at":         r.scheduledAt,
		"file_path":            r.filePath,
		"execuited_at":         r.execuitedAt,
		"execution_time_nano":  r.executionTimeNano,
		"verification_phrase":  r.verificationPhrase,
		"status":               r.status,
		"system_command":       r.sysCmd,
		"log_size_bytes":       logSize,
	}
}

func (r *Runner) GetStatus() string {
	return r.status
}

func (r *Runner) Execute() (string, error) {
	fmt.Println("Here")
	cmd := exec.Command("bash", "-c", r.sysCmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil

}

func (r *Runner) ExecutePayload(command string) (string, error) {
	/*
		this function accounts for the fact that an input command could also be "ls -al | grep jpg && pwd"
	*/
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (r *Runner) Logs() string {
	return r.log
}

func (r *Runner) ClearLogs() {

	r.log = ""
}

func (r *Runner) Restart() {
	// code
}

func (r *Runner) Stop() {
	// TODO: Need to plan out
}

// ----------------------- utility ---------------------------------
func calculateMD5(data interface{}) string {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	hash := md5.Sum(buffer.Bytes())
	return hex.EncodeToString(hash[:])
}
