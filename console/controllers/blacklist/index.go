package blacklist

import (
	"net/http"
	"time"
	"strconv"
	"strings"

	"smsgate/server/storage"
	"smsgate/utils"
	"smsgate/console/modules/db"
)

const pageSize = 100

//Index 黑名单显示主页
func Index(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	info := make(map[string]interface{})

	mobile := strings.TrimSpace(vars.Get("mobile"))
	page, _ := strconv.Atoi(vars.Get("page"))
	blacklists := make([]storage.SmsBlacklist, 0)

	err := db.RMysql.Engine.Desc("id").Limit(pageSize, pageSize*page).Find(&blacklists, storage.SmsBlacklist{Mobile: mobile})
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	previous := page - 1
	if previous < 0 {
		previous = 0
	}

	next := page + 1
	if len(blacklists) < pageSize {
		next = page
	}
	data := processBlacklist(blacklists)

	info["page"] = page
	info["data"] = data
	info["mobile"] = mobile
	info["next"] = next
	info["previous"] = previous

	utils.View.Execute(w, r, "blacklist/index.html", info)
}

func processBlacklist(blacklists []storage.SmsBlacklist) []map[string]interface{} {
	datas := make([]map[string]interface{}, 0)

	for _, blacklist := range blacklists {
		data := make(map[string]interface{})

		data["id"] = blacklist.Id
		data["mobile"] = utils.MobileFormat(blacklist.Mobile)
		data["createTime"] = time.Unix(int64(blacklist.CreateTime), 0).Format("2006-01-02 15:04:05")

		datas = append(datas, data)
	}
	return datas
}
