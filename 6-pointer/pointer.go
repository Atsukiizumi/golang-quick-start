package main

import (
	"fmt"
)

func swap(a int, b int) {
	fmt.Println("start to swap value...")
	var temp int
	temp = a
	b = a
	a = temp
}

func swap_point(pa *int, pb *int) {
	fmt.Println("start to swap point...")
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("a=", a, "b=", b)
	swap(a, b)
	fmt.Println("a=", a, "b=", b)
	swap_point(&a, &b)
	fmt.Println("a=", a, "b=", b)

	var p *int //一级地址
	p = &a
	fmt.Println("p=", p, "*p=", *p, "&a=", &a)

	var pp **int //二级地址
	pp = &p
	fmt.Println("pp=", pp, "*pp=", *pp, "**pp=", **pp, "&p=", &p)
}
