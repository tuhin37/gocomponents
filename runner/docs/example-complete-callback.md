## Runner example - Complete Callback & Failed Callback

If the success criterion is not defined, the runner cannot determine whether the system command execution was successful. However, it can ascertain if the system command exited with an exit code of 0 (zero). In such cases, the runner updates its status to `COMPLETED`. For any other non-zero exit code, the runner considers the system call to have failed and updates its own `status` to `FAILED`.

Additionally, you have the option to attach a user-defined function as an on-complete callback function to the runner object. This callback function is automatically triggered as soon as the runner's status is updated to `COMPLETED`. The runner object injects itself into the callback function, enabling users to access the properties and methods of the runner object from within the callback function.

Furthermore, you can attach another user-defined callback function for when the runner's system command execution exits with a non-zero exit code. In such situations, the runner will trigger the on-failed-callback while injecting its own instance into the callback function.

---

Example-1 (completed)

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("ls -l")
    runr.SetOnCompleteCallback(onCompleteCallback)
    runr.SetOnFailCallback(onFailCallback)
}

func main() {
    runr.Execute()
}

// callback function
func onCompleteCallback(r *runner.Runner) {
    fmt.Println("callback: onCompleteCallback")
    fmt.Println(r.GetStatus())
}

func onFailCallback(r *runner.Runner) {
    fmt.Println("callback: onFailCallback")
    fmt.Println(r.GetStatus())
}
```

Output

```shell
callback: onCompleteCallback
COMPLETED
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
    runr = runner.NewRunner("foo")
    runr.SetOnCompleteCallback(onCompleteCallback)
    runr.SetOnFailCallback(onFailCallback)
}

func main() {
    runr.Execute()
}

// on success callback
func onCompleteCallback(r *runner.Runner) {
    fmt.Println("callback: onCompleteCallback")
    fmt.Println(r.GetStatus())
}

// on fail callback
func onFailCallback(r *runner.Runner) {
    fmt.Println("callback: onFailCallback")
    fmt.Println(r.GetStatus())
}
```

Output

```sh
callback: onCompleteCallback
FAILED
```
