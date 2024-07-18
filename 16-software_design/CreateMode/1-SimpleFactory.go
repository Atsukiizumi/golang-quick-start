package main

import "fmt"

// =======抽象层========
type Fruit interface {
	Show()
}

// =======基础模块===========
type Apple struct {
	Fruit
}

type Banana struct {
	Fruit
}

type Pear struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("I am Apple")
}

func (banana *Banana) Show() {
	fmt.Println("I am Banana")
}

func (pear *Pear) Show() {
	fmt.Println("I am Pear")
}

// ==========工厂模块==============
type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = &Banana{}
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}
func main() {
	factory := &Factory{}

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
