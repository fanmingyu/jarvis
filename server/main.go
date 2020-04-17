package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"smsgate/server/controllers"
	"smsgate/server/filters"
	"smsgate/server/storage"
	"smsgate/utils"
	"smsgate/utils/registry"
)

var Routers = controllers.Routers

type SmsConfig struct {
	LogPath                 string
	RedisSentinelMasterName string
	RedisSentinelHosts      []string
	RedisPassWord           string
	RedisDb                 string
	Mysql                   mysql.Config
	Port                    string
	EtcdConf                registry.EtcdConf
}

var Conf SmsConfig

func main() {
	//必须要有配置文件
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

	//初始化监控
	utils.MonitorInit(Conf.RedisSentinelMasterName, Conf.RedisSentinelHosts, Conf.RedisPassWord, Conf.RedisDb)

	//初始化mysql
	storage.MysqlInit(&Conf.Mysql)

	//初始化内存中的存储信息
	storage.CacheInit()

	//服务注册
	r, err := registry.NewRegistry(Conf.EtcdConf)
	if err != nil {
		fmt.Printf("get registry failed. err:%v\n", err)
		os.Exit(-1)
	}

	ip, err := utils.GetLocalIp()
	if err != nil {
		fmt.Printf("get ip failed. err:%v\n", err)
		os.Exit(-1)
	}
	r.RegisterNode(&registry.Node{
		IP:           ip,
		Port:         Conf.Port,
		RegisterTime: time.Now().String(),
	}, 10*time.Second)

	//添加prometheus监控
	http.Handle("/metrics", promhttp.Handler())

	//添加filter功能
	//后期优化加载filter可以参考controller的加载方式，实现可扩展
	filters.Register("/**", func(rw http.ResponseWriter, r *http.Request) error {
		//暂时取消内网ip过滤功能
		// ip := utils.GetRemoteIp(r)
		// if utils.IsPublicIP(net.ParseIP(ip)) == false {
		// 	utils.ResponseJson(rw, -1, "please use public net ip", nil)
		// 	return errors.New("")
		// }
		return nil
	})

	//对外接口
	for k, v := range Routers {
		http.HandleFunc(k, func(f http.HandlerFunc) http.HandlerFunc {
			return filters.Handle(func(w http.ResponseWriter, r *http.Request) error {
				s := time.Now()
				f(w, r)
				cost := time.Now().Sub(s)

				log.Printf("http request. cost:%v, method:%v, uri:%v, host:%v, ip:%v, header:%v", cost, r.Method, r.RequestURI, r.Host, r.RemoteAddr, r.Header)
				return nil
			})
		}(v))
	}

	fmt.Printf("smsserver start. port:%v", Conf.Port)
	log.Printf("smsserver startup. port:%v", Conf.Port)
	log.Fatal(http.ListenAndServe(":"+Conf.Port, nil))
}
