package tpl

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"smsgate/console/modules/db"
	"smsgate/server/storage"
	"smsgate/utils"
)

const pageSize = 100

//Index 模板主页显示
func Index(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	info := make(map[string]interface{})

	keyword := strings.TrimSpace(vars.Get("keyword"))
	page, _ := strconv.Atoi(vars.Get("page"))
	tpls := make([]storage.SmsTpl, 0)
	id, _ := strconv.Atoi(vars.Get("id"))

	err := db.RMysql.Engine.Where("name LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Asc("id").Limit(pageSize, pageSize*page).Find(&tpls, storage.SmsTpl{Id: id})
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}
	previous := page - 1
	if previous < 0 {
		previous = 0
	}

	next := page + 1
	if len(tpls) < pageSize {
		next = page
	}

	data := processTpl(tpls)

	info["data"] = data
	info["page"] = page
	info["next"] = next
	info["previous"] = previous
	info["keyword"] = keyword

	utils.View.Execute(w, r, "tpl/index.html", info)
}

func processTpl(tpls []storage.SmsTpl) []map[string]interface{} {
	datas := make([]map[string]interface{}, 0)

	for _, tpl := range tpls {
		data := make(map[string]interface{})

		data["id"] = tpl.Id
		data["out_id"] = tpl.OutId
		data["name"] = tpl.Name
		data["content"] = tpl.Content

		if tpl.CreateTime != 0 {
			data["createTime"] = time.Unix(int64(tpl.CreateTime), 0).Format("2006-01-02 15:04:05")
		}

		if tpl.UpdateTime != 0 {
			data["updateTime"] = time.Unix(int64(tpl.UpdateTime), 0).Format("2006-01-02 15:04:05")
		}

		datas = append(datas, data)
	}
	return datas
}
