package registry

import (
	"testing"
	"time"
)

var (
	etcdEndpoints = []string{"127.0.0.1:2379"}
	prefix        = "/xxxx-servers-test/"
	reg           *Registry
)

func TestRegister(t *testing.T) {
	var err error
	reg, err = NewRegistry(etcdEndpoints, prefix, 10*time.Second)
	if err != nil {
		t.Errorf("new register error. err:%v", err)
	}

	node := &Node{
		IP:   "127.0.0.1",
		Port: "20000",
	}
	reg.RegisterNode(node, 10*time.Second)
}

func TestWatch(t *testing.T) {
	reg.WatchNodes()
	time.Sleep(1 * time.Second)

	nodes := reg.GetNodes()
	t.Logf("get nodes. nodes:%v", nodes)
	if len(nodes) != 1 || nodes[0].IP != "127.0.0.1" {
		t.Errorf("get nodes failed. nodes:%v", nodes)
	}
}
