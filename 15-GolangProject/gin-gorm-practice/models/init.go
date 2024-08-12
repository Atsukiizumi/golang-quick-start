package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	//操作数据库
	dsn := "sqlserver://sa:????@localhost:1433?database=gin_gorm_oj"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm init err:", err)
	}

	return db
}
