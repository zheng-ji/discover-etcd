Introduction
=======

[![Build Status](https://travis-ci.org/zheng-ji/discover-etcd.svg)](https://travis-ci.org/zheng-ji/discover-etcd)

Implement service discovery based on etcd By Go

Quick Start Etcd
-----------------

```
$ wget https://github.com/coreos/etcd/releases/download/v3.1.5/etcd-v3.1.5-linux-amd64.tar.gz
$ tar xzvf etcd-v3.1.5-linux-amd64.tar.gz
$ mv etcd-v3.1.5-linux-amd64 /opt/etcd

$ go get github.com/mattn/goreman
$ goreman -f ProcFile (本仓库提供) start
```


Quick-start
-----------------

```
go get github.com/zheng-ji/discover-etcd
```

Example
----------------

* [Service Discovery Link](https://github.com/zheng-ji/discover-etcd/blob/master/example/discovery_example/service_discovery_example.go)
* [Service Register Link](https://github.com/zheng-ji/discover-etcd/blob/master/example/register_example/service_register_example.go)

```go
# 服务发现
package main

import (
	"fmt"
	"github.com/zheng-ji/discover-etcd"
	"log"
	"time"
)

func main() {
	m, err := etcd.OnWatch([]string{
		"http://127.0.0.1:12379",
		"http://127.0.0.1:22379",
		"http://127.0.0.1:32379"}, "services/")
	if err != nil {
		log.Fatal(err)
	}

	for {
		m.Nodes.Range(func(k, v interface{}) bool {
			fmt.Printf("node:%s, ip=%s ext=%s\n",
                k, v.(*etcd.Node).Meta.IP, v.(*etcd.Node).Meta.Ext)
			return true
		})
		time.Sleep(time.Second * 5)
	}
}
```

```go
# 服务注册
package main

import (
	"fmt"
	"github.com/zheng-ji/discover-etcd"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := etcd.ServiceMeta{IP: "127.0.0.1", Ext: "Your Custome Data"} 
	s, err := etcd.Register(serviceName, serviceInfo, []string{
		"http://127.0.0.1:12379",
		"http://127.0.0.1:22379",
		"http://127.0.0.1:32379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Meta.IP)

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}
```
