package models

import (
	"GolangProject/iot-platform/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = InitDB()

//var RDB = InitRdb()

func InitDB() *gorm.DB {
	//操作数据库
	// dsn := define.SQLServerConn
	// sqlserver: sqlserver://username:password@ip:port?database=db_name
	// mysql: username:password@tcp(ip:port)/db_name?charset=utf8mb4&parseTime=True&loc=Local
	dsn := define.MySQLConn
	//db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm init err:", err)
	}
	//数据库迁移
	err = db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})
	if err != nil {
		log.Println("db.AutoMigrate err:", err)
	}

	return db
}

//func InitRdb() *redis.Client {
//	opt, _ := redis.ParseURL(define.RedisConn)
//	return redis.NewClient(opt)
//	/*return redis.NewClient(&redis.Options{
//		Addr:     "ip:port",
//		Password: "", // no password set
//		DB:       0,  // use default db
//	})*/
//}
