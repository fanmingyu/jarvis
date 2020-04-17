package tpl

import(
	"net/http"
	"strconv"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Update 更新一个模板
func Update(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	idStr := vars.Get("id")
	if idStr == "" {
		utils.ResponseJson(w, -1, "id is invaild", nil)
		return
	}

	id, _ := strconv.Atoi(idStr)
	tpls := make([]storage.SmsTpl, 0)

	err := db.RMysql.Engine.Where("(id) = ?", id).Find(&tpls)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.View.Execute(w, r, "tpl/edit.html", tpls[0])
}
