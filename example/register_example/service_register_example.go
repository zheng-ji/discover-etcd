package main

import (
	"fmt"
	"github.com/zheng-ji/discover-etcd"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := etcd.ServiceMeta{IP: "127.0.0.1", Ext: "test"}
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
