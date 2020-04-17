package report

import (
	"time"
	"log"

	"smsgate/report/modules/db"
	"smsgate/server/storage"
)

type Report struct {
	Msgid  string
	Mobile string
	Status string
	Time   int
}

func (r *Report) UpdateSmsRecord() error {
	data := storage.SmsRecord{
		ReportTime:     r.Time,
		ReportRecvTime: int(time.Now().Unix()),
		ReportStatus:   r.Status,
	}

	_, err := db.Mysql.Engine.Where("mobile=? AND msgid=?", r.Mobile, r.Msgid).Update(&data)
	if err != nil {
		log.Printf("sms report. update sms record failed. err:%v", err)
	}
	return err
}
