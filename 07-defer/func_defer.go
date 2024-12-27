package main

import "fmt"

func main() {
	fmt.Println(sum(3, 5))
}

func sum(a, b int) (s int) {
	defer func() {
		s -= 10
	}()
	s = a + b
	return
}
