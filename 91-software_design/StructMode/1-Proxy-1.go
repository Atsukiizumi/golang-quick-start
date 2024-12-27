package main

import (
	"fmt"
)

type Goods struct {
	Kind string //商品种类
	Fact bool   // 商品真伪
}

// ===== 抽象层 =====
// 抽象的购物主题Subject
type Shopping interface {
	Buy(goods *Goods)
}

// ===== 实现层 =====
// 具体的购物主题，实现了Shopping
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物，买了", goods.Kind)
}

type AmericaShopping struct{}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物，买了", goods.Kind)
}

type JapanShopping struct{}

func (js *JapanShopping) Buy(goods *Goods) {
	fmt.Println("去日本进行了购物，买了", goods.Kind)
}

// 海外代理
type OverseasProxy struct {
	shopping Shopping
}

func (op *OverseasProxy) Buy(goods *Goods) {
	//1. 先验货
	if op.distinguish(goods) == true {
		//2. 进行购买
		op.shopping.Buy(goods)
		//3. 海关安检
		op.check(goods)
	}

}

// 验货流程
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪")
	if goods.Fact != true {
		fmt.Println("发现假货", goods.Kind, "，不应该购买。")
	}
	return goods.Fact
}

// 安检流程
func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "]进行了海关检查，成功的带回祖国")
}

// 创建一个代理，并配置关联被代理的主题
func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "HK416",
		Fact: false,
	}

	g3 := Goods{
		Kind: "乌冬面",
		Fact: true,
	}

	// 不使用代理来完成从韩国购买任务
	var kshop Shopping = new(KoreaShopping)
	var ashop Shopping = new(AmericaShopping)
	var jshop Shopping = new(JapanShopping)
	if g1.Fact == true {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪")
		kshop.Buy(&g1)
		fmt.Println("对[", g1.Kind, "]进行了海关检查，成功的带回祖国")
	}

	fmt.Println("----------------以下 使用 代理模式--------------")
	var overseasProxy Shopping
	overseasProxy = NewProxy(ashop)
	overseasProxy.Buy(&g2)
	overseasProxy = NewProxy(jshop)
	overseasProxy.Buy(&g3)
}
