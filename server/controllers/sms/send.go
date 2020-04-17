package sms

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"smsgate/server/storage"
	"smsgate/server/workers"
	"smsgate/utils"
)

func Send(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	appName := r.PostFormValue("app")
	if len(appName) < 1 {
		utils.ResponseJson(w, -1, "app is invalid.", nil)
		return
	}
	app := storage.CacheData.GetApp(appName)
	if app.Name == "" {
		utils.ResponseJson(w, -1, "app is not exist.", nil)
		return
	}

	tplName := r.PostFormValue("tpl")
	if len(tplName) < 1 {
		utils.ResponseJson(w, -1, "tpl is invalid.", nil)
		return
	}
	tpl := storage.CacheData.GetTpl(tplName)
	if tpl.Name == "" {
		utils.ResponseJson(w, -1, "tpl is not exist.", nil)
		return
	}

	mobile := r.PostFormValue("mobile")
	if len(mobile) < 10 {
		utils.ResponseJson(w, -1, "mobile is invalid.", nil)
		return
	}

	mobiles := strings.Split(mobile, ",")
	if len(mobiles) > 100 {
		utils.ResponseJson(w, -1, "over limit 100.", nil)
		return
	}

	varStr := r.PostFormValue("vars")
	content, err := getContent(tpl.Content, varStr, app.Prefix)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	//	err = utils.VerifySignature(r.PostForm, app.Secret, "sign")
	//	if err != nil {
	//		utils.MonitorAdd("SIGN_NOT_MATCH", 1)
	//		//utils.ResponseJson(w, -1, err.Error(), nil)
	//		//return
	//	}

	sms := workers.SmsMessage{
		App:      appName,
		Mobile:   utils.MobileString(mobile),
		Content:  content,
		TplId:    tpl.Id,
		TplOutId: tpl.OutId,
	}

	sms.Produce()
	utils.ResponseJson(w, 0, "send sms success", nil)
}

//根据模板和参数获取正确的短信文本
func getContent(tplContent, varStr, prefix string) (string, error) {
	var err error

	varCount := strings.Count(tplContent, "$var")
	var vars []string
	json.Unmarshal([]byte(varStr), &vars)

	if varCount != len(vars) {
		err = errors.New("get content failed. the vars doesn't match tpl")
		return "", err
	}

	r := make([]string, 0)
	for i := 0; i < varCount; i++ {
		r = append(r, "{$var"+strconv.Itoa(i+1)+"}", vars[i])
	}
	replace := strings.NewReplacer(r...)
	content := replace.Replace(tplContent)

	return prefix + "||" + content, err
}
