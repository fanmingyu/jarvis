package sendcloud

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"smsgate/utils"
	"strings"
	"time"
)

func Send(mobile utils.MobileString, content string, out_id string) (string, error) {
	var r http.Request

	splitContent := strings.Split(content, "||")
	smsUser := "ifaxin_126445_ShXZWvyf"
	smsKey := "LFwLnTDB4MliGLDRDbUxGU6RfpmeJkkp"
	templateId := out_id
	phone := string(mobile)
	vars := splitContent[1]
	// md5Ctx := md5.New()
	signature := md5.Sum([]byte(smsKey + "&phone=" + phone + "&smsUser=" + smsUser + "&templateId=" + templateId + "&vars=" + vars + "&" + smsKey))
	r.ParseForm()
	r.Form.Add("smsUser", smsUser)
	r.Form.Add("templateId", templateId)
	r.Form.Add("phone", phone)
	r.Form.Add("vars", vars)
	r.Form.Add("signature", fmt.Sprintf("%x", signature))
	bodystr := strings.TrimSpace(r.Form.Encode())
	post("http://www.sendcloud.net/smsapi/send", bodystr)
	return "", nil
}

//发送POST请求
//url:请求地址，bodyStr:POST请求提交的数据
//content:请求放返回的内容
func post(url string, bodyStr string) (content string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyStr)))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	fmt.Println(content)
	return content
}
