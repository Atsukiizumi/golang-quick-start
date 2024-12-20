package oa_api

import "github.com/gin-gonic/gin"

func main() {
	//r := router
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":12221") // 监听并在 0.0.0.0:8080 上启动服务
}
