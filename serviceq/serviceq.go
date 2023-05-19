package serviceq

/*

	DESCRIPTION:
		A service queue is similar to a restaurant queue where tasks represent customers waiting to be served.
		In this analogy, a worker corresponds to the restaurant staff serving the tasks (customers) in the queue.
		Multiple workers can serve a single queue. The worker selects a task from the queue and begins working on it.
		If the task fails, the worker can retry it for a specified number of times, which we refer to as the retry count.
		When a task is completed successfully, it is moved to the 'completed' queue, and the worker proceeds to the next task.
		If a task fails repeatedly, it will be moved to the 'failed' queue, and the worker will move on to the next task from the 'pending' queue.

	DETAILS:
		QUEUES:
			There are 4 queues in a service queue.
			1. 'pending'. This is the main queue where all the tasks are added.
			2. 'retry'. This is the queue where all the tasks are moved when they fail. The tasks are retried for a specified number of times.
			2. 'completed'. This is the queue where all the successfully completed tasks are moved.
			3. 'failed'. This is the queue where all the failed tasks are moved.
		TASKS:
			A task object in a service queue is a json object. It has the following general structure.
				{
					"task": {}							// actual task
					"batchID": "1234567890"				// batch id
					svcQName": "drag"					// service queue name
					"next_attempt_number": 3 				// this means the task has been already attempted 2 times and failed. Next retry will be the 3rd attempt.
					"next_attempt_scheduledstamp": 1234567890 	// the 3rd attempt is scheduled at this timestamp.
				}
		RETRY:
			A task is retried for a specified number of times. This is called the retry count.
			the next retry timestamp for nth attempt is calculated using the following formula.
			next_attempt_scheduledstamp = currentEpoch + restingPeriod + ((n-1)*loopBackoff). this calculation is carried out when (n-1)th attempt fails.
			restingPeriod and loopBackoff are configurable parameters with default value 0. This allows the user to configure the retry mechanism.
			One can configure the retry mechanism to have a fixed delay between every retry. This can be done by setting loopBackoff to 0.
			Also, it can be configured so that the weight time between each retries keeps increasing.

		WORKER:
			A service queue, can be configured to have one, more or no worker attached. Default = 1. If there are no qorkes attached,
			then a service queue behaves just like a regular queue. One can update the number of workers at any point in time. If no workers are working on the queue,
			then the queue status becomes 'PENDING'. There is a '.Task()' method in which one can pass in a function handler. This function is a
			custom definition of how to process each task. .Start() method manually starts the worker(s) if already not running.
			.Stop() method can prematurely and on demand can stop all the worker(s). In this case all the pending and the retrying takss will end up in the failed queue.
			The status will become 'COMPLETE'. .Pause() method can pause running workes. Status becomes 'PAUSED'. .Resume() method does the opposite of .Pause() method.

		BATCH:
			A batch is a collection of tasks, seperated by time.
			A batch is created, with a eunique generated id, when the service queue is in 'COMPLETE' state and a new  task is added to the pending queue.
			If one or more new tasks are added, before the service queue status become 'COMPLETE', then they get added to the existing batch.
			When there are no tasks in the pending queue and in retry queue, then that means the ongoing batch is 'COMPLETE'. At this time a batch report
			is generated and is pushed to mongoDB. At this time the same batch repost is emmited to an event bridge as an event.








	PARAMETERS:
		1. name (string)
			This is the name of the service queue. This will be used to name the queues in redis.
		2. EnableAutoconsume (bool)
			A service que can be configured so that worker(s) start consuming automatically after the first task is added. default=true
		3. workerCount (int)
			The service queue can be configured to have multiple workers. This is the number of workers. default=1
		4. retryCount (int)
			Number of times a task will be retried before putting it in the failed queue. default=0
		5. consumerDelay (int)
			Delay between two pop operations. default=0
		6. status (string)
			Current status of the service queue. CREATED | PENDING | RUNNING | |PAUSED | COMPLETED |
		7. JobID (string)
			Jobs are batcjes ok tasks. One job is complete when status becomes COMPLETE. JobID is auto generated when the first task is added to the queue.
		8. startTime (int64)
			Epoch time when the first task is added to the queue.
		9. endTime (int64)
			Epoch time when the last task is added to the queue.
		10. redis (redis.RedisClient)
			Redis client object. This is used to connect to redis server.


	CONCEPTS:
		1. LoopBackOff (exponential staggered backoff) between every further retry
			if a task fails then the first retry will happen after `restingPeriod` seconds later. If the first retry fails
			then the second retry will happen after `restingPeriod + ((2-1)*loopBackoff)` seconds later.
			the generalized formula is that every nth retry will be rescheduled after `restingPeriod + ((n-1)*loopBackoff)` seconds from the (n-1)th failure.

			it is implemented by
			1. creating a json of retry task which looks like as follows
			{
				"task": {}
				"next_attempt_number": 3
				"next_attempt_scheduled": 1234567890
			}

			2. pushing the json to the queue for that timestamp. i.e. there exist a queue for each timestamp.
			3. the go code wakes up every 500ms, pull all the items from the queue for that timestamp and push them to the main queue.
			4. this deletes the queue for that timestamp.
			5. the main queue also has the same task format (json)






	METHODS:
		1. func NewServiceQ(name string, redisHost string, redisPort string, password ...string) *ServiceQ {}
			Initializes a new service queue object. except for the password every other field is mandatory.
			It uses these default parameters to create a service queue object.
				autoConsume = true
				workerCount = 1
				retryCount = 0
				consumerDelay = 0
				status = PENDING
			This function generates a jobID  (md5(current_epoch)) and captures startTime (epoch).
			Moreover, it creates and attaches a redis cliet object with the service queue object.
			Finally, Returns a pointer to the object.

		2. .Describe() map[string]interface{}
			This function returns a map of all the parameters of the service queue object

		3. .EnableAutostart() & .DisableAutostart()
			These methods can be used to enable or disable auto consume feature of the service queue.
			If the auto consume is enabled, then the worker(s) will start carrying out tarks automatically
			after the first task is added to the queue.

		4. .SetWorkerCount(int) & .GetWorkerCount()
			These methods can be used to set or get the number of workers for the service queue.
			By default the worker count is 1. If the worker count is set to 0, then the service queue will not start any worker.
			However, the tasks can still be added to the queue. This can be used to create a queue of tasks and then start the workers
			manually. workerless ques can be used as general queues. Worker count can be updated anytime, even when the service queue
			is already running. The queue is thread safe.

		5. .SetRetryCount(int) & .GetRetryCount()
			These methods can be used to set or get the retry count for the service queue.
			Retry count is the number of times a task will be retried before putting it in the failed queue.
			By default the retry count is 0. If the retry count is set to 0, then the task will not be retried.
			Retry count can be updated anytime, even when the service queue is already running. The queue is thread safe.

		6. .SetConsumerDelay(int) & .GetConsumerDelay()
			These methods can be used to set or get the consumer delay for the service queue.
			Consumer delay is the delay between two pop operations.
			By default the consumer delay is 0. If the consumer delay is set to 0, then the tasks will be popped as soon as they are pushed.
			Consumer delay can be updated anytime, even when the service queue is already running.

		7. .GetStatus()
			This method returns the current status of the service queue. | PENDING | RUNNING | PAUSED | COMPLETED |

		8. .GetJobID()
			This method returns the jobID of the service queue. JobID is auto generated when the first task is added to the queue.

		9. .GetStartTime()










*/

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tuhin37/goclient/redis"
)

