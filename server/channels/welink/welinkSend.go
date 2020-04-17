package welink

import(
	"strings"

	"smsgate/utils"
)

func SendCompatible(mobile utils.MobileString, content string) (string, error) {
	var sprdid string

	verify := strings.Contains(content, "验证码")
	if strings.HasPrefix(string(mobile), "00") {
		sprdid = WELINK_INTERNATIANAL_SPRDID
	} else if verify {
		sprdid = WELINK_VERIFY_SPRDID
	} else {
		sprdid = WELINK_NOTICE_SPRDID
	}

	return RequestWelink(mobile, content, sprdid, WELINK_NAME)
}

func SendNotice(mobile utils.MobileString, content string) (string, error) {
	return RequestWelink(mobile, content, WELINK_NOTICE_SPRDID, WELINK_NAME)
}

func SendVerify(mobile utils.MobileString, content string) (string, error) {
	return RequestWelink(mobile, content, WELINK_VERIFY_SPRDID, WELINK_NAME)
}

func SendMarketing(mobile utils.MobileString, content string) (string, error) {
	return RequestWelink(mobile, content, WELINK_MARKETING_SPRDID, WELINK_NAME)
}

func SendInternatel(mobile utils.MobileString, content string) (string, error) {
	return RequestWelink(mobile, content, WELINK_INTERNATIANAL_SPRDID, WELINK_NAME)
}

func SendCredit(mobile utils.MobileString, content string) (string, error) {
	return RequestWelink(mobile, content, WELINK_VERIFY_SPRDID, WELINK_CREDIT_NAME)
}
