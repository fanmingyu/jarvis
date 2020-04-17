package storage

import (
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestCache(t *testing.T) {
	config := mysql.Config{
		User:   "sms",
		Passwd: "sms",
		Addr:   "10.20.69.57:3306",
		DBName: "sms",
	}
	MysqlInit(&config)

	app := SmsApp{
		Name:    "test",
		Secret:  "test",
		Prefix:  "【xx签名】",
		Channel: "test",
		Worker:  "test",
	}

	Mysql.Engine.Insert(&app)

	tpl := SmsTpl{
		Name:    "normal",
		Content: "{$var1}",
	}
	Mysql.Engine.Insert(&tpl)

	blacklist := SmsBlacklist{
		Mobile: "15900000000",
	}
	Mysql.Engine.Insert(&blacklist)

	CacheInit()
	smsTpl := CacheData.GetTpl("normal")
	if smsTpl.Content != "{$var1}" {
		t.Errorf("test get tpl failed.")
	}

	smsApp := CacheData.GetApp("test")
	if smsApp.Secret != "test" || smsApp.Prefix != "【xx签名】" {
		t.Errorf("test get app failed.")
	}

	smsBlacklist := CacheData.GetBlacklist("15900000000")
	if smsBlacklist != 1 {
		t.Errorf("test get blacklist failed.")
	}
}