/*
	quecon: queue-consumer design pattern.
	Add items to queue and start a background consumer to process the queue.
	Define the name of the queue. -> This will be used to name the queue uin redis
	Also there are bunch of parameters associated with each queue. e.g. STATUS(RUNNING | PENDING | PAUSED | COMPLETED), start time, end time,
	A job also has two default queues, one for 'failed' items and another for 'completed' items.
	Each job has a unique ID, which is generated when the job is created.
	Each job has a unique name, which is provided by the user.
	Each job has a unique redis queue, which is generated using the job name.
	Each job has a unique redis queue for failed items, which is generated using the job name.
	Each job has a unique redis queue for completed items, which is generated using the job name.





	Methods:
		1. Initialize
		2. Describe => returns a map of all the parameters
		2. Push
		3. Pop
		4. Count
		5. Delete
		6. DeleteAll
		7. PushAny	=> convert struct to string and store
		8. PopAny	=> convert string to struct and return
		9. PushAnyCompressed => convert struct to string and compress and store
		10. PopAnyCompressed => convert string to struct and decompress and return
		11. Status
		12. Start
		13. Stop
		14. Pause
		15. Resume
		16. ConsumerCount (veriadic function, can be used to get or set the consumer count)
		17. ConsumerDelay (veriadic function, can be used to get or set the consumer delay)
		18. export Json
		19. export compressed json
		20. export binary
		18. import Json
		19. import compressed json
		20. import binary
*/

type ServiceQ struct {
	//---------- identity
	id       string // unique id of the service queue. when replicas are present
	name     string // name of the service queue
	isSilent bool   // if true, then the service queue will not log anything
	// ------- worker
	workerCount   int                    // number of workers, that be consuming the queue and carrying out the tasks
	autoStart     bool                   // if true, the worker will start automatically 'waitingPeriod' seconds after the first task is added to the queue
	waitingPeriod int                    // number of seconds the worker waits before start working
	restingPeriod int                    // sleep time between two pop operations
	mu            sync.RWMutex           // for making workerPostBox thread safe
	workerPostBox map[string]interface{} // [<worker-id>]:{status: RUNNING, processed: 2, failed: 1} // a syncer daemon will listen to res channel and update this map.
	// ------- retry config
	retryLimit  int // after these many retries, a job item will be put into failed queue
	loopBackoff int // every nth retry will be rescheduled after `restingPeriod + ((n-1)*loopBackoff)` seconds from the (n-1)th failure.
	// ------- batch report
	batchID        string
	status         string // RUNNING | PENDING | PAUSED | FINISHED
	totalSubmitted int64  // total number of tasks submitted to the queue
	totalPassed    int64  // total number of tasks successfully carried out
	totalFailed    int64  // total number of tasks failed
	startTime      int64  // epoch time when the job was started (first task was added to the queue)
	endTime        int64  // epoch time when the job was finished (last task was carried out)
	batchDuration  int64
	// ------- redis client
	redis redis.RedisClient // redis client
	// ------- custom task function
	taskFunction             func(interface{}) (bool, string)                     // this function will be called for each task
	batchEndCallback         func(map[string]interface{})                         // this function will be called for each time a batch is complete
	batchBeginCallback       func(map[string]interface{})                         // this function will be called for each time a batch starts
	workerPushUpdateCallback func(map[string]interface{}, map[string]interface{}) // this function will be called for each time a worker completes a task

	// --------channels
	start  chan bool        // send start signal to the worker
	stop   chan bool        // send stop signal to the worker
	pause  chan bool        // send pause signal to the worker
	resume chan bool        // send resume signal to the worker
	req    chan interface{} // send payload to worker
	res    chan interface{} // receive response from worker
}

