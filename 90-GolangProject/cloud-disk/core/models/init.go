package models

import (
	"GolangProject/cloud-disk/define"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = InitEngine()

//var RDB = InitRdb()

func InitEngine() *xorm.Engine {
	//操作数据库
	// dsn := define.SQLServerConn
	// sqlserver: sqlserver://username:password@ip:port?database=db_name
	// mysql: username:password@tcp(ip:port)/db_name?charset=utf8mb4&parseTime=True&loc=Local
	engine, err := xorm.NewEngine("mysql", define.MySQLConn)
	if err != nil {
		log.Println("Xorm New Engine init err:", err)
		return nil
	}

	return engine
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
