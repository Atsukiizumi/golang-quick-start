package main

import "fmt"

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写，表示该属性是对外能够访问的
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("Hero's Name:", this.Name)
	fmt.Println("Hero's Ad:", this.Ad)
	fmt.Println("Hero's Level:", this.Level)
}

func (this Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

func (this *Hero) SetName(name string) {
	//this 是调用该方法的对象的一个副本
	this.Name = name
}

func main() {
	//创建一个对象
	hero := Hero{Name: "Superman", Ad: 10, Level: 1}

	hero.Show()

	hero.SetName("Duck")
	hero.Show()

}
