package utils

import (
	"errors"
	"fmt"
	"time"

	// "github.com/garyburd/redigo/redis"

	"github.com/FZambia/sentinel"
	// "github.com/gomodule/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

type SentinelConfig struct {
	MasterName string   `json:"masterName"`
	Sentinels  []string `json:"sentinels"`
}

var RedisMaster string
var RedisSentinels []string
var RedisPassword string
var RedisDB string
var redisPool *redis.Pool

//KEY_TIME_TO_LIVE 键值对过期时间
const KEY_TIME_TO_LIVE = 120

func MonitorInit(master string, sentinels []string, password string, redisdb string) {
	RedisMaster = master
	RedisSentinels = sentinels
	RedisPassword = password
	RedisDB = redisdb
	if len(sentinels) == 1 {
		redisPool = newRedisPool()
	} else {
		redisPool = newSentinelPool()
	}
}

func newSentinelPool() *redis.Pool {
	sntnl := &sentinel.Sentinel{
		Addrs:      RedisSentinels,
		MasterName: RedisMaster,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   64,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr, redis.DialDatabase(1))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		},
	}
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   64,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr := RedisSentinels
			c, err := redis.Dial("tcp", masterAddr[0])
			if err != nil {
				return nil, err
			}
			//验证redis密码
			if _, err := c.Do("AUTH", RedisPassword); err != nil {
				c.Close()
				return nil, err
			}

			if _, err := c.Do("SELECT", RedisDB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return nil
			}
			return nil
		},
	}
}

func MonitorAdd(key string, count int) {
	conn := redisPool.Get()
	defer conn.Close()
	now := time.Now().Unix()
	now = now / 60 * 60
	_, err := conn.Do("HINCRBY", "H_MONITOR_"+key, now, count)
	if err != nil {
		fmt.Println(err)
	}
}
