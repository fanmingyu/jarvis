package storage

import (
	"log"
	"sync"
	"time"
)

type SmsApp struct {
	Id         int `xorm:"notnull pk autoincr"`
	Name       string
	Secret     string
	Prefix     string
	Channel    string
	Worker     string
	CreateTime int `xorm:"created"`
	UpdateTime int `xorm:"updated"`
}

type SmsTpl struct {
	Id         int `xorm:"notnull pk autoincr"`
	OutId      string
	Name       string
	Content    string
	CreateTime int `xorm:"created"`
	UpdateTime int `xorm:"updated"`
}

type SmsBlacklist struct {
	Id         int `xorm:"notnull pk autoincr"`
	Mobile     string
	CreateTime int `xorm:"created"`
}

type Tpl struct {
	tpl  map[string]SmsTpl
	Lock sync.RWMutex
}

type App struct {
	app  map[string]SmsApp
	Lock sync.RWMutex
}

type Blacklist struct {
	blacklist map[string]int
	Lock      sync.RWMutex
}

type MemoryData struct {
	tpl       Tpl
	app       App
	blacklist Blacklist
}

var CacheData = &MemoryData{}

//初始化
func CacheInit() {
	err := CacheData.loadInfo()
	if err != nil {
		panic(err)
	}
	cacheCron()
}

func cacheCron() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			CacheData.loadInfo()
		}
	}()
}

//根据key获取各种数据——get方法
func (cacheData *MemoryData) GetTpl(key string) SmsTpl {
	cacheData.tpl.Lock.RLock()
	defer cacheData.tpl.Lock.RUnlock()
	return cacheData.tpl.tpl[key]
}

func (cacheData *MemoryData) GetApp(key string) SmsApp {
	cacheData.app.Lock.RLock()
	defer cacheData.app.Lock.RUnlock()
	return cacheData.app.app[key]
}

func (cacheData *MemoryData) GetBlacklist(key string) int {
	cacheData.blacklist.Lock.RLock()
	defer cacheData.blacklist.Lock.RUnlock()
	return cacheData.blacklist.blacklist[key]
}

func (cacheData *MemoryData) loadInfo() error {
	tpls := make([]SmsTpl, 0)
	err := Mysql.Engine.Find(&tpls)
	if err != nil {
		log.Printf("get tplInfo failed. err:%v", err)
		return err
	}
	log.Printf("get tplInfo success. count:%d", len(tpls))
	cacheData.SetTplsInfo(tpls)

	apps := make([]SmsApp, 0)
	err = Mysql.Engine.Find(&apps)
	if err != nil {
		log.Printf("get appInfo failed. err:%v", err)
		return err
	}
	log.Printf("get appInfo success. count:%d", len(apps))
	cacheData.SetAppInfo(apps)

	list := make([]SmsBlacklist, 0)
	err = Mysql.Engine.Find(&list)
	if err != nil {
		log.Printf("get blacklist failed. err:%v", err)
		return err
	}
	log.Printf("get blacklist success. count:%d", len(list))
	cacheData.SetBlacklist(list)
	return nil
}

//set方法
func (cacheData *MemoryData) SetTplsInfo(tpls []SmsTpl) {
	cacheData.tpl.Lock.Lock()
	defer cacheData.tpl.Lock.Unlock()
	cacheData.tpl.tpl = map[string]SmsTpl{}

	for _, tpl := range tpls {
		cacheData.tpl.tpl[tpl.Name] = tpl
	}
}

func (cacheData *MemoryData) SetAppInfo(apps []SmsApp) {
	cacheData.app.Lock.Lock()
	defer cacheData.app.Lock.Unlock()
	cacheData.app.app = map[string]SmsApp{}

	for _, app := range apps {
		cacheData.app.app[app.Name] = app
	}
}

func (cacheData *MemoryData) SetBlacklist(list []SmsBlacklist) {
	cacheData.blacklist.Lock.Lock()
	defer cacheData.blacklist.Lock.Unlock()
	cacheData.blacklist.blacklist = map[string]int{}

	for _, member := range list {
		cacheData.blacklist.blacklist[member.Mobile] = 1
	}
}
