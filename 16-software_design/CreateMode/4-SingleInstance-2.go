package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var once sync.Once

// 标记
var initialized uint32

// 定义锁
var lock sync.Mutex

type singleton2 struct{}

var instance2 *singleton2

func GetInstance2() *singleton2 {
	//如果标记为被设置，直接返回，不加锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instance2
	}

	//为了线程安全，增加了互斥
	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance2 = new(singleton2)
		//设置标记位
		atomic.StoreUint32(&initialized, 1)
	}
	return instance2
}

func GetInstance2_2() *singleton2 {
	once.Do(func() {
		instance2 = new(singleton2)
	})
	return instance2
}

func (s *singleton2) SomeThing2() {
	fmt.Println("安全线程调用")
}

/*
如果多个线程或者协程同时首次调用GetInstance()方法有概率导致多个实例被创建，则违背了单例的设计初衷。
那么在上面的基础上进行修改，可以利用Sync.Mutex进行加锁，保证线程安全。
*/
func main() {
	s := GetInstance2()
	s.SomeThing2()

	s2 := GetInstance2_2()
	s2.SomeThing2()
}
