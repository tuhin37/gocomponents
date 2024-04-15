package main

import (
	"fmt"
	"sync"

	"github.com/tuhin37/gocomponents/runner"
)

var runr1 *runner.Runner
var runr2 *runner.Runner
var runr3 *runner.Runner
var runr4 *runner.Runner

func init() {
	// get IPv4 address of `facebook.com`
	runr1 = runner.NewRunner("wget https://upload.wikimedia.org/wikipedia/commons/3/38/Prometheus_software_logo.svg")
	runr2 = runner.NewRunner("wget https://upload.wikimedia.org/wikipedia/commons/a/a1/Grafana_logo.svg")
	runr3 = runner.NewRunner("wget https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg")
	runr4 = runner.NewRunner("ls -l | grep .svg")

}

func main() {
	// use wait group to wait for all the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(3)

	// launch runr1 as a goroutine, async, non-blocking call
	go func() {
		defer wg.Done()
		runr1.Execute()
	}()

	// launch runr1 as a goroutine, async, non-blocking call
	go func() {
		defer wg.Done()
		runr2.Execute()
	}()

	// launch runr1 as a goroutine, async, non-blocking call
	go func() {
		defer wg.Done()
		runr3.Execute()
	}()

	// wait for all the goroutines to finish
	wg.Wait()

	// list all the downloaded logo files. This only runs when the above 3 runner goroutines finish
	runr4.Execute()
	fmt.Println(runr4.Logs())

}
