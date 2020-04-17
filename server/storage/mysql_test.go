package storage

import (
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T) {
	config := mysql.Config{
		User:   "sms",
		Passwd: "sms",
		Addr:   "10.20.69.57:3306",
		DBName: "sms",
	}
	MysqlInit(&config)
	mysql := &SmsMysql{}

	sms := SmsRecord{
		App:           "pcn",
		Mobile:        "159999999",
		Content:       "sdfsdfsadf",
		Channel:       "weilai",
		RequestStatus: 1,
		CreateTime:    int(time.Now().Unix()),
	}
	err := mysql.InsertSmsRecord(&sms)
	if err != nil {
		t.Errorf("insert fail. err:%v", err)
	}
	t.Logf("insert success. sms id is:%v", sms.Id)

	err = mysql.UpdateSmsStatus(sms.Id, SMS_STATUS_SEND_SUCCESS, "123")
	if err != nil {
		t.Errorf("update fail. err:%v", err)
	}

	sms.RequestStatus = SMS_STATUS_SEND_SUCCESS
	has, err := Mysql.Engine.Get(&sms)
	if err != nil {
		t.Errorf("get failed. err:%v", err)
	}

	if !has {
		t.Errorf("test update fail.")
	}
}
