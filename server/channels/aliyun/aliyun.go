package aliyun

import (
	"fmt"
	"smsgate/utils"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func Send(mobile utils.MobileString, content string, out_id string) (string, error) {
	splitContent := strings.Split(content, "||")
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAIHYWi35UprXb8", "Or6KCTKeuX7sXLJ6vYB2afGzKJiVfa")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = string(mobile)
	request.SignName = splitContent[0]
	// request.TemplateCode = "SMS_172208844"
	request.TemplateCode = out_id
	request.TemplateParam = splitContent[1]

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return "", nil
}

// func SendNotice(mobile utils.MobileString, content string, out_id string) (string, error) {
// 	splitContent := strings.Split(content, "||")
// 	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAIHYWi35UprXb8", "Or6KCTKeuX7sXLJ6vYB2afGzKJiVfa")

// 	request := dysmsapi.CreateSendSmsRequest()
// 	request.Scheme = "https"

// 	request.PhoneNumbers = string(mobile)
// 	request.SignName = splitContent[0]
// 	// request.TemplateCode = "SMS_172208865"
// 	request.TemplateCode = out_id
// 	request.TemplateParam = splitContent[1]

// 	response, err := client.SendSms(request)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	fmt.Printf("response is %#v\n", response)
// 	return "", nil
// }
