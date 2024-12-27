# Cloud Disk 
GOLANG 官网：https://golang.google.cn/

GOLANG Document： https://pkg.go.dev/

## 技术栈
- go-zero
- xorm   
- mysql
- redis
- etcd
- gorm

## go-zero

go-zero[官网](https://go-zero.dev/) [DOCS](https://go-zero.dev/docs/tasks)

1. 安装goctl
```shell
>  go install github.com/zeromicro/go-zero/tools/goctl@latest
// 验证
> goctl --version
```
2. 初始化api
```shell
> goctl api new core
```
3. 启动服务
```shell
> go run core.go -f etc/core-api.yaml 
```
4. 使用api文件生成代码
```shell
> goctl api go -api core.api -dir . -style go_zero
```
5. 


## xorm

1. 安装
```shell
> go get -u github.com/go-xorm/xorm
// mysql驱动
> go get github.com/go-sql-driver/mysql
```
2. 

