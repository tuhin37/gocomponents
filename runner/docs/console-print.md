# Runner example - Console print

In this example, we'll observe the runner still functions as a blocking operation. However, this time, whatever the system command prints to `STDOUT`, the runner will immediately display it in the console in real-time.

This functionality is achieved by configuring the runner settings using `EnableConsole()` method

---
Example

```go
package main

import (
    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ping 8.8.8.8 -c 4") // instantiate a runner
    runr.EnableConsole()
}

func main() {
    runr.Execute()
}
```

Output

```shell
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=118 time=22.8 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=118 time=23.6 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=118 time=25.0 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=118 time=23.5 ms

--- 8.8.8.8 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 22.810/23.717/24.955/0.777 ms
```

Note: You can utilize the `.DisableConsole()` method to halt printing to the console. However, the runner will continue to capture all `STDOUT` outputs from the system call, storing them in its `logBuffer` property. You can access this buffer using the `.Logs()` method.
