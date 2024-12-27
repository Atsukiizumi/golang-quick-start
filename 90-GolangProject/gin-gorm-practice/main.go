package main

import (
	"GolangProject/gin-gorm-practice/router"
)

func main() {
	r := router.Router()
	r.Run() // 监听并在0.0.0.0:8080上启动服务
}
