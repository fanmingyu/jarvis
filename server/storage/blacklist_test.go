package storage

import (
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestBlacklist(t *testing.T) {
	config := mysql.Config{
		User:   "sms",
		Passwd: "sms",
		Addr:   "10.20.69.57:3306",
		DBName: "sms",
	}
	MysqlInit(&config)

	AddBlacklist("1257924688")
	CacheInit()

	if !InBlacklist("1257924688") {
		t.Errorf("test add blacklist failed")
	}
	RemoveBlacklist("1257924688")
	CacheInit()

	if InBlacklist("1257924688") {
		t.Errorf("test remove blacklist failed")
	}
}
