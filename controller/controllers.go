package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/gocomponents/serviceq"
)

var svcQ *serviceq.ServiceQ

func task(data interface{}) bool {
	fmt.Println("Task received")
	fmt.Println(data)
	return true
}

func init() {
	svcQ, _ = serviceq.NewServiceQ("drag", "localhost", "6379", "")
	svcQ.DisableAutostart()
	svcQ.SetWorkerConfig(1, 1, 2)
	svcQ.SetRetryConfig(5, 3)
	svcQ.SetTaskFunction(task)
	fmt.Println(svcQ.Describe())

	// 	svcQ.Start()

	// // svcQ.Push(map[string]interface{}{"name": "tuhin", "age": 30, "address": "Bangalore"})
}

func Add(c *gin.Context) {
	var bla interface{}

	c.BindJSON(&bla)
	fmt.Println("bla: ", bla)

	for _, bl := range bla.([]interface{}) {
		svcQ.Push(bl)
	}
	// svcQ.Push(bla)
}

func Start(c *gin.Context) {
	svcQ.Start()
}

func Stop(c *gin.Context) {
	svcQ.Stop()
}

func Status(c *gin.Context) {
	fmt.Println("jojo")
	svcQ.GetWorkerInfo()
	fmt.Println("yoyo")
	c.JSON(200, svcQ.GetWorkerInfo())
}
