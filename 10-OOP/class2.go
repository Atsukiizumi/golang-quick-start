package main

import "fmt"

type Human struct {
	name string
	age  int
}

type SuperMan struct {
	Human //superman类继承human类
	level int
}

// 重新定义父类的方法Eat()
func (this SuperMan) Eat() {
	fmt.Printf("Superman %s is eating...\n", this.name)
}

func (this SuperMan) BaseInfo() {
	fmt.Printf("Name  is %s\n", this.name)
	fmt.Printf("Age   is %d\n", this.age)
	fmt.Printf("Level is %d\n", this.level)
}

// 子类新方法
func (this SuperMan) Fly() {
	fmt.Printf("Superman %s is flying...\n", this.name)
}

func (this Human) Eat() {
	fmt.Printf("%s is eating...\n", this.name)
}

func (this Human) Walk() {
	fmt.Printf("%s is walking...\n", this.name)
}

func (this *Human) Grow() {
	fmt.Printf("%s is grow up!\n", this.name)
	this.age = this.age + 1
}
func (this Human) BaseInfo() {
	fmt.Println("Name = ", this.name)
	fmt.Println("Age = ", this.age)
}

func main() {
	human := Human{"Tom", 25}
	human.Walk()
	human.Eat()
	human.Grow()
	human.BaseInfo()

	sp := SuperMan{Human{"Gill", 12}, 20}
	sp.Fly()
	sp.Walk()
	sp.Eat()
	sp.Grow()
	sp.BaseInfo()

}
