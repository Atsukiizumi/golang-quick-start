package main

import "fmt"

// 抽象类，制作饮料，包裹一个模板的全部实现步骤
type Beverage interface {
	BoilWater() //煮开水
	Brew()      //冲泡
	PourInCup() //倒入杯中
	AddThings() //添加酌料

	WantAddThings() bool //是否加入酌料Hook
}

// 封装一套流程模板，让具体的制作流程继承且实现
type template struct {
	b Beverage
}

// 封装的固定模板
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	//子类可以重写该方法来决定是否执行下面动作
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

// 具体的模板子类 制作咖啡
type MakeCoffee struct {
	template
}

func (m MakeCoffee) BoilWater() {
	//TODO implement me
	panic("implement me")
}

func (m MakeCoffee) Brew() {
	//TODO implement me
	panic("implement me")
}

func (m MakeCoffee) PourInCup() {
	//TODO implement me
	panic("implement me")
}

func (m MakeCoffee) AddThings() {
	//TODO implement me
	panic("implement me")
}

func (m MakeCoffee) WantAddThings() bool {
	return false
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	makeCoffee.b = makeCoffee
	return makeCoffee
}

type MakeTea struct {
	template
}

func (m MakeTea) BoilWater() {
	//TODO implement me
	panic("implement me")
}

func (m MakeTea) Brew() {
	//TODO implement me
	panic("implement me")
}

func (m MakeTea) PourInCup() {
	//TODO implement me
	panic("implement me")
}

func (m MakeTea) AddThings() {
	//TODO implement me
	panic("implement me")
}

func (m MakeTea) WantAddThings() bool {
	//TODO implement me
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea

	return makeTea
}

func main() {
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()
	fmt.Println("-------------------")
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
