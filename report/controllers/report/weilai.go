package report

import (
	"net/http"
	"strings"
	"log"

	"smsgate/report/modules/report"
	"smsgate/utils"
	"smsgate/server/channels/weilai"
)

func Weilai(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	reportStr := r.PostFormValue("report")
	log.Printf("weilai report. report data:%v", reportStr)
	if reportStr == "" {
		utils.ResponseJson(w, -1, "report is invaild", nil)
		return
	}

	reports := strings.Split(reportStr, "^")
	for _, r := range reports {
		value := strings.Split(r, ",")
		if len(value) < 5 {
			break
		}

		if weilai.IsVerify(value[4]) && (value[2] == "DELIVRD" || value[2] == "0") {
			utils.MonitorAdd("WEILAI_VERIFY_ACCEPT_SUCCESS", 1)
		}

		reportTime := utils.GetUnixTime(value[3], "2006-01-02 15:04:05")
		data := report.Report{
			Msgid:  value[0],
			Mobile: value[1],
			Status: value[2],
			Time:   int(reportTime),
		}
		report.Queue <- data
	}
	utils.ResponseJson(w, 0, "success", nil)
}
