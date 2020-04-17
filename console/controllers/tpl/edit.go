package tpl

import (
	"net/http"
	"strconv"
	"strings"

	"smsgate/console/modules/db"
	"smsgate/server/storage"
	"smsgate/utils"
)

//Edit 编辑一个模板
func Edit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := strings.TrimSpace(r.PostFormValue("name"))
	if name == "" {
		utils.ResponseJson(w, -1, "name is invaild", nil)
		return
	}

	out_id := strings.TrimSpace(r.PostFormValue("out_id"))
	if out_id == "" {
		utils.ResponseJson(w, -1, "out_id is invaild", nil)
		return
	}

	content := strings.TrimSpace(r.PostFormValue("content"))
	if content == "" {
		utils.ResponseJson(w, -1, "content is invaild", nil)
		return
	}

	tpl := storage.SmsTpl{
		Name:    name,
		OutId:   out_id,
		Content: content,
	}

	idStr := r.PostFormValue("id")
	err := upsert(idStr, &tpl)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/tpl/index", 302)
}

//修改插入
func upsert(idStr string, tpl *storage.SmsTpl) error {
	if idStr == "" {
		_, err := db.WMysql.Engine.Insert(tpl)
		return err
	}

	id, _ := strconv.Atoi(idStr)
	_, err := db.WMysql.Engine.Where("(id) = ?", id).Update(tpl)

	return err
}
