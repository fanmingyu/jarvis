package send

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"smsgate/console/modules/node"
	"smsgate/utils"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

//Send 发送短信
func Send(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mobile := strings.TrimSpace(r.PostFormValue("mobile"))
	if mobile == "" {
		utils.ResponseJson(w, -1, "mobile is invaild", nil)
		return
	}

	app := r.PostFormValue("app")
	if app == "" {
		utils.ResponseJson(w, -1, "app is invaild", nil)
		return
	}

	var content []string
	content = append(content, r.PostFormValue("content"))
	vars, _ := json.Marshal(content)

	params := url.Values{
		"app":    {app},
		"mobile": {mobile},
		"tpl":    {"test"},
		"vars":   {string(vars)},
	}

	node := node.Registry.GetNodes()[0]
	smsURL := "http://" + node.GetUrl() + "/sms/send"
	resp, err := client.PostForm(smsURL, params)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result utils.Response
	err = json.Unmarshal(body, &result)

	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseJson(w, result.Code, result.Message, result.Data)
}