// initialize a service queue with default settings
func NewServiceQ(name string, redisHost string, redisPort string, password ...string) (*ServiceQ, error) {
	// if config exist in redis then this will overwrite it
	s := &ServiceQ{}

	s.startTime = time.Now().Unix()

	s.id = calculateMD5([]byte(fmt.Sprintf("%v", s.startTime) + s.name)[:]) // md5 (current_epoch + name)
	s.name = name
	s.isSilent = true

	s.workerCount = 1
	s.autoStart = true // if true consumer will start automatically, else will have to start manually
	s.workerPostBox = make(map[string]interface{})

	s.status = "CREATED"

	// initialize the signal channels
	s.start = make(chan bool, 32*32)
	s.stop = make(chan bool, 32*32)
	s.pause = make(chan bool, 32*32)
	s.resume = make(chan bool)

	// initialize the data channels
	s.req = make(chan interface{}, 32*32) // max number of workere are chosen to be 32
	s.res = make(chan interface{}, 32*32)

	// initialize a redis client
	s.redis = redis.NewRedis(redisHost, redisPort, "", 0)
	if len(password) == 1 {
		s.redis = redis.NewRedis(redisHost, redisPort, password[0], 0)
	}
	return s, nil
}

// beta
func NewSyncDown(name string, redisHost string, redisPort string, password ...string) (*ServiceQ, error) {
	// initialize a redis client
	s := &ServiceQ{}
	s.name = name

	// initialize a redis client
	s.redis = redis.NewRedis(redisHost, redisPort, "", 0)
	if len(password) == 1 {
		s.redis = redis.NewRedis(redisHost, redisPort, password[0], 0)
	}

	// get the config from redis
	config, ok := s.redis.Get(s.name + "-config")
	if !ok {
		return nil, fmt.Errorf("config not found in redis! ")
	}

	// json string to map
	configMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		return nil, err
	}

	// assign fetched values to a servce queue object
	s.autoStart = configMap["autoStart"].(bool) // if true consumer will start automatically, else will have to start manually
	s.status = configMap["status"].(string)
	s.startTime = int64(configMap["startTime"].(float64))
	s.endTime = int64(configMap["endTime"].(float64))
	s.workerCount = int(configMap["workerCount"].(float64))
	s.waitingPeriod = int(configMap["waitingPeriod"].(float64))
	s.restingPeriod = int(configMap["restingPeriod"].(float64))
	s.retryLimit = int(configMap["retryLimit"].(float64))
	s.loopBackoff = int(configMap["loopBackoff"].(float64))
	// return the object
	return s, nil
}

// defer it to clean up from redis
func (s *ServiceQ) Delete() {

	// delete config from redis
	s.redis.Unset(s.name + "-config")

	// delete pending queue from redis
	s.redis.Unset(s.name + "-pending")

	// delete failed queue from redis
	s.redis.Unset(s.name + "-failed")

	// delete passed queue from redis
	s.redis.Unset(s.name + "-passed")
}

// beta
func (s *ServiceQ) SyncUp() error {
	// package the service queue object and store it in redis
	config := map[string]interface{}{
		"name":          s.name,
		"workerCount":   s.workerCount,
		"autoStart":     s.autoStart,
		"waitingPeriod": s.waitingPeriod,
		"restingPeriod": s.restingPeriod,
		"retryLimit":    s.retryLimit,
		"loopBackoff":   s.loopBackoff,
		"status":        s.status,
		"startTime":     s.startTime,
		"endTime":       s.endTime,
		"taskFunction":  s.taskFunction,
	}

	// map to json
	json, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// store the config in redis
	ok := s.redis.Set(s.name+"-config", string(json))
	if !ok {
		return fmt.Errorf("failed to store the config in redis! ")
	}

	return nil
}

// ------------------ getters & setters ------------------

// returns a map of all the parameters
func (s *ServiceQ) Describe() map[string]interface{} {
	return map[string]interface{}{
		// --- id ---
		"id":   s.id,
		"name": s.name,
		// --- worker ---
		"worker_count":   s.workerCount,
		"auto_start":     s.autoStart,
		"waiting_period": s.waitingPeriod,
		"resting_period": s.restingPeriod,
		// --- retry ---
		"retry_limit": s.retryLimit,
		"loopBackoff": s.loopBackoff,
	}
}

func (s *ServiceQ) SetWorkerConfig(count int, waitingPeriod int, restingPeriod int, autoStart bool) {
	s.mu.Lock()
	prevWorkerCount := s.workerCount
	s.autoStart = autoStart
	s.workerCount = count
	s.waitingPeriod = waitingPeriod
	s.restingPeriod = restingPeriod
	status := s.status
	s.mu.Unlock()

	if count > 0 && autoStart && status == "PENDING" {
		s.Start()
	}

	// if a batch has started, then one can scale worker just by updating worker config (irrespective of autoStart)
	if count == 0 && (status == "RUNNING" || status == "PAUSED") {
		s.Stop()
	}

	// check scaleup
	if count != prevWorkerCount && (status == "RUNNING" || status == "PAUSED") {
		s.Scale()
	}
}

