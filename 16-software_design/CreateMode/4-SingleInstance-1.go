package main

import "fmt"

type singleton struct {
}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

/*
Singleton（单例）：在单例类的内部实现只生成一个实例，同时它提供一个静态的getInstance()工厂方法，让客户可以访问它的唯一实例；
为了防止在外部对其实例化，将其构造函数设计为私有；
在单例类内部定义了一个Singleton类型的静态对象，作为外部共享的唯一实例。

此例子是单例模式中的一种，属于“饿汉式”。
含义是，在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存。
饿汉式的好处是，不会出现线程并发创建，导致多个单例的出现，
但是缺点是如果这个单例对象在业务逻辑没有被使用，也会客观的创建一块内存对象。
*/
func main() {
	s := GetInstance()
	s.SomeThing()
}
