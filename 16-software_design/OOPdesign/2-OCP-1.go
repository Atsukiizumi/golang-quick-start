package main

import "fmt"

type Banker struct {
}

func (b *Banker) Save() {
	fmt.Println("存款业务")
}

func (b *Banker) Transfer() {
	fmt.Println("转账业务")
}

func (b *Banker) Pay() {
	fmt.Println("支付业务")
}

// 平铺式设计
func main() {
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()
}
