package main

import (
	"fmt"

	"github.com/tuhin37/gocomponents/serviceq"
)

// import

// Define the function to be passed to SetWorkerFunction
var task = func(t interface{}) bool {
	fmt.Println("from task  function: ", t)
	return false
}

func main() {
	svcQ, _ := serviceq.NewServiceQ("drag", "localhost", "6379", "")
	svcQ.SetWorkerConfig(1, 1, 0)
	svcQ.SetRetryConfig(5, 3)
	svcQ.SetTaskFunction(task)
	fmt.Println(svcQ.Describe())

	svcQ.Start()

	// svcQ.Push(map[string]interface{}{"name": "tuhin", "age": 30, "address": "Bangalore"})

	// svcQ.Delete()
}
