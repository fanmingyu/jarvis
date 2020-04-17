package stat

import(
	"net/http"
	"encoding/json"
	"sort"

	"smsgate/console/modules/status"
	"smsgate/server/storage"
	"smsgate/utils"
	"smsgate/console/modules/db"
)

//App 应用发送统计
func App(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	starttime, endtime := utils.GetUrlTime(vars, "")

	var data []struct{
		App    string
		Num    int
		Status string
	}
	db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("app, sum(num) as num, report_status as status").GroupBy("app, status").Find(&data)

	apps := make([]storage.SmsApp, 0)
	err := db.RMysql.Engine.Find(&apps)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	statData := make(map[string]map[string]int)
	for _, app := range apps {
		statData[app.Name] = make(map[string]int)
		statData[app.Name]["__SUM"] = 0
	}

	sum := 0
	for _, value := range data {
		title := status.Title.GetTitle(value.Status)
		statData[value.App][title] += value.Num
		statData[value.App]["__SUM"] += value.Num
		sum += value.Num
	}
	charData := [][]interface{}{}
	for k, v := range statData {
		charData = append(charData, []interface{}{k, v["__SUM"]})
	}
	sort.Slice(charData, func(i, j int) bool {
		return charData[i][1].(int) > charData[j][1].(int)
	})
	pieData, _ := json.Marshal(charData)

	info := make(map[string]interface{})

	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["data"] = statData
	info["statusTitle"] = status.Title.GetTitles()
	info["pieData"] = string(pieData)
	info["sum"] = sum

	utils.View.Execute(w, r, "stat/app.html", info)
}
