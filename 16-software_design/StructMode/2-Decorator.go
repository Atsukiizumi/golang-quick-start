package main

import "fmt"

// ------- 抽象层 -------
// 抽象的构建
type Phone interface {
	Show() //构件的功能
}

// 装饰器基础类（该类本应该为interface，但是Golang interface语法不可以有成员属性
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

// ------- 实现层 -------
// 具体的构件
type HuaWei struct{}

func (h *HuaWei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (x *XiaoMi) Show() {
	fmt.Println("秀出了XiaoMi手机")
}

// 具体的装饰类
type MoDecorator struct {
	Decorator //继承基础装饰器类
}

func (d *MoDecorator) Show() {
	d.phone.Show()       //调用被装饰构件的原方法
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) Show() {
	k.phone.Show()
	fmt.Println("手机壳的手机") //装饰额外的方法
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

// ------- 业务逻辑层 -------
func main() {
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show() //调用原构件方法

	fmt.Println("---------")
	//用贴膜装饰器装饰，得到新功能构件
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei) //通过HueWei ---> MoHuaWei
	moHuawei.Show()                   //调用装饰后新构件的方法

	fmt.Println("---------")
	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei) //通过HueWei ---> KeHuaWei
	keHuawei.Show()

	fmt.Println("---------")
	var keMoHuaWei Phone
	keMoHuaWei = NewMoDecorator(keHuawei) //通过KeHuaWei ---> KeMoHuaWei
	keMoHuaWei.Show()
}
