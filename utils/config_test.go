package utils

import (
	"io/ioutil"
	"testing"
)

type MyConf struct {
	LogPath       string            `json:"logPath"`
	RedisSentinel RedisSentinelConf `json:"redis_sentinel"`
	WorkerCount   int               `json:"worker_count"`
}

type RedisSentinelConf struct {
	MasterName string   `json:"masterName"`
	Sentinels  []string `json:"sentinels"`
}

var tmpfile = "/tmp/test-config-sample"

var confSample = `{
    "logPath":	"./logs/",
    "worker_count" : 10,
    "redis_sentinel": {
        "masterName": "def_master",
        "sentinels":  ["127.0.0.1:26479", "127.0.0.1:26579", "127.0.0.1:26679"]
    }
}`

func TestParseString(t *testing.T) {
	var conf MyConf
	Config.ParseString([]byte(confSample), &conf)
	t.Logf("conf string = %v", confSample)
	t.Logf("conf result = %+v", conf)

	if conf.WorkerCount != 10 {
		t.Errorf("test int failed")
	}

	if conf.LogPath != "./logs/" {
		t.Errorf("test string failed")
	}

	if conf.RedisSentinel.MasterName != "def_master" {
		t.Errorf("test struct failed")
	}

	if len(conf.RedisSentinel.Sentinels) != 3 {
		t.Errorf("test slice failed")
	}
}

func TestParseStringFail(t *testing.T) {
	wrongConf := confSample + ","

	defer func() {
		err := recover()
		t.Logf("has panic. err:%v", err)
	}()

	var conf MyConf
	Config.ParseString([]byte(wrongConf), &conf)

	// has not panic
	t.Errorf("test parse fail failed")
}

func TestParseFile(t *testing.T) {
	t.Logf("tmpfile = %v", tmpfile)

	ioutil.WriteFile(tmpfile, []byte(confSample), 0644)

	var conf MyConf
	Config.ParseFile(tmpfile, &conf)

	if conf.WorkerCount != 10 {
		t.Errorf("test int failed")
	}
}

func TestParseFileFail(t *testing.T) {
	notExistsFile := "/tmp/12345"

	defer func() {
		err := recover()
		t.Logf("has panic. err:%v", err)
	}()

	var conf MyConf
	Config.ParseFile(notExistsFile, &conf)

	// has not panic
	t.Errorf("test parse fail failed")
}
