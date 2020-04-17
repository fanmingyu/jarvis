package blacklist

import(
	"net/http"
	"strings"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/console/modules/db"
)

//Edit 编辑黑名单
func Edit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mobile := strings.TrimSpace(r.PostFormValue("mobile"))
	if mobile == "" {
		utils.ResponseJson(w, -1, "mobile is invaild.", nil)
		return
	}

	blacklist := storage.SmsBlacklist{
		Mobile: mobile,
	}
	_, err := db.WMysql.Engine.Insert(&blacklist)

	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/blacklist/index", 302)
}
