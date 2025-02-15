package main

import "fmt"

// 闭包实现
func main() {
	grow := Exp(2)

	for i := range 10 {
		fmt.Printf("2^%d=%d\n", i, grow())
	}
}

// 闭包
func Exp(n int) func() int {
	e := 1
	return func() int {
		temp := e
		e *= n
		return temp
	}
}

type Para map[string]interface{}

type Show struct {
	Para
}
