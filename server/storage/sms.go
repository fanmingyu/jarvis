package storage

import (
	"time"

	"smsgate/utils"
)

const (
	SMS_STATUS_SAVED        = 1
	SMS_STATUS_SEND_FAILED  = 2
	SMS_STATUS_SEND_SUCCESS = 3
)

type SmsStorage interface {
	InsertSmsRecord(app, mobile, content, channel, reportStatus string, tplId int) (int64, error)
	UpdateSmsStatus(ids []int64, status int, reportStatus string, msgId string) error
}

type SmsRecord struct {
	Id             int64 `xorm:"notnull pk autoincr"`
	App            string
	Mobile         string
	Content        string
	Channel        string
	Msgid          string
	CreateTime     int `xorm:"created"`
	RequestTime    int
	RequestStatus  int
	ReportTime     int
	ReportRecvTime int
	ReportStatus   string
	Num            int
	TplId          int
}

type SmsMysql struct{}

func (m *SmsMysql) InsertSmsRecord(app, mobile, content, channel, reportStatus string, tplId int) (id int64, err error) {
	sms := SmsRecord{
		App:           app,
		Content:       content,
		Channel:       channel,
		RequestStatus: SMS_STATUS_SAVED,
		Num:           utils.SmsCounter(content),
		Mobile:        mobile,
		ReportStatus:  reportStatus,
		TplId:         tplId,
	}
	_, err = Mysql.Engine.Insert(&sms)

	return sms.Id, err
}

func (m *SmsMysql) UpdateSmsStatus(ids []int64, status int, reportStatus string, msgId string) error {
	data := SmsRecord{
		RequestStatus: status,
		Msgid:         msgId,
		RequestTime:   int(time.Now().Unix()),
		ReportStatus:  reportStatus,
	}
	_, err := Mysql.Engine.In("id", ids).Update(&data)
	return err
}
