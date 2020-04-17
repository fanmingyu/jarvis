package storage

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type MysqlType struct {
	Config *mysql.Config
	Engine *xorm.Engine
}

var Mysql *MysqlType

func MysqlInit(config *mysql.Config) {
	Mysql = &MysqlType{}

	Mysql.Init(config)
}

//初始化mysql
func (m *MysqlType) Init(config *mysql.Config) {
	m.Config = config
	m.Config.Timeout = 2 * time.Second
	m.Config.AllowNativePasswords = true
	m.Config.Net = "tcp"

	url := m.Config.FormatDSN()
	var err error
	m.Engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		panic(err)
	}

	err = m.Engine.Ping()
	if err != nil {
		panic(err)
	}
}
