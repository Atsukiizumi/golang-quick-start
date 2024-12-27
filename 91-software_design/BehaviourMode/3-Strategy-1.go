package main

import (
	"fmt"
)

// 武器策略（抽象的策略）
type WeponStrategy interface {
	UseWeapon() //使用武器
}

// 具体的策略
type Ak47 struct {
}

func (a *Ak47) UseWeapon() {
	fmt.Println("使用Ak47 战斗")
}

// 具体策略
type Knife struct {
}

func (k *Knife) UseWeapon() {
	fmt.Println("使用匕首 战斗")
}

// 环境类
type Hero struct {
	strategy WeponStrategy //拥有一个抽象的策略
}

// 设置一个策略
func (h *Hero) SetWeaponStrategy(s WeponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon() //调用策略
}

func main() {
	hero := new(Hero)
	//更换策略
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
