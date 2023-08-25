[1692814922671340] command
```shell
ls -al
```

[1692814922671341] Response begin
```shell
total 24
drwxr-xr-x 1 drag drag  162 Aug 23 23:52 .
drwxr-xr-x 1 drag drag   52 May 31 01:02 ..
-rw-r--r-- 1 drag drag    0 Aug 23 23:51 foo.bar
drwxr-xr-x 1 drag drag  144 Aug 23 01:52 .git
-rw-r--r-- 1 drag drag 1429 May 10 00:38 go.mod
-rw-r--r-- 1 drag drag 8386 May 10 00:38 go.sum
drwxr-xr-x 1 drag drag   86 May  4 12:52 graph
drwxr-xr-x 1 drag drag   40 May  4 12:36 heap
-rw-r--r-- 1 drag drag   91 Aug 23 23:52 log.md
-rw-r--r-- 1 drag drag 3120 Aug 23 23:51 main.go
drwxr-xr-x 1 drag drag   38 May  4 12:26 queue
drwxr-xr-x 1 drag drag   18 Aug 18 22:56 runner
drwxr-xr-x 1 drag drag   68 May 20 22:33 serviceq
drwxr-xr-x 1 drag drag   12 May  4 12:55 set
drwxr-xr-x 1 drag drag   38 May  4 12:26 stack
drwxr-xr-x 1 drag drag   70 May  4 12:31 store
drwxr-xr-x 1 drag drag   38 May  4 12:51 tree
```
[1692814922674023] Response end
---
[1692814952540215] command
```shell
ls -al
```

[1692814952540245] Response begin
```shell
total 24
drwxr-xr-x 1 drag drag  162 Aug 23 23:52 .
drwxr-xr-x 1 drag drag   52 May 31 01:02 ..
-rw-r--r-- 1 drag drag    0 Aug 23 23:51 foo.bar
drwxr-xr-x 1 drag drag  144 Aug 23 01:52 .git
-rw-r--r-- 1 drag drag 1429 May 10 00:38 go.mod
-rw-r--r-- 1 drag drag 8386 May 10 00:38 go.sum
drwxr-xr-x 1 drag drag   86 May  4 12:52 graph
drwxr-xr-x 1 drag drag   40 May  4 12:36 heap
-rw-r--r-- 1 drag drag 1029 Aug 23 23:52 log.md
-rw-r--r-- 1 drag drag 3120 Aug 23 23:51 main.go
drwxr-xr-x 1 drag drag   38 May  4 12:26 queue
drwxr-xr-x 1 drag drag   18 Aug 18 22:56 runner
drwxr-xr-x 1 drag drag   68 May 20 22:33 serviceq
drwxr-xr-x 1 drag drag   12 May  4 12:55 set
drwxr-xr-x 1 drag drag   38 May  4 12:26 stack
drwxr-xr-x 1 drag drag   70 May  4 12:31 store
drwxr-xr-x 1 drag drag   38 May  4 12:51 tree
```
[1692814952542765] Response end
---
[1692910555013826] command
```shell
ls -al
```

[1692910555013865] Response begin
```shell
total 24
drwxr-xr-x 1 drag drag  162 Aug 23 23:52 .
drwxr-xr-x 1 drag drag   52 May 31 01:02 ..
-rw-r--r-- 1 drag drag    0 Aug 23 23:51 foo.bar
drwxr-xr-x 1 drag drag  144 Aug 23 01:52 .git
-rw-r--r-- 1 drag drag 1429 May 10 00:38 go.mod
-rw-r--r-- 1 drag drag 8386 May 10 00:38 go.sum
drwxr-xr-x 1 drag drag   86 May  4 12:52 graph
drwxr-xr-x 1 drag drag   40 May  4 12:36 heap
-rw-r--r-- 1 drag drag 1967 Aug 25 02:25 log.md
-rw-r--r-- 1 drag drag 2005 Aug 25 02:25 main.go
drwxr-xr-x 1 drag drag   38 May  4 12:26 queue
drwxr-xr-x 1 drag drag   18 Aug 18 22:56 runner
drwxr-xr-x 1 drag drag   68 May 20 22:33 serviceq
drwxr-xr-x 1 drag drag   12 May  4 12:55 set
drwxr-xr-x 1 drag drag   38 May  4 12:26 stack
drwxr-xr-x 1 drag drag   70 May  4 12:31 store
drwxr-xr-x 1 drag drag   38 May  4 12:51 tree
```
[1692910555016647] Response end
---
[1692910569139654] command
```shell
ls -al | grep main
```

[1692910569139656] Response begin
```shell
-rw-r--r-- 1 drag drag 2005 Aug 25 02:25 main.go
```
[1692910569143526] Response end
---
[1692910584551016] command
```shell
ls -al | grep main
```

[1692910584551040] Response begin
```shell
-rw-r--r-- 1 drag drag 2005 Aug 25 02:25 main.go
```
[1692910584554831] Response end
---
[1692910620822558] command
```shell
tree . | grep main
```

[1692910620822560] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692910620827050] Response end
---
[1692910660261006] command
```shell
tree . | grep main
```

[1692910660261079] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692910660264888] Response end
---
[1692910683717304] command
```shell
tree .
```

