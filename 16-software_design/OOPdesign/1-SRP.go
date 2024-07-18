package main

//单一职责原则

import "fmt"

type ClothesShop struct {
}

func (cs *ClothesShop) onShop() {
	fmt.Println("休闲的")
}

type ClothesWork struct {
}

func (cs *ClothesWork) onWork() {
	fmt.Println("工作的装扮")
}

func main() {
	cw := new(ClothesWork)
	cw.onWork()

	cs := new(ClothesShop)
	cs.onShop()
}
