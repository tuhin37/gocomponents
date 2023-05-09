package main

import (
	"net/http"

	"github.com/tuhin37/gocomponents/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 64 << 20 // 64 megabytes
	// -------------------------------------- hypd --------------------------------------
	r.POST("/add", controller.Add)
	r.GET("/start", controller.Start)
	r.GET("/stop", controller.Stop)
	r.GET("/status", controller.Status)

	// ------------------------------------- health -------------------------------------
	r.GET("health", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"service": "drag-test",
			"version": "1.0.6",
		})
	})

	// r.GET("/trigger", consumer.ConsumeAndInvite)
	r.Run(":5000")
}
