package workers

import (
	"testing"

	"github.com/go-sql-driver/mysql"

	"smsgate/server/storage"
)

func TestProcessMessage(t *testing.T) {
	config := mysql.Config{
		Addr:   "10.20.69.57:3306",
		DBName: "sms",
		User:   "sms",
		Passwd: "sms",
	}
	storage.MysqlInit(&config)
	storage.CacheInit()

	sms := SmsMessage{
		App:     "test",
		Mobile:  "15927046337",
		Content: "test content",
	}
	sms.Produce()

	var smsData storage.SmsRecord
	_, err := storage.Mysql.Engine.Id(sms.Id).Get(&smsData)
	if err != nil {
		t.Errorf("find sms failed. err:%v", err)
	}

	if smsData.App != "test" || smsData.Mobile != "15927046338" {
		t.Errorf("test sms produce failed. smsData:%v", smsData)
	}

	sms.Consume()
	var smsRecord storage.SmsRecord
	_, err = storage.Mysql.Engine.Id(sms.Id).Get(&smsRecord)
	if err != nil {
		t.Errorf("find sms failed. err:%v", err)
	}

	if smsRecord.RequestStatus != storage.SMS_STATUS_SEND_SUCCESS {
		t.Errorf("test sms consume failed. smsRecord:%v", smsRecord)
	}

}
