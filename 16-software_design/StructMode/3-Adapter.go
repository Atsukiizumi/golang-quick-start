package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

// 业务类，依赖V5接口
type IPhone struct {
	v V5
}

func NewPhone(v V5) *IPhone {
	return &IPhone{v}
}

func (p *IPhone) Charge() {
	fmt.Println("手机进行充电...")
	p.v.Use5V()
}

// 被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

/*
Target（目标抽象类）：目标抽象类定义客户所需接口，可以是一个抽象类或接口，也可以是具体类。
Adapter（适配器类）：适配器可以调用另一个接口，作为一个转换器，对Adaptee和Target进行适配，适配器类是适配器模式的核心，在对象适配器中，它通过继承Target并关联一个Adaptee对象使二者产生联系。
Adaptee（适配者类）：适配者即被适配的角色，它定义了一个已经存在的接口，这个接口需要适配，适配者类一般是一个具体类，包含了客户希望使用的业务方法，在某些情况下可能没有适配者类的源代码。
根据对象适配器模式结构图，在对象适配器中，客户端需要调用request()方法，而适配者类Adaptee没有该方法，但是它所提供的specificRequest()方法却是客户端所需要的。
为了使客户端能够使用适配者类，需要提供一个包装类Adapter，即适配器类。这个包装类包装了一个适配者的实例，从而将客户端与适配者衔接起来，在适配器的request()方法中调用适配者的specificRequest()方法。
因为适配器类与适配者类是关联关系（也可称之为委派关系），所以这种适配器模式称为对象适配器模式。
*/
func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
