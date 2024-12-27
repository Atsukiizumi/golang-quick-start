package main

import "fmt"

// const 来定义枚举类型
const (
	//可以在const()添加一个关键字iota，每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	// 常量（只读属性）
	const length int = 10

	fmt.Println("length= ", length)

	fmt.Println("beijing=", BEIJING)
	fmt.Println("shanghai=", SHANGHAI)
	fmt.Println("shenzhen=", SHENZHEN)

	fmt.Println("a=", a, " b=", b)
	fmt.Println("c=", c, " d=", d)
	fmt.Println("e=", e, " f=", f)
	fmt.Println("g=", g, " h=", h)
	fmt.Println("i=", i, " j=", j)
}
