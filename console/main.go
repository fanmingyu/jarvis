package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	"smsgate/console/controllers"
	"smsgate/console/modules/db"
	"smsgate/console/modules/node"
	"smsgate/console/modules/status"
	"smsgate/utils"
	"smsgate/utils/registry"
)

//Config 结构体来接收配置文件中的参数数据
type Config struct {
	LogPath  string
	RMysql   mysql.Config
	WMysql   mysql.Config
	Port     string
	EtcdConf registry.EtcdConf
}

var routers = controllers.Controllers
var conf Config
var funcs = controllers.ViewFuncs

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error: the config file is invaild.")
		os.Exit(-1)
	}

	utils.Config.ParseFile(os.Args[1], &conf)

	if !utils.IsDir(conf.LogPath) {
		fmt.Printf("Error: the log path is invaild.")
		os.Exit(-1)
	}

	utils.LoggerInit(conf.LogPath)

	//初始化MySQL
	db.RMysql.Init(&conf.RMysql)
	db.WMysql.Init(&conf.WMysql)

	//初始化view
	utils.View.Init()
	utils.View.SetFuncs(funcs)

	//初始化服务注册器registry
	node.InitRegistry(conf.EtcdConf)

	//加载状态报告资源
	status.Init()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	for k, v := range routers {
		http.HandleFunc(k, middleWare(v))
	}

	fmt.Printf("sms console start. port:%v", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}

func middleWare(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authgateUsername := r.Header["Authgate-Username"]
		username := "anonymous"
		if len(authgateUsername) > 0 {
			username = authgateUsername[0]
		}
		ctx := context.WithValue(r.Context(), "username", username)
		start := time.Now()
		ctx = context.WithValue(ctx, "start", start)

		r = r.WithContext(ctx)

		handler(w, r)
		cost := time.Now().Sub(start)
		log.Printf("http request. cost:%v, method:%v, uri:%v, host:%v, ip:%v", cost, r.Method, r.RequestURI, r.Host, r.RemoteAddr)
	}
}
