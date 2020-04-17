package status

import (
	"encoding/csv"
	"log"
	"os"
)

//ReportStatus 状态表的信息
type ReportStatus struct {
	Status   string
	Net      string
	Info     string
	Solution string
}

//Titles 回执状态的标题
type Titles map[string]string

//Status 缓存回执状态表
var Status = map[string]ReportStatus{}

//Title 一般统计的回执
var Title = Titles{
	"__FAIL":     "提交失败",
	"__COMMIT":   "提交成功",
	"DELIVRD":    "接收成功",
	"UNDELIVERD": "接收失败",
	"0":          "接收成功",
}

//TitleFull 全部回执的统计
var TitleFull = Titles{
	"":           "未提交",
	"__BlACK":    "黑名单",
	"__FAIL":     "提交失败",
	"__COMMIT":   "提交成功",
	"0":          "接收成功",
	"DELIVRD":    "接收成功",
	"UNDELIVERD": "接收失败",
}

//GetTitles 获取所有回执的标题
func (t Titles) GetTitles() map[string]string {
	titles := make(map[string]string)
	for k, v := range t {
		titles[v] = k
	}
	return titles
}

//GetTitle 获取某个回执对应的标题
func (t Titles) GetTitle(status string) string {
	if v, ok := t[status]; ok {
		return v
	}
	return t["UNDELIVERD"]
}

//Init 将回执状态表加载到缓存中
func Init() {
	file, err := os.Open("../static/status.csv")
	if err != nil {
		log.Printf("init report status failed. err:%v", err)
		return
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Printf("read file failed. err:%v", err)
		return
	}

	for _, record := range records {
		Status[record[0]] = ReportStatus {
			Status:   record[0],
			Net:      record[1],
			Info:     record[2],
			Solution: record[3],
		}
	}
}