func (s *ServiceQ) SetRetryConfig(retryLimit int, loopBackoff int) {
	s.retryLimit = retryLimit
	s.loopBackoff = loopBackoff
	// s.SyncUp()
}

func (s *ServiceQ) GetStatus() string {
	return s.status // overall status, cimbining all the workers
}

func (s *ServiceQ) GetStatusInfo() map[string]interface{} {
	pendingCount, _ := s.redis.Qlength(s.name + "-pending")
	passedCount, _ := s.redis.Qlength(s.name + "-passed")
	failedCount, _ := s.redis.Qlength(s.name + "-failed")

	if s.status == "RUNNING" || s.status == "PAUSED" {
		s.batchDuration = time.Now().Unix() - s.startTime
		s.endTime = time.Now().Unix()
	}
	return map[string]interface{}{
		"status":          s.status,
		"batch_id":        s.batchID,
		"total_submitted": s.totalSubmitted,
		"total_pending":   s.totalSubmitted - s.totalFailed - s.totalPassed,
		"total_success":   s.totalPassed,
		"total_failed":    s.totalFailed,
		"start_time":      s.startTime,
		"end_time":        s.endTime,
		"batch_duration":  s.batchDuration,

		// unifiled queue data. not per batch
		"pending_q_length": pendingCount,
		"passed_q_length":  passedCount,
		"failed_q_length":  failedCount,

		"worker_reports": s.workerPostBox,
	}
}

func (s *ServiceQ) EnableAutostart() {
	s.autoStart = true

	if s.workerCount > 0 {
		s.Start()
	}
	// s.SyncUp()
}

func (s *ServiceQ) DisableAutostart() {
	s.autoStart = false
	// s.SyncUp()
}

func (s *ServiceQ) Verbose() {
	s.isSilent = false
	// s.SyncUp()
}

func (s *ServiceQ) Silent() {
	s.isSilent = true
	// s.SyncUp()
}

// ------------------ worker operations ------------------
func (s *ServiceQ) SetTaskFunction(f func(interface{}) (bool, string)) {
	s.taskFunction = f
}

// this callback funbction gets called everytime a batch completes.
func (s *ServiceQ) SetBatchEndCallback(f func(map[string]interface{})) {
	s.batchEndCallback = f
}

// this callback function is called when the first worker starts processing the batch (could be help full if workers are scheduled at a future time)
func (s *ServiceQ) SetBatchBeginCallback(f func(map[string]interface{})) {
	s.batchBeginCallback = f
}

// gets called everytime a worker pushes update to the postbox
func (s *ServiceQ) SetWorkerPushUpdateCallback(f func(map[string]interface{}, map[string]interface{})) {
	s.workerPushUpdateCallback = f
}

func (s *ServiceQ) Start() error {
	/*
		1. only responsible for starting the workers
		2. this changes s.status from PENDING to RUNNING
		3. if autostart is set, then PENDING -> RUNNING is autotriggered
	*/

	// read serviceQ state
	s.mu.RLock()
	status := s.status
	desiredWorkerCount := s.workerCount
	// runningWorkerCount := len(s.workerPostBox)
	pendingTaskCount, _ := s.redis.Qlength(s.name + "-pending")
	s.mu.RUnlock()

	// only move forward if in PENDING state
	if status != "PENDING" {
		return fmt.Errorf("can not start! service is not in PENDING state")
	}

	// move forward if desired worker count is more than 0
	if desiredWorkerCount == 0 {
		return fmt.Errorf("can not start! worker count is 0")
	}

	// move forward if there are pending tasks
	if pendingTaskCount == 0 {
		return fmt.Errorf("can not start! no pending tasks")
	}

	// signal workes to start working
	s.mu.Lock()
	s.status = "START" // the workers poll this value to start working
	s.mu.Unlock()

	// create workers
	for i := 0; i < desiredWorkerCount; i++ {
		go worker(s, s.start, s.stop, s.pause, s.resume, s.req, s.res) // create worker(s)
	}

	return nil
}

func (s *ServiceQ) Stop() error {

	s.mu.RLock()
	status := s.status

	// allowed state transitions: PENDING, RUNNING, PAUSED -> STOPPED
	if status != "PENDING" && status != "RUNNING" && status != "PAUSED" {
		return fmt.Errorf("serviceQ can not be stopped, it is not pending, running or paused state")
	}

	// signal all workers to kill themselves, the last to die will do the exit formalities. coded in the worker function
	s.status = "STOPPED"
	s.mu.RUnlock()
	return nil
}

func (s *ServiceQ) Pause() error {
	s.mu.RLock()
	status := s.status
	s.mu.RUnlock()

	if status != "RUNNING" {
		return fmt.Errorf("can not pause. serviceq was not running")
	}

	// signal all workers to pause themselves
	s.mu.Lock()
	s.status = "PAUSED"
	s.mu.Unlock()
	return nil
}

func (s *ServiceQ) Resume() error {
	s.mu.RLock()
	status := s.status
	s.mu.RUnlock()

	if status != "PAUSED" {
		return fmt.Errorf("can not resume. serviceq was not paused")
	}

	// signal all workers to resume themselves
	s.mu.Lock()
	s.status = "RUNNING"
	s.mu.Unlock()

	return nil
}

