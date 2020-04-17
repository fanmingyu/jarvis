package blacklist

import(
	"net/http"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Delete 删除一个黑名单
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	id := vars.Get("id")
	if id == "" {
		utils.ResponseJson(w, -1,"id is invaild", nil)
		return
	}

	blacklist := new(storage.SmsBlacklist)
	_, err := db.WMysql.Engine.Where("id = ?", id).Delete(blacklist)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/blacklist/index", 302)
}
