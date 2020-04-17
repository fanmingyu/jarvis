package stat

import(
	"net/http"
	"sort"
	"encoding/json"

	"smsgate/utils"
	"smsgate/console/modules/db"
)

//NumberRank 发送条数分级
var NumberRank = map[int]string {
	1 : "1条",
	2 : "2条",
	3 : "3条",
	4 : "3条以上",
}

//Number 发送条数统计
func Number(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	starttime, endtime := utils.GetUrlTime(vars, "")

	var data []struct {
		Title int
		Num   int
		Sum   int
	}
	db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("if(num<4, num, 4) as title, sum(num) as num").GroupBy("title").Find(&data)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Title < data[j].Title
	})

	sum := 0
	for index, value := range data {
		sum += value.Num
		data[index].Sum = sum
	}

	charData := [][]interface{}{}
	for _, v := range data {
		charData = append(charData, []interface{}{NumberRank[v.Title], v.Num})
	}
	sort.Slice(charData, func(i, j int) bool {
		return charData[i][1].(int) > charData[j][1].(int)
	})
	pieData, _ := json.Marshal(charData)

	info := make(map[string]interface{})

	info["data"] = data
	info["numberRank"] = NumberRank
	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["sum"] = sum
	info["pieData"] = string(pieData)

	utils.View.Execute(w, r, "stat/number.html", info)
}
