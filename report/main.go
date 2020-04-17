package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	"smsgate/report/controllers"
	"smsgate/report/modules/db"
	"smsgate/report/modules/report"
	"smsgate/utils"
)

type Config struct {
	LogPath                 string
	Mysql                   mysql.Config
	Port                    string
	RedisSentinelMasterName string
	RedisSentinelHosts      []string
	RedisPassWord           string
	RedisDb                 string
}

var Conf Config
var routers = controllers.Controllers
var workerNumber = 20

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: the config is invaild.")
		os.Exit(-1)
	}

	utils.Config.ParseFile(os.Args[1], &Conf)

	if !utils.IsDir(Conf.LogPath) {
		fmt.Println("Error: the log path is invaild.")
		os.Exit(-1)
	}

	//初始化日志
	utils.LoggerInit(Conf.LogPath)

	//redis初始化
	utils.MonitorInit(Conf.RedisSentinelMasterName, Conf.RedisSentinelHosts, Conf.RedisPassWord, Conf.RedisDb)

	//初始化mysql
	db.Mysql.Init(&Conf.Mysql)

	//启动worker
	report.WorkerStart(workerNumber)

	for k, v := range routers {
		http.HandleFunc(k, func(f http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				s := time.Now()
				f(w, r)
				cost := time.Now().Sub(s)

				log.Printf("http request. cost:%v, method:%v, uri:%v, host:%v, ip:%v", cost, r.Method, r.RequestURI, r.Host, r.RemoteAddr)
			}
		}(v))
	}

	fmt.Printf("sms report start. port:%v", Conf.Port)
	log.Printf("sms report startup. port:%v", Conf.Port)
	log.Fatal(http.ListenAndServe(":"+Conf.Port, nil))
}
