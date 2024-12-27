package main

import "fmt"

type Season uint8

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

// 通过给自定义类型添加方法来返回其字符串表现形式
func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	}
	return ""
}

func main() {
	fmt.Println(Season(0))
	fmt.Println(Season(1))
	fmt.Println(Season(2))
	fmt.Println(Season(3))
	fmt.Println(Spring, Summer, Autumn, Winter)
}
