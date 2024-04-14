# Runner example - Kill

For emergency stop execution of the system command, we have provided `.Kill()` method to forcefully and ondemand kill and ongoing running system command. If a runner is killed, then the `status` is set to `KILLED`In cases of emergency where the execution of a system command needs to be immediately stopped, we've implemented the `.Kill()` method as a kill switch. This method allows for the forceful termination of any ongoing system command execution.

Upon invoking `.Kill()`, the runner sets its `status` to `KILLED`, indicating that the execution has been forcefully terminated. This provides a clear indication of the abrupt interruption of the command's execution.

---

Example 

```go
package main

import (
    "fmt"
    "time"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ping 8.8.8.8")
    runr.EnableConsole()
}

func main() {
    go runr.Execute()

    // Introduce a delay of 2 seconds
    time.Sleep(2 * time.Second)

    // print status before killed
    fmt.Println("status: ", runr.GetStatus())

    // kill the runner
    runr.Kill()

    // print status after killed
    fmt.Println("status: ", runr.GetStatus())
}
```

Output

```shell
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=59 time=15.7 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=59 time=16.8 ms
status:  RUNNING
status:  KILLED
```

Explation:

The code performs the following steps

- Instantiate a new runner object `runr` with a system command `ping 8.8.8.8` which will run for ever if not stopped.

- The runner is configured to print the logs from the system command to the console

- The runner is made to execute the system command as a go routine. In this case the execution happens on a seperate thred and the execution becomes non blocking or async. 

- The main function wait for 2 seconds, while the execution is going on on a seperate thresh and the logs from the execution are being printed on the console.

- After 2 seconds, just for a check, the status of the runner is made to print and as expected it turned out to be in a`RUNNING` state

- At this point the runner is killed by invoking `.Kill()` method. 

- Immidiately after this lets check the status of the runner object and as expected it was found to be `KILLED`
