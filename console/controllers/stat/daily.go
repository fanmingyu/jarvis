package stat

import (
	"net/http"
	"time"

	"smsgate/utils"
	"smsgate/console/modules/db"
	"smsgate/console/modules/status"
)

//Daily 按照日期来统计数据
func Daily(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var data []struct {
		Title  string
		Status string
		Num    int
	}
	info := make(map[string]interface{})

	starttime, endtime := utils.GetUrlTime(vars, "-48h")

	db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("report_status as status, sum(num) as num, from_unixtime(create_time, '%Y-%m-%d') as title").GroupBy("title, report_status").Find(&data)

	statData := make(map[string]map[string]int)
	for i := starttime; i < endtime; i += 86400 {
		date := time.Unix(i, 0).Format("2006-01-02")
		statData[date] = make(map[string]int)
		statData[date]["DELIVRD"] = 0
		statData[date]["__SUM"] = 0
	}

	for _, value := range data {
		title := status.TitleFull.GetTitle(value.Status)
		statData[value.Title][title] += value.Num
		statData[value.Title]["__SUM"] += value.Num
	}

	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["data"] = statData
	info["statusTitle"] = status.TitleFull.GetTitles()

	utils.View.Execute(w, r, "stat/daily.html", info)
}
