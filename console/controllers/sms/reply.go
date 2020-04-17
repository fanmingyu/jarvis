package sms

import(
	"net/http"
	"strconv"
	"strings"

	"smsgate/utils"
	"smsgate/console/modules/db"
	"smsgate/report/modules/reply"
)

func Reply(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	page, _ := strconv.Atoi(vars.Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(vars.Get("pageSize"))
	if pageSize < 1 || pageSize > 1000 {
		pageSize = 100
	}

	mobile := strings.TrimSpace(vars.Get("mobile"))

	starttime, endtime := utils.GetUrlTime(vars, "")

	var data []reply.SmsReply
	err := db.RMysql.Engine.Desc("id").Where("create_time BETWEEN ? AND ?", starttime, endtime).Limit(pageSize, pageSize*(page - 1)).Find(&data, reply.SmsReply{Mobile: mobile})
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	for k, v := range data {
		data[k].Mobile = utils.MobileFormat(v.Mobile)
	}

	previous := page - 1
	next := page + 1
	if len(data) < pageSize {
		next = page
	}

	info := make(map[string]interface{})

	info["data"] = data
	info["page"] = page
	info["pageSize"] = pageSize
	info["mobile"] = mobile
	info["next"] = next
	info["previous"] = previous
	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")

	utils.View.Execute(w, r, "sms/reply.html", info)
}
