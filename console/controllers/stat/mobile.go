package stat

import(
	"net/http"
	"strconv"

	"smsgate/utils"
	"smsgate/console/modules/db"
)

//Mobile 号码发送统计
func Mobile(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	starttime, endtime := utils.GetUrlTime(vars, "")

	limit, _ := strconv.Atoi(vars.Get("limit"))
	if limit < 1 || limit > 1000 {
		limit = 50
	}

	data := []struct{
		Mobile string
		Total int
		Num int
	}{}
	db.RMysql.Engine.Table("sms_record").Where("create_time BETWEEN ? AND ?", starttime, endtime).Select("mobile, count(*) as total, sum(num) as num").GroupBy("mobile").Desc("num").Limit(limit).Find(&data)

	info := map[string]interface{}{}
	info["start"] = vars.Get("start")
	info["end"] = vars.Get("end")
	info["limit"] = limit
	info["data"] = data

	utils.View.Execute(w, r, "stat/mobile.html", info)
}
