package weilai

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"errors"

	"smsgate/utils"
)

//未来无线接口参数
const WEILAI_API_URL = "http://123.58.255.70:8860/sendSms"

const WEILAI_VERIFY_CUST_CODE = "500072"

const WEILAI_VERIFY_PASS = "xxxx"

const WEILAI_NOTICE_CUST_CODE = "xxx"

const WEILAI_NOTICE_PASS = "xxxx"

const WEILAI_SEND_FAILED = "WEILAISMS_SEND_FAILED"

const WEILAI_SEND_SUCCESS = "WEILAISMS_SEND_SUCCESS"

var client = &http.Client{
	Timeout: 5 * time.Second,
}

//未来无线请求包
type WeilaiRequest struct {
	CustCode string `json:"cust_code"`
	SpCode   string `json:"sp_code"`
	Content  string `json:"content"`
	Mobiles  string `json:"destMobiles"`
	Sign     string `json:"sign"`
	Uid      string `json:"uid"`
}

//未来无线返回包
type WeilaiResponse struct {
	RespCode       string                   `json:"respCode"`
	TotalChargeNum int                      `json:"totalChargeNum"`
	RespMsg        string                   `json:"respMsg"`
	Result         []map[string]interface{} `json:"result"`
}


//调用weilai第三方接口
func RequestWeilai(mobile utils.MobileString, content, custCode, pass string) (string, error) {
	var weilaiResp WeilaiResponse
	req := &WeilaiRequest{
		CustCode: custCode,
		SpCode:   "",
		Content:  content,
		Mobiles:  string(mobile),
		Uid:      custCode,
	}

	//生成签名
	h := md5.New()
	h.Write([]byte(content))
	h.Write([]byte(pass))
	req.Sign = hex.EncodeToString(h.Sum(nil))

	reqJson, _ := json.Marshal(req)

	//发送请求
	start := time.Now()
	resp, err := client.Post(WEILAI_API_URL, "application/json;charset=utf-8", bytes.NewBuffer([]byte(string(reqJson))))
	cost := time.Now().Sub(start)

	if err != nil {
		log.Printf("Weilai request failed. phone:%v, cost:%v, err:%v", mobile, cost, err)
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &weilaiResp); err != nil {
		log.Printf("Weilai request failed. unmarshal response body fail. phone:%v", mobile)
		return "", err
	}

	if weilaiResp.RespCode != "0" {
		err = errors.New(weilaiResp.RespMsg)
		log.Printf("Weilai request failed. phone:%v, err:%v", mobile, err)
		return "", err
	}

	result := weilaiResp.Result[0]

	log.Printf("Weilai request success. phone:%v, cost:%v, resp:%v", mobile, cost, result)
	return result["msgid"].(string), nil
}

//IsVerify 根据uid来判断一条短信是否是验证码短信
func IsVerify(uid string) bool {
	if uid == WEILAI_VERIFY_CUST_CODE {
		return true
	}

	return false
}
