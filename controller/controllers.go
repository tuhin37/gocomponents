package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/gocomponents/serviceq"
)

var svcQ *serviceq.ServiceQ

func task(data interface{}) (bool, string) {
	// fmt.Println("TASK_FUNCTION: ", data)

	// if the task fails
	// return false, "NO_RETRY the user does not have whatsapp number"

	return true, "success"
}

func init() {
	svcQ, _ = serviceq.NewServiceQ("drag", "localhost", "6379", "")
	svcQ.DisableAutostart()
	svcQ.SetWorkerConfig(0, 1, 2)
	svcQ.SetRetryConfig(5, 3)
	svcQ.SetTaskFunction(task)
	fmt.Println(svcQ.Describe())

	// 	svcQ.Start()

	// // svcQ.Push(map[string]interface{}{"name": "tuhin", "age": 30, "address": "Bangalore"})
}

func Describe(c *gin.Context) {
	c.AsciiJSON(200, svcQ.Describe())
}

func SetWorker(c *gin.Context) {
	var workerConfig map[string]interface{}
	c.BindJSON(&workerConfig)

	workerCount := int(workerConfig["worker_count"].(float64))
	waitingPeriod := int(workerConfig["waiting_period"].(float64))
	restingPeriod := int(workerConfig["resting_period"].(float64))
	isAutoStart := workerConfig["auto_start"].(bool)

	svcQ.SetWorkerConfig(workerCount, waitingPeriod, restingPeriod)
	if isAutoStart {
		svcQ.EnableAutostart()
	} else {
		svcQ.DisableAutostart()
	}
	c.AsciiJSON(200, svcQ.Describe())
}

func GetStatusInfo(c *gin.Context) {
	c.AsciiJSON(200, svcQ.GetStatusInfo())
}

func Add(c *gin.Context) {
	var tasks interface{}

	c.BindJSON(&tasks)

	for _, task := range tasks.([]interface{}) {
		svcQ.Push(task)
	}

	c.AsciiJSON(200, gin.H{"msg": fmt.Sprintf("%d tasks submitted", len(tasks.([]interface{})))})

	config := svcQ.Describe()
	// if auto start possible then start
	if config["auto_start"].(bool) && config["worker_count"].(int) > 0 {
		svcQ.Start() // start all workers
	}
}

func Start(c *gin.Context) {
	err := svcQ.Start()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"msg": err.Error()})
		return
	}

	c.AsciiJSON(200, svcQ.GetStatusInfo())

}

func Stop(c *gin.Context) {
	svcQ.Stop()
}

func Pause(c *gin.Context) {
	err := svcQ.Pause()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"msg": err.Error()})
		return
	}

	c.AsciiJSON(200, svcQ.GetStatusInfo())
}