[1692910683717306] Response begin
```shell
.
├── foo.bar
├── go.mod
├── go.sum
├── graph
│   ├── dir-graph.go
│   ├── undir-graph.go
│   └── weighted-graph.go
├── heap
│   ├── maxheap.go
│   └── minheap.go
├── log.md
├── main.go
├── queue
│   ├── linklist.go
│   └── redis.go
├── runner
│   └── runner.go
├── serviceq
│   ├── doc
│   │   ├── serviceQ.postman_collection.json
│   │   └── svcQ-state-diagram.png
│   ├── main.go.bak
│   ├── README.md
│   └── serviceq.go
├── set
│   └── set.go
├── stack
│   ├── linklist.go
│   └── redis.go
├── store
│   ├── hashmap-expiry.go
│   ├── hashmap.go
│   └── redis.go
└── tree
    ├── btree.go
    └── quadtree.go

11 directories, 26 files
```
[1692910683720222] Response end
---
[1692910743430555] command
```shell
tree .
```

[1692910743430558] Response begin
```shell
.
├── foo.bar
├── go.mod
├── go.sum
├── graph
│   ├── dir-graph.go
│   ├── undir-graph.go
│   └── weighted-graph.go
├── heap
│   ├── maxheap.go
│   └── minheap.go
├── log.md
├── main.go
├── queue
│   ├── linklist.go
│   └── redis.go
├── runner
│   └── runner.go
├── serviceq
│   ├── doc
│   │   ├── serviceQ.postman_collection.json
│   │   └── svcQ-state-diagram.png
│   ├── main.go.bak
│   ├── README.md
│   └── serviceq.go
├── set
│   └── set.go
├── stack
│   ├── linklist.go
│   └── redis.go
├── store
│   ├── hashmap-expiry.go
│   ├── hashmap.go
│   └── redis.go
└── tree
    ├── btree.go
    └── quadtree.go

11 directories, 26 files
```
[1692910743434328] Response end
---
[1692910766147189] command
```shell
tree . |grep main
```

[1692910766147876] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692910766151332] Response end
---
[1692910773900166] command
```shell
tree . |grep main
```

[1692910773900168] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692910773903234] Response end
---
[1692910991773991] command
```shell
tree . |grep main
```

[1692910991773992] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692910991777751] Response end
---
[1692911250191957] command
```shell
tree . |grep main
```

[1692911250191958] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692911250196201] Response end
---
[1692911338293444] command
```shell
tree . |grep main
```

[1692911338293446] Response begin
```shell
├── main.go
│   ├── main.go.bak
```
[1692911338296799] Response end
---
[1692911573140859] command
```shell
ls  |grep main
```

[1692911573140873] Response begin
```shell
main.go
```
[1692911573143086] Response end
---
[1692912221349412] command
```shell
ping google.com -c 4
```

[1692912221349433] Response begin
```shell
PING google.com (142.250.193.142) 56(84) bytes of data.
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=1 ttl=117 time=62.8 ms
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=2 ttl=117 time=11.9 ms
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=3 ttl=117 time=12.1 ms
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=4 ttl=117 time=12.1 ms

--- google.com ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3002ms
rtt min/avg/max/mdev = 11.859/24.723/62.835/22.004 ms
```
[1692912224452351] Response end
---
[1692912291348492] command
```shell
ping google.com -c 4
```

[1692912291348524] Response begin
```shell
PING google.com (142.250.193.142) 56(84) bytes of data.
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=1 ttl=117 time=11.5 ms
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=2 ttl=117 time=14.8 ms
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=3 ttl=117 time=15.0 ms
64 bytes from maa05s25-in-f14.1e100.net (142.250.193.142): icmp_seq=4 ttl=117 time=11.8 ms

--- google.com ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3003ms
rtt min/avg/max/mdev = 11.509/13.291/15.001/1.626 ms
```
[1692912294366088] Response end
---
[1692944459559902] command
```shell
ping google.com -c 4
```

[1692944459559922] Response begin
```shell
PING google.com (172.217.31.206) 56(84) bytes of data.
64 bytes from maa03s28-in-f14.1e100.net (172.217.31.206): icmp_seq=1 ttl=114 time=18.5 ms
64 bytes from maa03s28-in-f14.1e100.net (172.217.31.206): icmp_seq=2 ttl=114 time=18.0 ms
64 bytes from maa03s28-in-f14.1e100.net (172.217.31.206): icmp_seq=3 ttl=114 time=17.5 ms
64 bytes from maa03s28-in-f14.1e100.net (172.217.31.206): icmp_seq=4 ttl=114 time=18.6 ms

--- google.com ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 17.454/18.125/18.629/0.457 ms
```
[1692944462784894] Response end
---
[1692944481781438] command
```shell
ls  |grep main
```

[1692944481781457] Response begin
```shell
main.go
```
[1692944481787794] Response end
---
[1692944728911959] command
```shell
ls  |bla main
```

[1692944728911974] Response begin
```shell
```
[1692944728913888] Response end
---
[1692944996264869] command
```shell
ls  |bla main
```

[1692944996264885] Response begin
```shell
```
[1692944996266561] Response end
---
[1692945014157651] command
```shell
ls  |grep main
```

[1692945014157713] Response begin
```shell
main.go
```
[1692945014159718] Response end
---
[1692945045043489] command
```shell
ls  |bla main
```

[1692945045043490] Response begin
```shell
```
[1692945045045311] Response end
---
[1692945124890113] command
```shell
ls  |grep main
```

[1692945124890115] Response begin
```shell
main.go
```
[1692945124892185] Response end
---
[1692945167731671] command
```shell
ls  |grep main
```

[1692945167731679] Response begin
```shell
main.go
```
[1692945167734185] Response end
---
