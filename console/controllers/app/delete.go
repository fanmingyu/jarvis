package app

import(
	"net/http"
	"strconv"

	"smsgate/utils"
	"smsgate/console/modules/db"
	"smsgate/server/storage"
)

//Delete 删除一个app
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	idStr := vars.Get("id")
	if idStr == "" {
		utils.ResponseJson(w, -1, "id is invaild", nil)
		return
	}
	id, _ := strconv.Atoi(idStr)

	app := new(storage.SmsApp)
	_, err := db.WMysql.Engine.Where("(id) = ?", id).Delete(app)

	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/app/index", 302)
}