// scale number of workers
func (s *ServiceQ) Scale() error {
	/*
		1. scalling only makes sense when service is in RUNNING or PAUSED state
		2. if paused and scaled up, then more worker will be created and they will also be in pasued state
		3. if paused and scaled down, then some worker(s) will be deleted
		4. if running and scaled up, then more worker will be created and they will also be in running state
		5. if running and scaled down, then some worker(s) will be deleted
	*/

	// gather info
	s.mu.RLock()
	status := s.status
	desiredWorkerCount := s.workerCount
	runningWorkerCount := len(s.workerPostBox)
	s.mu.RUnlock()

	if status != "RUNNING" && status != "PAUSED" {
		return fmt.Errorf("can not scale. serviceq needs to be in RUNNING or PAUSED state")
	}

	// scale up
	if desiredWorkerCount > runningWorkerCount {
		workersToCreate := desiredWorkerCount - runningWorkerCount
		for i := 0; i < workersToCreate; i++ {
			go worker(s, s.start, s.stop, s.pause, s.resume, s.req, s.res) // create worker(s)
		}
	}

	// get all the worker ids
	var workerIDs []string
	s.mu.RLock()
	for workerID, _ := range s.workerPostBox {
		workerIDs = append(workerIDs, workerID)
	}
	s.mu.RUnlock()

	// scale down
	if desiredWorkerCount < runningWorkerCount {
		workersToKill := runningWorkerCount - desiredWorkerCount
		for i := 0; i < workersToKill; i++ {
			workerID := workerIDs[i]
			s.mu.Lock()
			delete(s.workerPostBox, workerID) // aks the workers to commit cuicide, make sure worker can die while in paused
			s.mu.Unlock()
		}
	}

	return nil
}

