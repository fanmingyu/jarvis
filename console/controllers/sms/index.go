package sms

import (
	"net/http"
	"strconv"
	"strings"

	"smsgate/server/storage"
	"smsgate/utils"
	"smsgate/console/modules/db"
)

//Index 主页显示
func Index(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	info := make(map[string]interface{})
	page, _ := strconv.Atoi(vars.Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(vars.Get("pageSize"))
	if pageSize < 1 || pageSize > 1000 {
		pageSize = 100
	}

	mobile := strings.TrimSpace(vars.Get("mobile"))
	status := strings.TrimSpace(vars.Get("status"))

	starttime, endtime := utils.GetUrlTime(vars, "")

	sms := make([]storage.SmsRecord, 0)
	err := db.RMysql.Engine.Desc("id").Where("create_time BETWEEN ? AND ?", starttime, endtime).Limit(pageSize, pageSize*(page - 1)).Find(&sms, storage.SmsRecord{Mobile: mobile, ReportStatus: status})

	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	previous := page - 1
	next := page + 1
	if len(sms) < pageSize {
		next = page
	}

	data := processSms(sms)

	info["data"] = data
	info["page"] = page
	info["pageSize"] = pageSize
	info["mobile"] = mobile
	info["next"] = next
	info["previous"] = previous
	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["status"] = status

	utils.View.Execute(w, r, "sms/index.html", info)
}

func processSms(smsRecords []storage.SmsRecord) []map[string]interface{} {
	datas := make([]map[string]interface{}, 0)

	for _, sms := range smsRecords {
		data := make(map[string]interface{})

		data["mobile"] = utils.MobileFormat(sms.Mobile)
		data["app"] = sms.App
		data["content"] = sms.Content
		data["count"] = sms.Num
		data["channel"] = sms.Channel
		data["tplId"] = sms.TplId
		data["createTime"]  = sms.CreateTime
		data["requestTime"]  = sms.RequestTime
		data["reportTime"]  = sms.ReportTime
		data["reportRecvTime"]  = sms.ReportRecvTime

		data["msgid"] = sms.Msgid
		data["reportStatus"] = sms.ReportStatus
		if sms.RequestStatus == 1 {
			data["status"] = "已保存"
		} else if sms.RequestStatus == 2 {
			data["status"] = "提交失败"
		} else if sms.RequestStatus == 3 {
			data["status"] = "提交成功"
		}

		datas = append(datas, data)
	}

	return datas
}
