package test

import(
	"errors"
	"strings"
	"strconv"
	"time"

	"smsgate/utils"
)

func Send(mobile utils.MobileString, content string) (string, error) {
	if strings.Contains(content, "fail") {
		err := errors.New("test wrong send")
		return "", err
	}

	msgId := strconv.FormatInt(time.Now().UnixNano(), 10)
	return msgId, nil
}
