package tpl

import(
	"net/http"
	"strconv"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Delete 删除一个模板
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	idStr := vars.Get("id")
	if idStr == "" {
		utils.ResponseJson(w, -1, "id is invaild", nil)
		return
	}
	id, _ := strconv.Atoi(idStr)
	tpl := storage.SmsTpl{}

	_, err := db.WMysql.Engine.Where("(id) = ?", id).Delete(&tpl)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}
	http.Redirect(w, r, "/tpl/index", 302)
}
