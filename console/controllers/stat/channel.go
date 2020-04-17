package stat

import (
	"net/http"
	"encoding/json"
	"sort"

	"smsgate/console/modules/status"
	"smsgate/server/channels"
	"smsgate/utils"
	"smsgate/console/modules/db"
)

//Channel 按照通道来统计数据
func Channel(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var data []struct{
		Channel string
		Num     int
		Status  string
	}
	info := make(map[string]interface{})

	starttime, endtime := utils.GetUrlTime(vars, "")

	db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("channel, report_status as status, sum(num) as num").GroupBy("channel, status").Find(&data)

	statData := make(map[string]map[string]int)
	for channel := range channels.Channels {
		statData[channel] = make(map[string]int)
		statData[channel]["__SUM"] = 0
		statData[channel]["DELIVRD"] = 0
	}

	for _, value := range data {
		title := status.Title.GetTitle(value.Status)
		statData[value.Channel][title] += value.Num
		statData[value.Channel]["__SUM"] += value.Num
	}

	chartData := [][]interface{}{}
	for k, v := range statData {
		chartData = append(chartData, []interface{}{k, v["__SUM"]})
	}
	sort.Slice(chartData, func(i, j int) bool {
		return chartData[i][1].(int) > chartData[j][1].(int)
	})
	pieData, _ := json.Marshal(chartData)

	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["data"] = statData
	info["statusTitle"] = status.Title.GetTitles()
	info["pieData"] = string(pieData)

	utils.View.Execute(w, r, "stat/channel.html", info)
}
