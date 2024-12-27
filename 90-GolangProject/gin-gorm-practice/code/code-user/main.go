package main

import "fmt"

func main() {
	// 标准的输入
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Printf("a=%d, b=%d\n", a, b)
}
