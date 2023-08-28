package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/gocomponents/runner"
)

// instantiate a runner
var runr *runner.Runner

// initialize the serviceQ with some default values
func init() {
	runr = runner.NewRunner()                           // instantiate a runner
	runr.SetLogFile("log.md")                           // logfile name; if not defined
	runr.SetVerificationPhrase("main", successCallback) // verification phrase, if found in the output then status - > SUCCEEDED
	runr.EnableConsole()                                // the system output will be printed on console
	runr.SetWaitingPeriod(10)                           // wait 10s before execuiting
	runr.SetTimeout(5)                                  // set timeout of 5 seconds.
	runr.OnNewLineCallback(logCallback)                 // attach callback on log
}

func main() {
	r := gin.Default()

	// setup routes
	r.POST("/exec-payload", execPayload)
	r.POST("/exec-payload-async", execPayloadAsync)
	r.GET("/state", getState)
	r.GET("/status", getStatus)
	r.GET("/kill", kill)
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
	stdout, _ := runr.Execute(command, failCallback, completeCallback)

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
// onLogLineCallback
func logCallback(logLine []byte) {
	fmt.Println("Log: ", string(logLine))
}

// onSuccessCallback
func successCallback(stdout []byte) {
	fmt.Println("execuition successful: ", string(stdout))
}

// onFailCallback
func failCallback(stdout []byte) {
	fmt.Println("execuition failed: ", string(stdout))
}

// onCompleteCallback
func completeCallback(stdout []byte) {
	fmt.Println("execuition completed: ", string(stdout))
}

// onTimeoutCallback
func timeoutCallback(stdout []byte) {
	fmt.Println("execuition timedout: ", string(stdout))
}
