package weilai

import (
	"strings"

	"smsgate/server/channels/welink"
	"smsgate/utils"
)

func SendCompatible(mobile utils.MobileString, content string) (string, error) {
	var custCode string
	var pass string

	verify := strings.Contains(content, "验证码")
	if strings.HasPrefix(string(mobile), "00") {
		return welink.RequestWelink(mobile, content, welink.WELINK_INTERNATIANAL_SPRDID, welink.WELINK_NAME)
	}
	if verify {
		custCode = WEILAI_VERIFY_CUST_CODE
		pass = WEILAI_VERIFY_PASS
	} else {
		custCode = WEILAI_NOTICE_CUST_CODE
		pass = WEILAI_NOTICE_PASS
	}

	return RequestWeilai(mobile, content, custCode, pass)
}

func SendNotice(mobile utils.MobileString, content string) (string, error) {
	return RequestWeilai(mobile, content, WEILAI_NOTICE_CUST_CODE, WEILAI_NOTICE_PASS)
}

func SendVerify(mobile utils.MobileString, content string) (string, error) {
	return RequestWeilai(mobile, content, WEILAI_VERIFY_CUST_CODE, WEILAI_VERIFY_PASS)
}
