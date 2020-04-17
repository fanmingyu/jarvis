package stat

import(
	"net/http"
	"sort"
	"encoding/json"

	"smsgate/utils"
	"smsgate/console/modules/db"
	"smsgate/server/channels"
)

var costRank = map[int]string {
	0 : "&lt; 10秒",
	1 : "&lt; 20秒",
	2 : "&lt; 30秒",
	3 : "&lt; 40秒",
	4 : "&lt; 50秒",
	5 : "&lt; 60秒",
	6 : "60秒及以上",
}


//Cost 接收时长统计
func Cost(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	info := make(map[string]interface{})

	starttime, endtime := utils.GetUrlTime(vars, "")

	var data []struct {
		Rank int
		Num  int
	}
	channel := vars.Get("channel")

	if channel == "" {
		db.RMysql.Engine.Table("sms_record").Where("report_status=? AND create_time BETWEEN ? AND ?", "DELIVRD", starttime, endtime).Select("if(report_time>create_time, ((report_time-create_time) div 10), 0) AS rank, sum(num) AS num").GroupBy("rank").Find(&data)
	} else {
		db.RMysql.Engine.Table("sms_record").Where("channel=? AND report_status=? AND create_time BETWEEN ? AND ?", channel, "DELIVRD", starttime, endtime).Select("if(report_time>create_time, ((report_time-create_time) div 10), 0) AS rank, sum(num) AS num").GroupBy("rank").Find(&data)
	}

	statData := make(map[int]map[string]int)
	for rank := range costRank {
		statData[rank] = map[string]int {
			"num" : 0,
		}
	}

	for _, value := range data {
		if _, ok := costRank[value.Rank]; !ok {
			value.Rank = 6
		}
		statData[value.Rank]["num"] += value.Num
	}

	sum := 0
	for i := 0; i < len(costRank); i++ {
		sum += statData[i]["num"]
		statData[i]["addUp"] = sum
	}

	charData := [][]interface{}{}
	for k, v := range statData {
		charData = append(charData, []interface{}{costRank[k], v["num"], k})
	}

	sort.Slice(charData, func(i, j int) bool{
		return charData[i][2].(int) < charData[j][2].(int)
	})
	pieData, _ := json.Marshal(charData)

	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["data"] = statData
	info["pieData"] = string(pieData)
	info["costRank"] = costRank
	info["sum"] = sum
	info["channels"] = channels.Channels
	info["channel"] = channel

	utils.View.Execute(w, r, "stat/cost.html", info)
}
