package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a=", a, "b=", b)
	c := 100
	return c
}

// 返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a, "b=", b)
	c := 100
	d := 200
	return c, d
}

// 返回多个返回值，带形参
func foo3(a string, b int) (x int, y int) {
	fmt.Println("a=", a, "b=", b)
	// 给形参赋值
	x = 500
	y = 1000
	return x, y
}

func foo4(a string, b int) (x, y int) {
	fmt.Println("a=", a, "b=", b)
	// 给形参赋值
	x = 200
	y = 400
	return x, y
}

func main() {
	c := foo1("hello golang", 10)
	fmt.Println("c=", c)

	x, y := foo2("hello golang", 100)
	fmt.Println("x=", x, "y=", y)

	a, b := foo3("hello golang", 200)
	fmt.Println("a=", a, "b=", b)

	x, y = foo4("hello golang", 300)
	fmt.Println("x=", x, "y=", y)
}
