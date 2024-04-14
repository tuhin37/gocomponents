package main

import (
	"github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
	runr = runner.NewRunner("ping 8.8.8.8 -c 5")
	runr.EnableConsole()
	runr.SetWaitingPeriod(3) // wait 3 seconds before executing the system command
}

func main() {
	runr.Execute()
}
