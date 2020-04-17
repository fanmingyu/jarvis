package utils

import (
	"encoding/json"
	"io/ioutil"
)

var Config = conf{}

type conf struct{}

// 从Json字符串解析出配置
func (c conf) ParseString(content []byte, result interface{}) {
	err := json.Unmarshal(content, &result)
	if err != nil {
		panic("json unmarshal failed. err:" + err.Error())
	}
}

// 从Json配置文件解析出配置
func (c conf) ParseFile(file string, result interface{}) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic("read file failed. err:" + err.Error())
	}

	c.ParseString(content, &result)
}
