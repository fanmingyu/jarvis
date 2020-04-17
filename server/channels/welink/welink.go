package welink

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"smsgate/utils"
)

//接口url
const WELINK_API_URL = "http://cf.51welink.com/submitdata/Service.asmx/g_Submit"

const WELINK_NAME = "xxx"
const WELINK_CREDIT_NAME = "xxxx"

const WELINK_PWD = "xxxx"

const WELINK_NOTICE_SPRDID = "xxx"
const WELINK_VERIFY_SPRDID = "xxxx"
const WELINK_INTERNATIANAL_SPRDID = "xxxx"
const WELINK_MARKETING_SPRDID = "xxxx"

const WELINK_SEND_SUCCESS = "WELINKSMS_SEND_SUCCESS"
const WELINK_SEND_FAILED = "WELINKSMS_SEND_FAILED"

var verifySPNumbers = []string{"106920478000", "106915898000", "106901408000"}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

type WelinkResponse struct {
	Status int    `xml:"State"`
	Msgid  string `xml:"MsgID"`
	Msg    string `xml:"MsgState"`
}

//请求微网接口
func RequestWelink(mobile utils.MobileString, content, sprdid, user string) (string, error) {
	params := url.Values{
		"sname":   {user},
		"spwd":    {WELINK_PWD},
		"scorpid": {""},
		"sprdid":  {sprdid},
		"sdst":    {string(mobile)},
		"smsg":    {content},
	}

	start := time.Now()
	resp, err := client.PostForm(WELINK_API_URL, params)
	cost := time.Now().Sub(start)
	if err != nil {
		log.Printf("welinkRequest failed. phone:%v, cost:%v, err:%v", mobile, cost, err)
		return "", err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var welinkResp WelinkResponse
	err = xml.Unmarshal(body, &welinkResp)
	if err != nil {
		log.Printf("welinkResponse unmarshal fail. phone:%v, err:%v", mobile, err)
		return "", err
	}
	if welinkResp.Status != 0 {
		log.Printf("welinkRequest failed. phone:%v, errCode:%v, message:%v", mobile, welinkResp.Status, welinkResp.Msg)
		err = errors.New("welink commit failed.")
		return "", err
	}

	log.Printf("welinkRequest Success. phone:%v, cost:%v, resp:%v", mobile, cost, welinkResp)
	return welinkResp.Msgid, nil
}

//IsVerify 根据spNumber来判断一条短信是否是验证码短信
func IsVerify(spNumber string) bool {
	for _, number := range verifySPNumbers {
		if spNumber == number {
			return true
		}
	}

	return false
}
