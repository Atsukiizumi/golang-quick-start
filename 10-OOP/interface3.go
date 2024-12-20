package main

import "fmt"

// 起重机接口
type Crane interface {
	JackUp() string
	Hoist() string
}

// 起重机A
type CraneA struct {
	work int
}

func (this CraneA) Work() {
	fmt.Println("使用起重机A")
}

func (this CraneA) JackUp() string {
	this.Work()
	return "jackup"
}

func (this CraneA) Hoist() string {
	this.Work()
	return "hoist"
}

type CraneB struct {
	work int
}

func (this CraneB) Boot() {
	fmt.Println("使用起重机B")
}

func (this CraneB) JackUp() string {
	this.Boot()
	return "jackup"
}

func (this CraneB) Hoist() string {
	this.Boot()
	return "hoist"
}

type ConstructionCompany struct {
	Crane Crane // 只根据Crane类型存放起重机
}

func (this *ConstructionCompany) Build() {
	fmt.Println(this.Crane.JackUp())
	fmt.Println(this.Crane.Hoist())
	fmt.Println("建筑完成")
}

func main() {
	//使用起重机A
	company := ConstructionCompany{CraneA{}}
	company.Build()
	fmt.Println()
	//使用起重机B
	company.Crane = CraneB{}
	company.Build()
	fmt.Println("Done.")
}
