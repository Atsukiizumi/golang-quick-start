package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类

}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}
func (this Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this Dog) Sleep() {
	fmt.Println("Dog is sleeping...")
}
func (this Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color:", animal.GetColor())
	fmt.Println("type:", animal.GetType())
}

/*
多态的基本要素"

	1.有一个父类（有接口）
	2.有子类（实现了父类的全部接口方法）
	3.父类类型的变量（指针）指向（引用）子类的具体数据变量
*/
func main() {
	var animal AnimalIF //接口的数据类型，父类指针
	animal = &Cat{"Green"}
	animal.Sleep() //调用的是cat的sleep

	animal = &Dog{"Blue"}
	animal.Sleep() //调用的是Dog的sleep

	cat := Cat{"Red"}
	dog := Dog{"Blue"}
	showAnimal(&cat)
	showAnimal(&dog)
}
