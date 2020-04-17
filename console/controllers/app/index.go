package app

import(
	"net/http"
	"log"
	"time"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Index app页面显示主页
func Index(w http.ResponseWriter, r *http.Request) {
	apps := make([]storage.SmsApp, 0)
	err := db.RMysql.Engine.Find(&apps)
	if err != nil {
		log.Printf("find apps failed. err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	data := processApp(apps)
	utils.View.Execute(w, r, "app/index.html", data)
}

func processApp(apps []storage.SmsApp) []map[string]interface{} {
	datas := make([]map[string]interface{}, 0)

	for _, app := range apps {
		data := make(map[string]interface{})

		data["name"] = app.Name
		data["secret"] = app.Secret
		data["prefix"] = app.Prefix
		data["channel"] = app.Channel
		data["worker"] = app.Worker
		data["id"] = app.Id

		if app.CreateTime != 0 {
			data["createTime"] = time.Unix(int64(app.CreateTime), 0).Format("2006-01-02 15:04:05")
		}

		if app.UpdateTime != 0 {
			data["updateTime"] = time.Unix(int64(app.UpdateTime), 0).Format("2006-01-02 15:04:05")
		}

		datas = append(datas, data)
	}
	return datas
}