// this runs in the background
func worker(s *ServiceQ, start chan bool, stop chan bool, pause chan bool, resume chan bool, req chan interface{}, res chan interface{}) {

	// these parameters will be sent to response channel after each task
	workerParams := map[string]interface{}{}
	workerParams["created_at"] = time.Now().Unix()
	workerParams["id"] = fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))) // auto generate a workerID md5(time.Now().Unix())
	workerParams["status"] = "CREATED"
	workerParams["failed_count"] = 0
	workerParams["success_count"] = 0
	workerParams["uptime"] = time.Now().Unix() - workerParams["created_at"].(int64) // will be calculated at the time of report from start time and current time

	s.mu.Lock()
	s.workerPostBox[workerParams["id"].(string)] = workerParams
	s.mu.Unlock()

	// this is used to detect repeating retry tasks
	var prevChecksum string

	/*
		if there's one retry task, in the pending queue, which is scheduled for a future timestamp.
		just so that the Qtask is not poppped and pushed back to the pending queue, again and again.
		introduce a staggered delay, so that the task is not retried immediately and save CPU cycles.
	*/
	// samePopLoopBackOffDelay := s.loopBackoff

	var svcqStatus string
	for {
		s.mu.RLock()
		svcqStatus = s.status // check if the status has changed
		s.mu.RUnlock()

		if svcqStatus == "START" {
			s.mu.Lock()
			s.status = "RUNNING"
			s.mu.Unlock()
			workerLOG(&workerParams, s, "i'm the first one. doing start formalities")
			batchBeginReport := workerEntryFormalities(&workerParams, s) // system level formalities
			s.batchBeginCallback(batchBeginReport)                       // call user function if set
		}

		if svcqStatus == "RUNNING" {
			// sleep for the waiting period
			workerLOG(&workerParams, s, "waiting for "+strconv.Itoa(s.waitingPeriod)+" seconds before starting")
			time.Sleep(time.Duration(s.waitingPeriod) * time.Second)
			for {

				// check every pop, if ststus is still "RUNNING". it could mutate in "PAUSED" or "STOPPED" or "FINISHED"
				if s.status != "RUNNING" {
					break
				}

				// check mailbox before every pop, if the worker is deleted from the mailbox, then the worker commits cuicide
				if _, ok := s.workerPostBox[workerParams["id"].(string)]; !ok {
					s.mu.Lock()
					runningWorkerCount := len(s.workerPostBox)
					s.mu.Unlock()

					// if this is the last worker, do the exit formalities
					if runningWorkerCount == 0 {
						workerLOG(&workerParams, s, "i am the last one. doing exit formalities")
						exitReport := workerExitFormalities(&workerParams, s)
						s.batchEndCallback(exitReport) // call user function if set
						workerLOG(&workerParams, s, "worker killed")
						return // kill this worker, if this is the last worker
					}

					workerLOG(&workerParams, s, "worker killed")
					return // kill this worker
				}

				// POP a Qtask
				s.mu.Lock()
				// check if this is the last pop, before popping, add a
				QtaskString, err := s.redis.Pop(s.name + "-pending")
				s.mu.Unlock()
				if err != nil {
					workerHandlePopFailure(&workerParams, s, err)
					return // this will kill worker
				}

				// ATP: a task string is popped from the pending queue
				// deserialize the json string
				var Qtask map[string]interface{}
				json.Unmarshal([]byte(QtaskString), &Qtask)

				/*	calculate the checksum of the task body and compare with the checksum of the previous task.
					if the checksum is same, that mostly means its retry task with a future timestamp,
					and its keep getting polled again and again. In that case icrease delay.
				*/
				currChecksum := calculateMD5(Qtask["task"])
				if currChecksum == prevChecksum {

					// wait for the delay and continue the loop
					workerLOG(&workerParams, s, "popped the same Qtask "+Qtask["task_id"].(string)+" again. delaying "+""+" seconds loopbackoff")
					// time.Sleep(time.Duration(samePopLoopBackOffDelay) * time.Second)

				}

				// update the prevChecksum
				prevChecksum = currChecksum

				// if the task has a future timestamp, which means its a retry task. and push is back to the que for future execution
				if int64(Qtask["next_attempt_scheduled"].(float64)) > time.Now().Unix() {
					workerLOG(&workerParams, s, "retry task "+Qtask["task_id"].(string)+" has a future timestamp: "+strconv.Itoa(int(Qtask["next_attempt_scheduled"].(float64)))+" re-adding to pending queue")
					s.mu.Lock()
					s.redis.Push(s.name+"-pending", QtaskString)
					s.mu.Unlock()
					time.Sleep(5 * time.Second) // wait for the delay and continue the loop
					continue
				}

				/*
					Execute the task. This is a user function call. The user function,
					Should return "true" inicating that the task wass  successful or "false" if the task failed.
					Also the user function should return a remark. this remarks is added against the Qtask object in the report.
					A remark has two parts, first part is directive and send part is message
					This remark's directive part is used to decide whether to retry the task or not.
					e.g. "NO_RETRY whatsapp not registered", "SUCCESS user added to engine", "RETRY 360dialog queue full"
					so "NO_RETRY", "SUCCESS" and "RETRY" are the directive part of remarks and the rest is the message part
				*/
				isTaskPass, remark := s.taskFunction(Qtask["task"])

				// get the directive part of the remark
				remarkDirective := strings.Split(remark, " ")[0]
				isRetryPossible := true
				if remarkDirective == "NO_RETRY" {
					isRetryPossible = false
				}

				// task failed
				if !isTaskPass {
					// if retry is configured and possible
					if isRetryPossible && s.retryLimit != 0 && int(Qtask["next_attempt_number"].(float64)) <= s.retryLimit {
						// increase the attempt number
						Qtask["next_attempt_number"] = int(Qtask["next_attempt_number"].(float64)) + 1

						// caculate and set rescheduled timestamp. next_attempt_scheduled_timestamp = currentEpoch + restingPeriod + ((n-1)*loopBackoff)
						delay := int64(s.restingPeriod) + int64((Qtask["next_attempt_number"].(int)-1))*int64(s.loopBackoff)
						Qtask["next_attempt_scheduled"] = time.Now().Unix() + delay
						// fmt.Println("next_attempt_scheduled", Qtask["next_attempt_scheduled"])
						workerLOG(&workerParams, s, "task "+Qtask["task_id"].(string)+" failed, retrying at "+strconv.Itoa(int(Qtask["next_attempt_scheduled"].(int64)))+" epoch")

						// add remark to the Qtask
						Qtask["remark"] = remark

						// convert map to json ([]byte)
						UpdatedQtaskJson, _ := json.Marshal(Qtask)

						// add it back to the pending queue
						s.mu.Lock()
						s.redis.Push(s.name+"-pending", string(UpdatedQtaskJson)) // re-add the task to the pending queue for future execution
						s.endTime = time.Now().Unix()
						s.batchDuration = s.endTime - s.startTime
						s.mu.Unlock()
						workerLOG(&workerParams, s, "Qtask "+Qtask["task_id"].(string)+" added to pending queue for retry, waiting for 1s")
						time.Sleep(1 * time.Second) // wait 1s if the task is pushed for retry. (ratelimit if there are two or more retry tasks in in the pending queue for future timestamp, checksome wont catch becaue its not the same Qtask that gets popped everytime.)
						continue                    // move on the the next Qtast POP
					}

					// ATP: Retry limit reached, push the task to the failed queue
					workerLOG(&workerParams, s, "retry limit reached, pushing the task to failed queue")

					// Process Qtask object before pushng it to the failed queue - for better reporting
					lastRetryAttempt := int(Qtask["next_attempt_number"].(float64))
					lastRetryTimestamp := time.Now().Unix()
					Qtask["total_retry_attempts"] = lastRetryAttempt
					Qtask["last_retry_timestamp"] = lastRetryTimestamp
					Qtask["remark"] = remark
					delete(Qtask, "next_attempt_number")
					delete(Qtask, "next_attempt_scheduled")

					// package Qtask
					UpdatedQtaskJson, _ := json.Marshal(Qtask)

					// push Qtask to redis failed queue
					s.redis.Push(s.name+"-failed", string(UpdatedQtaskJson))

					// worker push updated report to the postbox
					workerParams["failed_count"] = workerParams["failed_count"].(int) + 1
					workerParams["uptime"] = time.Now().Unix() - workerParams["created_at"].(int64) // will be calculated at the time of report from start time and current time
					workerParams["status"] = "RUNNING"
					s.mu.Lock()
					s.workerPostBox[workerParams["id"].(string)] = workerParams
					s.workerPushUpdateCallback(workerParams, Qtask) // send to user defined function
					s.totalFailed += 1
					s.endTime = time.Now().Unix()
					s.batchDuration = s.endTime - s.startTime
					s.mu.Unlock()
					time.Sleep(time.Duration(s.restingPeriod) * time.Second)
					continue // task pushed to failed list, move on to the next QTASK POP
				}

				// ATP: task passed
				// Process Qtask object before pushng it to the failed queue - for better reporting
				lastRetryAttempt := int(Qtask["next_attempt_number"].(float64))
				lastRetryTimestamp := time.Now().Unix()
				Qtask["total_retry_attempts"] = lastRetryAttempt
				Qtask["last_retry_timestamp"] = lastRetryTimestamp
				Qtask["remark"] = remark
				delete(Qtask, "next_attempt_number")
				delete(Qtask, "next_attempt_scheduled")

				// package Qtask
				UpdatedQtaskJson, _ := json.Marshal(Qtask)

				// push the task to the passed queue
				s.redis.Push(s.name+"-passed", string(UpdatedQtaskJson))

				// worker push updated report to the postbox
				workerParams["success_count"] = workerParams["success_count"].(int) + 1
				workerParams["uptime"] = time.Now().Unix() - workerParams["created_at"].(int64) // will be calculated at the time of report from start time and current time
				workerParams["status"] = "RUNNING"
				s.mu.Lock()
				s.workerPostBox[workerParams["id"].(string)] = workerParams
				s.workerPushUpdateCallback(workerParams, Qtask) // send to user defined function
				s.totalPassed += 1
				s.endTime = time.Now().Unix()
				s.batchDuration = s.endTime - s.startTime
				s.mu.Unlock()

				// sleep for the resting period, between two pop operations
				workerLOG(&workerParams, s, "sleeping for resting period of "+strconv.Itoa(s.restingPeriod)+" seconds")
				time.Sleep(time.Duration(s.restingPeriod) * time.Second)
			}
		}

		if svcqStatus == "PAUSED" {
			workerLOG(&workerParams, s, "worker is paused")
			// worker push updated report to the postbox
			workerParams["uptime"] = time.Now().Unix() - workerParams["created_at"].(int64) // will be calculated at the time of report from start time and current time
			workerParams["status"] = "PAUSED"
			s.mu.Lock()
			s.workerPostBox[workerParams["id"].(string)] = workerParams
			s.totalFailed += 1
			s.endTime = time.Now().Unix()
			s.batchDuration = s.endTime - s.startTime
			s.mu.Unlock()
			time.Sleep(time.Duration(s.restingPeriod) * time.Second)

			for {
				if s.status == "RUNNING" || s.status == "STOPPED" { // only these two status are allowed to break the loop
					break
				}
			}
		}

		if svcqStatus == "STOPPED" {
			s.mu.Lock()
			runningWorkerCount := len(s.workerPostBox)
			delete(s.workerPostBox, (workerParams)["id"].(string))
			s.mu.Unlock()

			// if this is the last worker, do the exit formalities
			if runningWorkerCount == 1 {
				workerLOG(&workerParams, s, "i am the last one. doing exit formalities")
				batchEndReport := workerExitFormalities(&workerParams, s)
				s.batchEndCallback(batchEndReport) // send to user defined function
				workerLOG(&workerParams, s, "worker killed")
				return // kill this worker, if this is the last worker
			}

			workerLOG(&workerParams, s, "worker killed")
			return // kill this worker
		}
	}
}

