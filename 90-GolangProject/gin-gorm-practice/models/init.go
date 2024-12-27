package models

import (
	"GolangProject/gin-gorm-practice/define"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

var RDB = InitRdb()

func Init() *gorm.DB {
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

	return db
}

func InitRdb() *redis.Client {
	opt, _ := redis.ParseURL(define.RedisConn)
	return redis.NewClient(opt)
	/*return redis.NewClient(&redis.Options{
		Addr:     "ip:port",
		Password: "", // no password set
		DB:       0,  // use default db
	})*/
}
