package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn //连接信息

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线通知
func (this *User) Online() {
	//用户上线，将用户加入到OnlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上限消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线通知
func (this *User) Offline() {
	//用户下线，从OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this, "已下线")
}

// 给当前user对应的客户端发送信息
func (this *User) SendMessage(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户都有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "] " + user.Name + ": 在线...\t\n"
			this.SendMessage(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式: rename| 张三
		rename := strings.Split(msg, "|")[1]
		//判断name是否存在
		_, ok := this.server.OnlineMap[rename]
		if ok {
			this.SendMessage("当前用户名被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[rename] = this
			this.server.mapLock.Unlock()

			this.Name = rename
			this.SendMessage("您已经更新用户名：" + rename)
		}
	} else if len(msg) > 3 && msg[:3] == "pm|" {
		//消息格式：pm|张三|消息内容
		// 获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMessage("消息格式不正确，请使用\"pm|张三|你好啊\"格式。\n")
			return
		}

		//获取对方的用户对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMessage("当前用户不存在\n")
			return
		}

		remoteMsg := strings.Split(msg, "|")[2]
		if remoteMsg == "" {
			this.SendMessage("无消息内容，请重新发送。\n")
			return
		}
		remoteUser.SendMessage("【" + this.Name + "】对您说：" + remoteMsg + "\n")

	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channel的方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\r\n"))
	}
}
