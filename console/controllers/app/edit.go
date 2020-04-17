package app

import (
	"net/http"
	"strconv"
	"strings"

	"smsgate/server/storage"
	"smsgate/console/modules/db"
	"smsgate/utils"
)

//Edit 编辑一个app
func Edit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := strings.TrimSpace(r.PostFormValue("name"))
	if name == "" {
		utils.ResponseJson(w, -1, "name is invaild", nil)
		return
	}

	secret := strings.TrimSpace(r.PostFormValue("secret"))
	if secret == "" {
		utils.ResponseJson(w, -1, "secret is invaild", nil)
		return
	}

	prefix := strings.TrimSpace(r.PostFormValue("prefix"))
	if prefix == "" {
		utils.ResponseJson(w, -1, "prefix is invaild", nil)
		return
	}

	channel := r.PostFormValue("channel")
	if channel == "" {
		utils.ResponseJson(w, -1, "channel is invaild", nil)
		return
	}

	worker := r.PostFormValue("worker")
	if worker == "" {
		utils.ResponseJson(w, -1, "worker is invaild", nil)
		return
	}

	idStr := r.PostFormValue("id")
	app := storage.SmsApp{
		Name:    name,
		Secret:  secret,
		Prefix:  prefix,
		Channel: channel,
		Worker:  worker,
	}

	err := Upsert(idStr, &app)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/app/index", 302)
}

//Upsert 修改一条数据，如果没有，那就插入
func Upsert(idStr string, app *storage.SmsApp) error {
	if idStr == "" {
		_, err := db.WMysql.Engine.Insert(app)
		return err
	}

	id, _ := strconv.Atoi(idStr)
	_, err := db.WMysql.Engine.Where("(id) = ?", id).Update(app)
	return err
}
