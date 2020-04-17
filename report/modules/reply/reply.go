package reply

import (
	"smsgate/report/modules/db"
)

type SmsReply struct {
	Id         int `xorm:"notnull pk autoincr"`
	Msgid      string
	SpCode     string
	Mobile     string
	Content    string
	RecvTime   int
	CreateTime int `xorm:"created"`
	Account    string
}

func (r *SmsReply) Insert() error {
	_, err := db.Mysql.Engine.Insert(r)

	return err
}
