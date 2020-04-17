package stat

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"smsgate/console/modules/db"
	"smsgate/console/modules/status"
	"smsgate/utils"
)

//Index 首页
func Index(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	start := vars.Get("start")
	if start == "" {
		start = time.Now().Add(-1 * time.Hour).Format("2006-01-02 15:04")
	}

	end := vars.Get("end")
	if end == "" {
		end = time.Now().Add(1 * time.Minute).Format("2006-01-02 15:04")
	}

	starttime := utils.GetUnixTime(start, "2006-01-02 15:04")
	endtime := utils.GetUnixTime(end, "2006-01-02 15:04") + 59

	var data []struct {
		Status string
		Num    int
		Cost   float64
	}
	db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("report_status as status, sum(num) as num, avg(if(report_time>create_time, report_time-create_time, 0)) AS cost").GroupBy("status").Find(&data)

	sum := 0
	cost := 0.0
	statData := map[string]int{}
	for _, v := range data {
		title := status.TitleFull.GetTitle(v.Status)
		statData[title] += v.Num

		if title == status.TitleFull["DELIVRD"] {
			cost = v.Cost
		}

		sum += v.Num
	}

	chartData := [][]interface{}{}
	for k, v := range statData {
		chartData = append(chartData, []interface{}{k, v})
	}
	sort.Slice(chartData, func(i, j int) bool {
		return chartData[i][1].(int) > chartData[j][1].(int)
	})
	pieData, _ := json.Marshal(chartData)

	info := map[string]interface{}{}
	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["statusTitle"] = status.TitleFull.GetTitles()
	info["data"] = statData
	info["pieData"] = string(pieData)
	info["sum"] = sum
	info["cost"] = cost

	utils.View.Execute(w, r, "stat/index.html", info)
}
