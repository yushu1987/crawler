package mysql

import (
	"crawler/lib"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db     *gorm.DB
	dbname = "database"
)

func InitMysqlConn() {
	var err error

	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = lib.Config.GetString(dbname + ".addr")
	config.User = lib.Config.GetString(dbname + ".user")
	config.Passwd = lib.Config.GetString(dbname + ".pass")
	config.DBName = lib.Config.GetString(dbname + ".dbname")

	DSN := config.FormatDSN()
	db, err = gorm.Open("mysql", DSN)
	if err != nil {
		fmt.Printf("mysql connect error:%s", err.Error())
		panic("failed to connect database")
	}

	// 初始化连接池
	db.DB().SetMaxIdleConns(lib.Config.GetInt(dbname + ".idle_conns"))
	db.DB().SetMaxOpenConns(lib.Config.GetInt(dbname + ".max_conns"))
	idleTimeout := lib.Config.GetInt(dbname + ".idle_timeout")
	if idleTimeout > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(idleTimeout) * time.Second)
	}

	db.LogMode(false)
}

func Close() {
	db.Close()
}
