# Runner example - Advanced Overrides

In this section we will look at some advanced override examples

---

Example-1: reuse same runner instance with a different command

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("whoami")
}

func main() {
    runr.Execute()
    fmt.Println(runr.Logs())

    // reuse same runner with different command
    runr.Execute("pwd")
    fmt.Println(runr.Logs())
}
```

Output

```shell
drag

/home/drag/programming/personal/go/module/gocomponents
```

---

Example-2: attach multiple system command to the same runner instance at once

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr *runner.Runner

func init() {
    runr = runner.NewRunner("")
}

func main() {

    // reuse runner instance with multiple commands
    runr.Execute("pwd", "whoami", "which python")
    fmt.Println(runr.Logs())
}
```

Output

```shell
/home/drag/programming/personal/go/module/gocomponents
drag
/usr/bin/python
```

---

Example-3: Multiple runner instance running sequentially

```go
package main

import (
    "fmt"

    "github.com/tuhin37/gocomponents/runner"
)

var runr1 *runner.Runner
var runr2 *runner.Runner

func init() {
    // get IPv4 address of `facebook.com`
    runr1 = runner.NewRunner("nslookup facebook.com |grep Address |sed -n '2p' |awk '{print $2}'")
    runr2 = runner.NewRunner("")
}

func main() {
    // find IP address of `facebook.com`
    runr1.Execute()
    facebookIPv4 := runr1.Logs()
    fmt.Println("facebook's IPv4 address:  ", facebookIPv4)

    // ping the same IP address 5 times
    runr2.Execute("ping " + facebookIPv4 + " -c 5")
    fmt.Println(runr2.Logs())
}
```

Output

```shell
facebook's IPv4 address:   157.240.23.35
PING 157.240.23.35 (157.240.23.35) 56(84) bytes of data.
64 bytes from 157.240.23.35: icmp_seq=1 ttl=57 time=11.9 ms
64 bytes from 157.240.23.35: icmp_seq=2 ttl=57 time=11.5 ms
64 bytes from 157.240.23.35: icmp_seq=3 ttl=57 time=10.8 ms
64 bytes from 157.240.23.35: icmp_seq=4 ttl=57 time=10.3 ms
64 bytes from 157.240.23.35: icmp_seq=5 ttl=57 time=11.4 ms

--- 157.240.23.35 ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4005ms
rtt min/avg/max/mdev = 10.349/11.182/11.852/0.532 ms
```

---

Example-4: Three runners running in parallel (non-blocking, goroutines, async) and the final runner only runs when the first three runner finishes.

```go
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
    // download svg logo files
	runr1 = runner.NewRunner("wget https://upload.wikimedia.org/wikipedia/commons/3/38/Prometheus_software_logo.svg")
	runr2 = runner.NewRunner("wget https://upload.wikimedia.org/wikipedia/commons/a/a1/Grafana_logo.svg")
	runr3 = runner.NewRunner("wget https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg")
	
    // list svg files from current directory
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

	// launch runr2 as a goroutine, async, non-blocking call
	go func() {
		defer wg.Done()
		runr2.Execute()
	}()

	// launch runr3 as a goroutine, async, non-blocking call
	go func() {
		defer wg.Done()
		runr3.Execute()
	}()

	// wait for all the three runnes to finish
	wg.Wait()

	// list all the downloaded svg files. This runner runs only when the above three runners finish
	runr4.Execute()
	fmt.Println(runr4.Logs())
}
```



Output

```sh
-rw-r--r--  1 drag drag  6753 Jan 14  2023 Grafana_logo.svg
-rw-r--r--  1 drag drag  2777 Jan  9  2020 Prometheus_software_logo.svg
-rw-r--r--  1 drag drag 49983 Mar 21  2022 Tux.svg
```
