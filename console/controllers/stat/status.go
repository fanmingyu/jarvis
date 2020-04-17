package stat

import (
	"net/http"
	"sort"

	"smsgate/console/modules/status"
	"smsgate/utils"
	"smsgate/console/modules/db"
)

//Status 失败状态统计
func Status(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	starttime, endtime := utils.GetUrlTime(vars, "")

	var data []struct {
		Report status.ReportStatus
		Status string
		Num    int
		Sum    int
	}
	db.RMysql.Engine.Table("sms_record").Where("report_status!=? AND report_status!=? AND report_status!=? AND create_time BETWEEN ? AND ?", "__COMMIT", "DELIVRD", "0", starttime, endtime).Select("report_status as status, sum(num) as num").GroupBy("status").Find(&data)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Num > data[j].Num
	})

	sum := 0
	for i := 0; i < len(data); i++ {
		data[i].Report.Info = status.Status[data[i].Status].Info
		data[i].Report.Net = status.Status[data[i].Status].Net
		data[i].Report.Solution = status.Status[data[i].Status].Solution
		sum += data[i].Num
		data[i].Sum = sum
	}

	info := make(map[string]interface{})

	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["data"] = data
	info["sum"] = sum

	utils.View.Execute(w, r, "stat/status.html", info)
}
