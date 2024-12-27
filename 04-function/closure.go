package main

import "fmt"

func main() {
	grow := Exp(2)
	for i := range 10 {
		fmt.Printf("2^%d=%d\n", i, grow())
	}
}

func Exp(n int) func() int {
	e := 1
	return func() int {
		temp := e
		e *= n
		return temp
	}
}
