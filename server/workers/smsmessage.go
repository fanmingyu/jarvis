package workers

import (
	"log"
	"strings"

	"smsgate/server/channels"
	"smsgate/server/storage"
	"smsgate/utils"
)

// SmsMessage 定义短信结构体
type SmsMessage struct {
	Id      []int64
	App     string
	Mobile  utils.MobileString
	Content string
	TplId   int
	//第三方短信平台中的模板ID
	TplOutId string
}

// 生产一条短信
func (message *SmsMessage) Produce() {
	app := storage.CacheData.GetApp(message.App)
	worker := app.Worker

	white, black := storage.DivideBlacklist(string(message.Mobile))

	for _, mobile := range black {
		_, err := Workers[worker].Storage.InsertSmsRecord(message.App, mobile, message.Content, app.Channel, "__BLACK", message.TplId)
		if err != nil {
			log.Printf("insert db failed. phone:%s, err:%s", message.Mobile, err)
			continue
		}
	}

	for _, mobile := range white {
		id, err := Workers[worker].Storage.InsertSmsRecord(message.App, mobile, message.Content, app.Channel, "", message.TplId)
		if err != nil {
			log.Printf("insert db failed. phone:%s, err:%s", message.Mobile, err)
			continue
		}
		message.Id = append(message.Id, id)
	}

	if len(white) < 1 {
		return
	}

	message.Mobile = utils.MobileString(strings.Join(white, ","))
	//入队列
	Workers[worker].Queue.Enqueue(*message)
	log.Printf("produce sms success. phone:%s, app:%s", message.Mobile, message.App)
}

//处理一条短信
func (message *SmsMessage) Consume() {
	app := storage.CacheData.GetApp(message.App)
	mobiles := strings.Split(string(message.Mobile), ",")

	chargeNumber := utils.SmsCounter(message.Content) * len(mobiles)
	monitorNodeStr := strings.ToUpper(app.Channel) + "_SMS_SEND_"

	chanHandle, ok := channels.Channels[app.Channel]
	if !ok {
		log.Printf("send sms failed. channel is not exit. phone:%s, chan:%s", message.Mobile, app.Channel)
		return
	}

	worker := Workers[app.Worker]
	msgId, err := chanHandle.SendFunc(message.Mobile, message.Content, message.TplOutId)
	if err != nil {
		utils.MonitorAdd(monitorNodeStr+"FAILED", chargeNumber)
		log.Printf("sms send failed. mobile:%s, content:%v, err:%v", message.Mobile, message.Content, err)
		worker.Storage.UpdateSmsStatus(message.Id, storage.SMS_STATUS_SEND_FAILED, "__FAIL", "")
		return
	}

	utils.MonitorAdd(monitorNodeStr+"SUCCESS", chargeNumber)
	log.Printf("sms send success. mobile:%s, msgId: %s, content:%v", message.Mobile, msgId, message.Content)
	worker.Storage.UpdateSmsStatus(message.Id, storage.SMS_STATUS_SEND_SUCCESS, "__COMMIT", msgId)
}
