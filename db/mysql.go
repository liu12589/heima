package db

import (
	"AIAssistServer/constants"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", constants.UserName, constants.Password, constants.Host, constants.Port, constants.Dbname, constants.Timeout)
	ms, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Database connection failure, error= :%s", err.Error())
	}
	db = ms
	//fmt.Println("数据库连接成功")
}

func GetDb() *gorm.DB {
	return db
}
