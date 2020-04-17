package controllers

import(
	"testing"
)

func TestViewFuncs(t *testing.T) {
	s := getChannelName("testfa")

	if s != "" {
		t.Errorf("test get name failed.")
	}

	s = getChannelName("test")
	if s != "测试通道" {
		t.Errorf("test get channel name failed")
	}

	s = getWorkerName("testf")
	if s != "" {
		t.Errorf("test get worker name failed. the result is %v", s)
	}

	s = getWorkerName("test")
	if s != "测试队列" {
		t.Errorf("test get worker name failed. the result is %v", s)
	}

	m := calcPercent(5, 7)
	if m != "71.43" {
		t.Errorf("test calcPercent failed. the result is %v", m)
	}
}
