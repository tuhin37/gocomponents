package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/gocomponents/runner"
)

// instantiate a runner
var runr *runner.Runner

// initialize the serviceQ with some default values
func init() {
	runr = runner.NewRunner("ping google.com -c 4") // instantiate a runner
	runr.SetLogFile("log.md")                       // logfile name; if not defined
	// runr.SetVerificationPhrase(".go")               // verification phrase, if found in the output then status - > SUCCEEDED
	// runr.EnableConsole() // the system output will be printed on console
	// runr.SetWaitingPeriod(10)                       // wait 10s before execuiting
	// runr.SetTimeout(5) // set timeout of 5 seconds.

	runr.SetOnNewLineCallback(logCallback)
}

func main() {
	r := gin.Default()
	// -------------------------------------- hypd --------------------------------------
	// r.POST("/update", update)
	// r.POST("/exec", exec)
	// r.POST("/exec-async", execAsync)
	r.POST("/exec-payload", execPayload)
	r.POST("/exec-payload-async", execPayloadAsync)

	// r.GET("/logs", logs)
	r.GET("/state", getState)
	r.GET("/status", getStatus)
	r.GET("/kill", kill)
	// r.GET("/restart", restart)

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

func getStatus(c *gin.Context) {
	c.AsciiJSON(200, runr.GetStatus())
}

func kill(c *gin.Context) {
	runr.Kill()
	c.AsciiJSON(200, runr.GetStatus())
}

func execPayload(c *gin.Context) {
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

	command := data["instruction"]
	_ = command

	// ATP: command holds the system call command. e.g. "ls -al | grep main && tree ."
	stdout, _ := runr.Execute(command)

	c.String(200, string(stdout))
}

func execPayloadAsync(c *gin.Context) {
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

	command := data["instruction"]
	_ = command

	// ATP: command holds the system call command. e.g. "ls -al | grep main && tree ."
	go runr.Execute(command)

	c.String(200, "ok")
}

// ------------------------------------- callbacks -------------------------------------
func logCallback(logLine []byte) {
	fmt.Println("Log: ", string(logLine))
}
