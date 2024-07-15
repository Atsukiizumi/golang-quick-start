package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	// 用户连接基本信息
	Ip   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User //存储在线用户
	mapLock   sync.RWMutex     //锁

	//消息广播的channel
	Message chan string
}

// 监听Message广播消息channel的

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 监听message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发给全部在线用户
		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			user.C <- msg
		}

		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "{" + user.Addr + "}" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//当前链接业务
	fmt.Println("连接建立成功")

	// 用户上线，将用户加入到OnlineMap
	user := NewUser(conn)

	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//广播当前用户上限消息
	this.BroadCast(user, "已上线")

	//当前handler阻塞
	select {}
}

// 启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//close listen socket
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("listener close err:", err)
			return
		}
	}(listener)

	//启动监听msg
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}

}
