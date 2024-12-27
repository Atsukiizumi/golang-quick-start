package main

import "fmt"

//简单工厂模式  + “开闭原则” =    工厂方法模式

// ============ 抽象层 =============
// 水果类（抽象接口）
type Fruits interface {
	Show()
}

// 工厂类（抽象接口）
type AbstractFactoryF interface {
	CreateFruit() Fruits
}

// ======基础类模块=======
type Apples struct {
	Fruits
}

type Bananas struct {
	Fruits
}

type Pears struct {
	Fruits
}

func (apples *Apples) Show() {
	fmt.Println("苹果")
}

func (bananas *Bananas) Show() {
	fmt.Println("香蕉")
}

func (pears *Pears) Show() {
	fmt.Println("梨")
}

// ======== 工厂模块 =========
type AppleFactory struct {
	AbstractFactoryF
}

func (fac *AppleFactory) CreateFruit() Fruits {
	var fruits Fruits

	fruits = new(Apples)

	return fruits
}

func main() {
	//先需要一个工厂
	var appleFac AbstractFactoryF
	appleFac = new(AppleFactory)
	//生产对应的水果
	var apples Fruits
	apples = appleFac.CreateFruit()
	apples.Show()
}
