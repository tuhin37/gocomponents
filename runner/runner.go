package runner

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var STDPIPE = make(chan []byte)

// ------------------------- ServiceQ -------------------------
type Runner struct {
	id                 string // id of the runner. md5(epoch+sysCmd)
	isConsole          bool   // if set to true then the shell's commands output is printed on console. default=false
	timeout            int    // maximum duration (in seconds) the system call has before declearing failed. default=0
	waitingPeriod      int    // number of seconds the runner waits before executing the system call. default=0
	logFile            string // specifiy a file to write logs. if  set to "" then no log files will be written. default=""
	execuitedAt        int64  // epoch then the system call was execuited
	executionTimeNano  int64  // units of nano seconds ellapsed since the system call has started. default=0
	verificationPhrase string // if this string is found the in the log then that means the output is vefified, and now the status becomes SUCCESSFUL. default="" that means verification is disabled and now the status can become FINISHED
	status             string // PENDING | RUNNING | COMPLETED | FAILED | TIMEDOUT | SUCCEEDED. default=PENDING. [default=PENDING, COMPLETED => the exitcode of the system call was 0, FAILED => non-zero, SUCCESSFUL => the verification string was found]
	sysCmd             string // the system command to be execuited. e.g. "ls -al | grep foo". Required
	logBuffer          []byte // stdout of the system call is kept here. default="", max=1MB
}

// ----------------------- constructor -----------------------
func NewRunner(cmd string) *Runner {
	// TODO take a second optional string parameter. if that is set to "RUN_ONCE" then this function's behaviour will be modified and

	// initialize a runner with default settings and a system command
	r := &Runner{}
	r.id = calculateMD5([]byte(fmt.Sprintf("%v", time.Now().UnixNano()) + cmd)[:]) // md5 (current_epoch + name)
	r.sysCmd = cmd
	r.status = "PENDING"

	return r
}

// enable console print
func (r *Runner) EnableConsole() {
	r.isConsole = true
}

func (r *Runner) DisableConsole() {
	r.isConsole = false
}

func (r *Runner) SetTimeout(timeout int) {
	r.timeout = timeout
}

// if set then the rinner will wait that much seconds before execuiting the system call. default=0
func (r *Runner) SetWaitingPeriod(wp int) {
	if wp > 0 {
		r.waitingPeriod = wp
	}
}

// provide a filename for log
func (r *Runner) SetLogFile(file string) {
	r.logFile = file
}

// if verification phrase is provided, then the runner finds the phrase in the output of the system call
// if found then
func (r *Runner) SetVerificationPhrase(phrase string) {
	r.verificationPhrase = phrase
}

func (r *Runner) GetState() map[string]interface{} {
	var logSize = len(r.logBuffer)
	return map[string]interface{}{
		"id":                  r.id,
		"console":             r.isConsole,
		"timeout":             r.timeout,
		"waiting_period":      r.waitingPeriod,
		"file_path":           r.logFile,
		"execuited_at":        r.execuitedAt,
		"execution_time_nano": r.executionTimeNano,
		"verification_phrase": r.verificationPhrase,
		"status":              r.status,
		"system_command":      r.sysCmd,
		"log_size_bytes":      logSize,
	}
}

func (r *Runner) GetStatus() string {
	return r.status
}

func (r *Runner) Execute(commands ...string) ([]byte, error) {
	// Combine multiple commands into a single string if provided
	var command string
	if len(commands) > 0 {
		command = strings.Join(commands, " && ")
	} else {
		command = r.sysCmd
	}

	// initiate a command variable
	cmd := exec.Command("bash", "-c", command)
	r.sysCmd = command

	// set commands STDOUT to the STDOUT of the system
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return []byte{}, err
	}

	var file *os.File
	// Open the file in append mode
	if r.logFile != "" { // TODO check if a proper file or /path/to/file
		file, err = os.OpenFile(r.logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return []byte{}, err
		}
		defer file.Close()
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// instantiate a scanner object. will read command's output stream from here
	scanner := bufio.NewScanner(stdout)

	wg.Add(1)
	go func() {
		defer wg.Done()

		if r.logFile != "" {
			_, err := file.WriteString("[" + strconv.Itoa(int(time.Now().UnixMicro())) + "]" + " command" + "\n```shell\n" + r.sysCmd + "\n```\n\n" + "[" + strconv.Itoa(int(time.Now().UnixMicro())) + "]" + " Response begin\n```shell\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}

		for scanner.Scan() {
			stdoutLine := scanner.Text()

			r.logBuffer = append(r.logBuffer, scanner.Bytes()...)
			r.logBuffer = append(r.logBuffer, '\n')

			// write to a file here
			if r.logFile != "" {
				_, err := file.WriteString(stdoutLine + "\n")
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}

			// print to console
			if r.isConsole {
				fmt.Println(stdoutLine)
			}
		}

		// write footer after command is executed
		if r.logFile != "" {
			_, err := file.WriteString("```\n" + "[" + strconv.Itoa(int(time.Now().UnixMicro())) + "]" + " Response end\n" + "---\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}

	}()

	// Execute the command
	r.logBuffer = []byte{} // reset the buffer
	r.status = "RUNNING"
	r.execuitedAt = time.Now().UnixNano()
	if err := cmd.Start(); err != nil {
		return []byte{}, err
	}

	// timeoutNano := int64(r.timeout * 1000000000)
	// for {
	// 	r.executionTimeNano = time.Now().UnixNano() - r.execuitedAt
	// 	if r.executionTimeNano > timeoutNano {
	// 		// return prematurely
	// 		r.status = "TIMEDOUT"
	// 		return r.logBuffer, nil
	// 	}
	// }

	// Wait for the system call to finish
	err = cmd.Wait()

	// capture command exit timestamp

	// capture the exit code of the system call
	exitCode := cmd.ProcessState.ExitCode()

	// update status based on exit code
	if exitCode == 0 {
		r.status = "COMPLETED"
	} else {
		r.status = "FAILED"
	}

	// search and match output and set status to VERIFIED
	if r.verificationPhrase != "" && strings.Contains(string(r.logBuffer), r.verificationPhrase) {
		r.status = "SUCCEEDED"
	}

	// Wait for all goroutines to finish before returning
	wg.Wait()

	return r.logBuffer, nil
}

func (r *Runner) Logs() string {
	return string(r.logBuffer)
}

func (r *Runner) ClearLogs() {
	r.logBuffer = []byte{}
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
