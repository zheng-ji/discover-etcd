package etcd

import (
	"testing"
	//"time"
)

func TestRegister(t *testing.T) {
	serviceName := "s-test"
	serviceInfo := ServiceMeta{IP: "127.0.0.1", Ext: "test"}
	_, err := Register(serviceName, serviceInfo, []string{
		"http://127.0.0.1:12379",
		"http://127.0.0.1:22379",
		"http://127.0.0.1:32379",
	})

	if err != nil {
		t.Errorf("err:%s", err.Error())
	}

}
