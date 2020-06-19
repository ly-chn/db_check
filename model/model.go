package model

import (
	"db_check/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var gormDb *gorm.DB

func GetDb() (db *gorm.DB) {
	return gormDb
}

func InitDb() {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname)
	// 打开连接
	gormDb, err = gorm.Open("mysql", url)
	if err != nil {
		log.Fatal("MySQL 数据库无法连接", err)
	}

	// 防止表名加复数
	gormDb.SingularTable(true)
	// 最大空闲
	gormDb.DB().SetMaxIdleConns(10)
	// 最大开启
	gormDb.DB().SetMaxOpenConns(100)
}
