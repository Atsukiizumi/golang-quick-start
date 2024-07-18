package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("Cat is eating...")
}

// 给小猫添加一个 可以睡觉的方法 （使用继承来实现）
type CatB struct {
	Cat
}

func (c *CatB) Sleep() {
	fmt.Println("CatB is sleeping...")
}

// 给小猫添加一个 可以睡觉的方法 （使用组合的方式）
type CatC struct {
	C *Cat
}

func (c *CatC) Sleep() {
	fmt.Println("CatC is sleeping...")
}

func main() {
	//继承
	cb := new(CatB)
	cb.Sleep()
	cb.Eat()

	//组合
	cc := new(CatC)
	cc.C = new(Cat)
	cc.C.Eat()
	cc.Sleep()
}
