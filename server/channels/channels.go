package channels

import (
	"smsgate/server/channels/aliyun"
	"smsgate/server/channels/sendcloud"
	"smsgate/utils"
)

type ChannelHandle struct {
	Name     string
	SendFunc func(mobile utils.MobileString, content string, out_id string) (msgId string, err error)
}

var Channels = map[string]ChannelHandle{
	// "weilai": ChannelHandle{
	// 	Name:     "未来无线兼容通道",
	// 	SendFunc: weilai.SendCompatible,
	// },
	// "weilaiNotice": ChannelHandle{
	// 	Name:     "未来无线通知",
	// 	SendFunc: weilai.SendNotice,
	// },
	// "weilaiVerify": ChannelHandle{
	// 	Name:     "未来无线验证码",
	// 	SendFunc: weilai.SendVerify,
	// },
	// "welink": ChannelHandle{
	// 	Name:     "微网兼容通道",
	// 	SendFunc: welink.SendCompatible,
	// },
	// "welinkNotice": ChannelHandle{
	// 	Name:     "微网通知",
	// 	SendFunc: welink.SendNotice,
	// },
	// "welinkVerify": ChannelHandle{
	// 	Name:     "微网验证码",
	// 	SendFunc: welink.SendVerify,
	// },
	// "welinkMarketing": ChannelHandle{
	// 	Name:     "微网营销",
	// 	SendFunc: welink.SendMarketing,
	// },
	// "welinkInternatel": ChannelHandle{
	// 	Name:     "微网国际",
	// 	SendFunc: welink.SendInternatel,
	// },
	// "test": ChannelHandle{
	// 	Name:     "测试通道",
	// 	SendFunc: test.Send,
	// },
	// "welinkCredit": ChannelHandle{
	// 	Name:     "微网通道",
	// 	SendFunc: welink.SendCredit,
	// },
	"aliyunSms": ChannelHandle{
		Name:     "阿里云短信通道",
		SendFunc: aliyun.Send,
	},
	"sendCloud": ChannelHandle{
		Name:     "SendCloud短信通道",
		SendFunc: sendcloud.Send,
	},
}
