package report

import (
	"net/http"
	"log"

	"smsgate/utils"
	"smsgate/report/modules/report"
	"smsgate/server/channels/welink"
)

func Welink(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Printf("welink report. report data:%v", r.PostForm)
	msgid := r.PostFormValue("MsgID")
	if msgid == "" {
		utils.ResponseJson(w, -1, "msgid is invaild", nil)
		return
	}

	mobile := r.PostFormValue("MobilePhone")
	if mobile == "" {
		utils.ResponseJson(w, -1, "mobile is invaild", nil)
		return
	}

	status := r.PostFormValue("ReportResultInfo")
	spNumber := r.PostFormValue("SPNumber")
	reportState := r.PostFormValue("ReportState")
	timeStr := r.PostFormValue("ReportTime")

	recvTime := int(utils.GetUnixTime(timeStr, "2006-01-02 15:04:05"))

	if welink.IsVerify(spNumber) && reportState == "True" {
		utils.MonitorAdd("WELINK_VERIFY_ACCEPT_SUCCESS", 1)
	}

	data := report.Report{
		Msgid:  msgid,
		Mobile: mobile,
		Status: status,
		Time:   recvTime,
	}
	log.Printf("welink report. report data is %v", data)

	report.Queue <- data

	utils.ResponseJson(w, 0, "success", nil)
}
