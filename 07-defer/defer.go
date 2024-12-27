package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	defer fmt.Println("main::defer 1")
	defer fmt.Println("main::defer 2")

	fmt.Println("main::log 1")
	fmt.Println("main::log 2")

	defer func1()
	defer func2()
	defer func3()
}
