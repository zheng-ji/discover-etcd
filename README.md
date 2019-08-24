Introduction
=======

discover-etcd implement service discovery based on etcd

Quick Start Etcd
-----------------

```
$ wget https://github.com/coreos/etcd/releases/download/v3.1.5/etcd-v3.1.5-linux-amd64.tar.gz
$ tar xzvf etcd-v3.1.5-linux-amd64.tar.gz
$ mv etcd-v3.1.5-linux-amd64 /opt/etcd

$ go get github.com/mattn/goreman
$ goreman -f ProcFile<本仓库提供> start
```


Quick-start
-----------------

```
go get github.com/zheng-ji/discover-etcd
```

```
# 服务注册端

```
example/service_register_example.go
``

```

# 服务发现端

```
example/service_discovery_example.go
``