func workerHandlePopFailure(workerParams *map[string]interface{}, s *ServiceQ, err error) {
	/* 	handle pop failure
	1. if the pop operation has failed, first check is it because the queue is empty
	2. if the above is true, that means no more task is remaining. print that.
	3. check if this is the last worker. if yes, do exit formalities (system + user) and kill the worker
	4. if this is not the last worker, then just kill this worker
	5. when this function returns, the parent function will kill this worker
	6. if the pop operation has failed for some other reason, then kill the worker with the error message
	*/
	if strings.Trim(err.Error(), "\"") == "redis: nil" {
		workerLOG(workerParams, s, "no Qtask pending") // do exit formalities, i.e. moving Qtasks from pending to failed, send report to mongodb, updating s.status to STOPPED
		s.mu.Lock()
		runningWorkerCount := len(s.workerPostBox)
		delete(s.workerPostBox, (*workerParams)["id"].(string))
		s.mu.Unlock()

		if runningWorkerCount == 1 {
			workerLOG(workerParams, s, "i am the last one. doing exit formalities")
			batchReport := workerExitFormalities(workerParams, s)
			s.batchEndCallback(batchReport) // send to user defined function
			workerLOG(workerParams, s, "worker terminated")
			return // kill this worker
		}
		workerLOG(workerParams, s, "worker terminated")
		return
	}

	// if the pop failed for some other reason, e.g. redis connection error, kill with the error msg
	workerLOG(workerParams, s, "worker terminated! pop failed with error: "+err.Error())
}

func workerEntryFormalities(workerParams *map[string]interface{}, s *ServiceQ) map[string]interface{} {
	// doing the entry formalities
	workerLOG(workerParams, s, "example entry formality task")

	s.mu.Lock()
	defer s.mu.Unlock()
	s.endTime = time.Now().Unix()
	s.batchDuration = s.endTime - s.startTime

	// return batch report
	return map[string]interface{}{
		"serviceq_id":     s.id,
		"serviceq_name":   s.name,
		"status":          s.status,
		"batch_id":        s.batchID,
		"total_submitted": s.totalSubmitted,
		"total_pending":   s.totalSubmitted - s.totalFailed - s.totalPassed,
		"total_success":   s.totalPassed,
		"total_failed":    s.totalFailed,
		"start_time":      s.startTime,
		"end_time":        s.endTime,
		"batch_duration":  s.batchDuration,
		"passed_tasks":    []interface{}{},
		"failed_tasks":    []interface{}{},
		"pending_tasks":   []interface{}{},
	}

}

