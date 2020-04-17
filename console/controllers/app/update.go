package app

import(
	"net/http"
	"strconv"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Update 修改一个app
func Update(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	info := make(map[string]interface{})

	idStr := vars.Get("id")
	if idStr == "" {
		utils.ResponseJson(w, -1, "id is invaild", nil)
		return
	}
	id, _ := strconv.Atoi(idStr)

	apps := make([]storage.SmsApp, 0)
	err := db.RMysql.Engine.Where("(id) = ?", id).Find(&apps)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	chans := getChannels()
	workers := getWorkers()
	info["channels"] = chans
	info["workers"] = workers
	info["data"] = apps[0]

	utils.View.Execute(w, r, "app/edit.html", info)
}
