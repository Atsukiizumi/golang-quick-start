package main

import "fmt"

/*
练习：
商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景
*/

// 销售策略
type SellStrategy interface {
	GetPrice(price float64) float64
}

type StrategyA struct {
}

func (s *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("使用策略A...")
	return price * 0.8
}

type StrategyB struct{}

func (s *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("使用策略B...")
	if price >= 200 {
		price -= 100
	}
	return price
}

// 环境类
type Goods struct {
	Price    float64
	Strategy SellStrategy
}

// 设置策略
func (goods *Goods) SetStrategy(s SellStrategy) {
	goods.Strategy = s
}

func (goods *Goods) SellPrice() float64 {
	fmt.Println("原价值", goods.Price, " .")
	return goods.Strategy.GetPrice(goods.Price)
}

func main() {
	nikke := Goods{
		Price: 200.0,
	}
	//执行策略A
	nikke.SetStrategy(new(StrategyA))
	fmt.Println("上午nikke卖，", nikke.SellPrice())

	//执行策略B
	nikke.Price = 260
	nikke.SetStrategy(new(StrategyB))
	fmt.Println("下午nikke卖，", nikke.SellPrice())
}
