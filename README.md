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
