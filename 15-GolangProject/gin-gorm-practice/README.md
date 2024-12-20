## Gin-Gorm GOLANG 练手框架

GOLANG 官网：https://golang.google.cn/

GOLANG Document： https://pkg.go.dev/

GORM 官网：https://gorm.io/
```shell
go get -u gorm.io/gorm
#MySQL数据库驱动
go get -u gorm.io/driver/mysql
#SQL Server数据库驱动
go get -u gorm.io/driver/sqlserver
```
Gorm数据库连接
```go
//MySQL模板
func Init() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        log.Println("gorm init err:", err)
	}
	return db
}

//SQL Server模板
func Init() *gorm.DB {
    //操作数据库
    dsn := "sqlserver://user:password@localhost:port?database=dbname"
    db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Println("gorm init err:", err)
    }

    return db
}
```


Gin 官网：https://gin-gonic.com/zh-cn/
```shell
go get -u github.com/gin-gonic/gin
```

## 整合 Swagger

项目地址：https://github.com/swaggo/gin-swagger
接口访问：http://localhost:8080/swagger/index.html

安装：

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

如何使用：

1. 使用 gin-swagger 规则为 api 和 main 函数添加注释，如下所示：

```go
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
```

2. 使用 `swag init` 命令生成文档，生成的文档将存储在 `docs/` 处。


## 依赖：
    - jwt
    - email
    - redis
    - google/uuids

## 安装 jwt

项目地址：https://github.com/golang-jwt/jwt

```shell
go get -u github.com/golang-jwt/jwt/v5
```

## 安装 email
```shell
go get -u github.com/jordan-wright/email
```

## 安装 redis
项目地址：https://github.com/redis/go-redis

```shell
go get -u github.com/redis/go-redis/v9
```

```go
var ctx  = context.Background()

```

## 安装 Google-uuid
项目地址：https://github.com/google/uuid

```shell
go get -u github.com/google/uuid
```

## TODO List:

1. 更新框架
   - [ ] Swagger
         项目地址：https://github.com/go-swagger/go-swagger
   - [ ] Email -> 使用原生mail库
   - [ ]
2. 新增框架
   - [ ] Go-Zero
   项目地址：https://github.com/tal-tech/go-zero
   - [ ] 日志
   - [ ] NSQ 消息队列 
   项目地址:https://github.com/nsqio/nsq
   - [ ] TiDB?/ETCD?
   - [ ] gRPC?
   - [ ] ~~Vue3.5(咕咕咕)~~
