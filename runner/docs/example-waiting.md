# Runner example - Waiting Period

You have the option to set the runner to wait for a specific amount of time before it executes the system command. Think of it as a delay timer, similar to the shutter delay feature found in DSLR.

---

Example

```go
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
```

Output

```shell

```
