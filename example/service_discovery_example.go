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
			fmt.Printf("node:%s, ip=%s ext=%s\n", k, v.(*etcd.Node).Meta.IP, v.(*etcd.Node).Meta.Ext)
			return true
		})
		time.Sleep(time.Second * 5)
	}
}
