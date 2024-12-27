package main

import "fmt"

//------ 抽象层 ------

// 抽象的观察者
type Listener interface {
	OnTeacherComing() //观察者得知通知后要触发的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// ------- 实现层 ------
// 观察者学生
type Student1 struct {
	Badthing string
}

func (s *Student1) OnTeacherComing() {
	fmt.Println("一号同学 停止", s.Badthing)
}

func (s *Student1) DoBadthing() {
	fmt.Println("一号同学 正在", s.Badthing)
}

type Student2 struct {
	Badthing string
}

func (s *Student2) OnTeacherComing() {
	fmt.Println("二号同学 停止", s.Badthing)
}

func (s *Student2) DoBadthing() {
	fmt.Println("二号同学 正在", s.Badthing)
}

// 通知者班长
type ClassMonitor struct {
	listeners []Listener //需要通知的全部观察者集合
}

func (c *ClassMonitor) AddListener(listener Listener) {
	c.listeners = append(c.listeners, listener)
}

func (c *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range c.listeners {
		if l == listener {
			c.listeners = append(c.listeners[:index], c.listeners[index+1:]...)
			break
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, l := range c.listeners {
		//依次调用全部观察的具体动作
		l.OnTeacherComing()
	}
}

func main() {
	s1 := &Student1{
		Badthing: "阅读课外书籍",
	}
	s2 := &Student2{
		Badthing: "与女同学打闹",
	}

	cm := new(ClassMonitor)
	fmt.Println("上课了，老师还没来，同学们都在忙自己的事情...")
	s1.DoBadthing()
	s2.DoBadthing()

	cm.AddListener(s1)
	cm.AddListener(s2)

	fmt.Println("老师来了，班长通知到学生就位...")
	cm.Notify()
}
