package models

import (
	"TinyURL/extend/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB
var err error

func Setup() {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBConf.User,
		conf.DBConf.Password,
		conf.DBConf.Host,
		conf.DBConf.Port,
		conf.DBConf.DBName)
	DB, err = gorm.Open(conf.DBConf.DBType, connectStr)
	if err != nil {
		fmt.Println("连接数据库失败--golang")
		time.Sleep(10 * time.Second)
		DB, err = gorm.Open(conf.DBConf.DBType, connectStr)
		if err != nil {
			panic(err.Error())
		}
	}
	if DB.Error != nil {
		fmt.Println("连接数据库失败--DB:", DB.Error)
	}

	DB.LogMode(conf.DBConf.Debug)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.AutoMigrate(&Link{}, &User{})
}
