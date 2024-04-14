# Runner example - Success Criteria

Optionally, you can define a string as a success criterion. The runner will search for this string in the output of the system call. If the runner finds the specified string, it will consider the system call to have returned the expected output, and consequently set its `status` as `SUCCEEDED`.

---

Example-1 (seccess)

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ping 8.8.8.8 -c 5")
    runr.SetSuccessCriteria("icmp_seq") // success criteria
}

func main() {
    runr.Execute()
    fmt.Println(runr.GetStatus())
}
```

Output

```shell
SUCCEEDED
```

---

Example-2 (fail)

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ping 8.8.8.8 -c 5")
    runr.SetSuccessCriteria("foo") // success criteria
}

func main() {
    runr.Execute()
    fmt.Println(runr.GetStatus())
}
```

Output

```shell
FAILED
```

---

Moreover, you have the option to attach a user-defined function as a success callback function. When the runner's `status` becomes `SUCCEEDED`, this function is automatically invoked. The runner also injects itself into the scope of the callback function, allowing users to access all properties and methods of the runner object, such as `.logs()`.

Example-3 (callback)

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ping 8.8.8.8 -c 5")
    runr.SetSuccessCriteria("icmp_seq", onSuccessCallback) // configure success criteria and callback function
}

func main() {
    runr.Execute()
}

// callback function
func onSuccessCallback(r *runner.Runner) {
    fmt.Println(r.Logs())
}
```

Output

```shell
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=59 time=11.6 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=59 time=12.3 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=59 time=15.1 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=59 time=15.3 ms
64 bytes from 8.8.8.8: icmp_seq=5 ttl=59 time=15.3 ms

--- 8.8.8.8 ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4006ms
rtt min/avg/max/mdev = 11.560/13.909/15.347/1.641 ms
```

Here the logs are printed from within the callback function.
