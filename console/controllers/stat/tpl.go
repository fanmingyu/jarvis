package stat

import(
	"net/http"
	"fmt"

	"smsgate/utils"
	"smsgate/console/modules/db"
)

//Tpl 模板发送统计
func Tpl(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	starttime, endtime := utils.GetUrlTime(vars, "")

	var data []struct{
		Tpl    int
		Num    int
		Sum    int
		AvgNum string
	}
	err := db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("tpl_id as tpl, count(*) as num, sum(num) as sum").GroupBy("tpl").Desc("sum").Find(&data)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	sum := 0
	for k, v := range data {
		data[k].AvgNum = fmt.Sprintf("%.2f", float64(v.Sum)/float64(v.Num))
		sum += v.Sum
	}

	info := make(map[string]interface{})

	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["data"] = data
	info["sum"] = sum

	utils.View.Execute(w, r, "stat/tpl.html", info)
}
