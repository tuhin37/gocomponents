# Runner - Simple

#### Explaination

Import the go components library in your code. runner is defined in this library

```go
import (
    "github.com/tuhin37/gocomponents/runner"
)
```

Instantiate a global variable called `runr` 

```go
var runr *runner.Runner
```

Configure the `runr` object in an init function. In this case we will set the system command that we want to execute and the system command is `ping 8.8.8.8 -c 4`

```go
func init() {
    runr = runner.NewRunner("ping 8.8.8.8 -c 4") // instantiate a runner
}
```

In the main function let go runner execute the system command. Please note that this is a blocking function call.

```go
runr.Execute()
```

The output of the system call is stored in `logBuffer` property of the `runr` object. To print the output of the system call, after the system call is done execuiting invoke the `.Logs()`  method. This method returns string.

```go
fmt.Println(runr.Logs())
```

---

#### Complete code - `main.go`

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ping 8.8.8.8 -c 4")
}

func main() {
    runr.Execute()
    fmt.Println(runr.Logs())
}
```

Output

```shell
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=118 time=22.8 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=118 time=22.6 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=118 time=22.6 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=118 time=26.5 ms

--- 8.8.8.8 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 22.570/23.625/26.542/1.685 ms
```
