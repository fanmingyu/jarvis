package sms

import(
	"net/http"
	"strconv"

	"smsgate/utils"
	"smsgate/server/storage"
	"smsgate/report/modules/reply"
)

func Reply(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	account := r.Form["account"]
	if len(account) < 1 {
		utils.ResponseJson(w, -1, "account is invaild", nil)
		return
	}

	time := r.Form["time"]
	if len(time) < 1 {
		utils.ResponseJson(w, -1, "time is invaild", nil)
		return
	}
	timeUnix, _ := strconv.Atoi(time[0])
	if timeUnix < 0 {
		utils.ResponseJson(w, -1, "time is illegal", nil)
		return
	}

	var data []reply.SmsReply
	err := storage.Mysql.Engine.Where("account=? AND create_time>=?", account[0], timeUnix).Find(&data)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseJson(w, 0, "get reply success", data)
}
