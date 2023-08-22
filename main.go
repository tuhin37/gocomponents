package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/gocomponents/runner"
)

// instantiate a runner
var runr *runner.Runner

// initialize the serviceQ with some default values
func init() {
	runr = runner.NewRunner("ping google.com -c 4") // serviceQ name is drag, redis running in loclahost:6379, no password
	runr.SetLogFile("log.md")
	runr.EnableConsole()

}

func main() {
	r := gin.Default()
	// -------------------------------------- hypd --------------------------------------
	// r.POST("/update", update)
	r.GET("/exec", execute)
	r.POST("/execp", execp)
	// r.GET("/logs", logs)
	r.GET("/state", getState)
	// r.GET("/status", status)
	// r.GET("/stop", stop)
	// r.GET("/restart", restart)

	// ------------------------------------- health -------------------------------------
	r.GET("health", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	r.Run(":5000")
}

// ------------------------------------- controller -------------------------------------

func getState(c *gin.Context) {
	c.AsciiJSON(200, runr.GetState())
}

func execute(c *gin.Context) {
	stdout, _ := runr.Execute()
	c.String(200, stdout)

}

func execp(c *gin.Context) {
	// Read the request body
	requestBodyBytes, err := c.GetRawData()
	if err != nil {
		log.Println("Error reading request body:", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// Parse the request body JSON into a map
	var data map[string]string
	if err := json.Unmarshal(requestBodyBytes, &data); err != nil {
		log.Println("Error parsing JSON:", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Iterate through the map and print key-value pairs
	for key, value := range data {
		log.Printf("Key: %s, Value: %v\n", key, value)
	}

	command := data["instruction"]

	// ATP: command holds the system call command. e.g. "ls -al | grep main && tree ."

	stdout, _ := runr.ExecutePayloadStream(command)

	c.String(200, stdout)
}
