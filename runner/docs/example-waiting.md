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
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=59 time=14.7 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=59 time=13.8 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=59 time=13.7 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=59 time=14.4 ms
64 bytes from 8.8.8.8: icmp_seq=5 ttl=59 time=13.8 ms

--- 8.8.8.8 ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4006ms
rtt min/avg/max/mdev = 13.739/14.072/14.656/0.367 ms
```