func workerExitFormalities(workerParams *map[string]interface{}, s *ServiceQ) map[string]interface{} {
	// doing the entry formalities
	workerLOG(workerParams, s, "example exit formality task")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.endTime = time.Now().Unix()
	s.batchDuration = s.endTime - s.startTime

	// get all the pending Qtasks (in case of force stopped or retry failed tasks)
	var pendingQtasks []interface{}
	var pendingQtaskStr string
	var pendingQtask interface{}
	totalPendingCount, _ := s.redis.Qlength(s.name + "-pending")
	for i := int64(0); i < totalPendingCount; i++ {
		pendingQtaskStr, _ = s.redis.Pop(s.name + "-pending")
		json.Unmarshal([]byte(pendingQtaskStr), &pendingQtask)
		pendingQtasks = append(pendingQtasks, pendingQtask)
	}

	// get all Qtasks from the passed lists
	var passedQtasks []interface{}
	var passQtaskStr string
	var passedQtask interface{}
	totalPAssedCount, _ := s.redis.Qlength(s.name + "-passed")
	for i := int64(0); i < totalPAssedCount; i++ {
		passQtaskStr, _ = s.redis.Pop(s.name + "-passed")
		json.Unmarshal([]byte(passQtaskStr), &passedQtask)
		passedQtasks = append(passedQtasks, passedQtask)
	}

	// get all Qtasks from the failed lists
	var failedQtasks []interface{}
	var failedQtaskStr string
	var failedQtask interface{}
	totalFailedCount, _ := s.redis.Qlength(s.name + "-failed")
	for i := int64(0); i < totalFailedCount; i++ {
		failedQtaskStr, _ = s.redis.Pop(s.name + "-failed")
		json.Unmarshal([]byte(failedQtaskStr), &failedQtask)
		failedQtasks = append(failedQtasks, failedQtask)
	}

	if s.status == "RUNNING" {
		s.status = "FINISHED"
	}

	// return batch report
	return map[string]interface{}{
		"serviceq_id":     s.id,
		"serviceq_name":   s.name,
		"status":          s.status,
		"batch_id":        s.batchID,
		"total_submitted": s.totalSubmitted,
		"total_pending":   s.totalSubmitted - s.totalFailed - s.totalPassed,
		"total_success":   s.totalPassed,
		"total_failed":    s.totalFailed,
		"start_time":      s.startTime,
		"end_time":        s.endTime,
		"batch_duration":  s.batchDuration,
		"passed_tasks":    passedQtasks,
		"failed_tasks":    failedQtasks,
		"pending_tasks":   pendingQtasks,
	}
}

func workerLOG(workerParams *map[string]interface{}, s *ServiceQ, log string) {
	if s.isSilent {
		return
	}
	fmt.Println(time.Now().Unix(), " | ", (*workerParams)["id"].(string), " | ", log)
}

// ----------------------------------------- queue operations -----------------------------------------

// push a new item to the queue
func (s *ServiceQ) Push(task interface{}) error {

	// check if the pending queue is empty
	pendingQLen, err := s.redis.Qlength(s.name + "-pending")
	if err != nil {
		return err
	}

	// reset batch if the tasl is pushed after previous batch is finished (pending queue is empty)
	if pendingQLen == 0 {
		// start a new batch
		s.mu.Lock()
		s.batchID = calculateMD5([]string{strconv.Itoa(int(time.Now().Unix())), s.name}) // md5(current_epoch + serviceQ_name)
		s.status = "PENDING"                                                             // worker will update this
		s.totalSubmitted = 0
		s.totalFailed = 0
		s.totalPassed = 0
		s.startTime = time.Now().Unix()
		s.endTime = s.startTime
		s.batchDuration = 0
		s.mu.Unlock()
	}

	// construct the task object
	Qtask := map[string]interface{}{
		"task":                   task,
		"task_id":                calculateMD5([]string{strconv.Itoa(int(time.Now().UnixNano())), s.name, fmt.Sprintf("%v", task)}),
		"batch_id":               s.batchID,
		"serviceq_id":            s.id,
		"created_at":             time.Now().Unix(),
		"next_attempt_number":    0,
		"next_attempt_scheduled": 0,
		"remark":                 "",
	}

	// convert map to json ([]byte)
	QtaskJson, err := json.Marshal(Qtask)
	if err != nil {
		return err
	}

	// stringify the json and push to the queue
	err = s.redis.Push(s.name+"-pending", string(QtaskJson))

	if err == nil {
		s.mu.Lock()
		s.totalSubmitted++
		s.mu.Unlock()
	}

	// start the serviceQ if autoStart is true
	if s.autoStart && s.status != "RUNNING" {
		s.Start()
	}

	return err // this error is from redis push operation
}

// ------------------------------------------ helper functions ------------------------------------------

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

func StringSliceDifference(biggSlice, smallSlice []string) []string {
	// Create a map to hold the values in smaller slice
	m := make(map[string]bool)
	for _, item := range smallSlice {
		m[item] = true
	}

	// Create a slice to hold the difference
	var diff []string

	// Check each item in slice1
	for _, item := range biggSlice {
		// If the item is not in the map, add it to the difference slice
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}
