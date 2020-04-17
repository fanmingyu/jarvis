package sms

import (
	"encoding/json"
	"net/http"

	"smsgate/server/storage"
	"smsgate/server/workers"
	"smsgate/utils"
)

type SmsMobileContent struct {
	Mobile  string `json:"mobile"`
	Content string `json:"content"`
}

//Batch 兼容接口
func Batch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	appName := r.PostFormValue("app")
	if len(appName) < 1 {
		utils.ResponseJson(w, -1, "app is invaild.", nil)
		return
	}

	//TODO 临时变更，需要早日去除
	appName = "pmarketing"

	data := r.PostFormValue("data")
	if len(data) < 1 {
		utils.ResponseJson(w, -1, "data is invaild", nil)
		return
	}

	app := storage.CacheData.GetApp(appName)
	if app.Name == "" {
		utils.ResponseJson(w, -1, "app is not exist", nil)
		return
	}

	var records []SmsMobileContent
	json.Unmarshal([]byte(data), &records)
	if len(records) < 1 {
		utils.ResponseJson(w, -1, "no sms", nil)
		return
	}

	if len(records) > 100 {
		utils.ResponseJson(w, -1, "over limit 100", nil)
		return
	}

	prefix := app.Prefix
	for _, record := range records {
		sms := workers.SmsMessage{
			App:     appName,
			Mobile:  utils.MobileString(record.Mobile),
			Content: prefix + record.Content,
		}
		sms.Produce()
	}

	utils.ResponseJson(w, 0, "batch send success", nil)
}
