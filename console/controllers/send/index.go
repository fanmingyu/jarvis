package send

import(
	"net/http"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Index 发送测试短信的主页面
func Index(w http.ResponseWriter, r *http.Request) {
	apps := make([]storage.SmsApp, 0)
	err := db.RMysql.Engine.Find(&apps)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	info := map[string]interface{} {
		"data": apps,
	}
	utils.View.Execute(w, r, "send/index.html", info)
}
