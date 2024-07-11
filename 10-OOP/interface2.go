package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 该如何区分 此时引用的底层数据类型到底是什么？

	//给interface{}提供"类型断言"的机制
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is a string")
		fmt.Println("value=", value)
	} else {
		fmt.Println("arg is not a string")
	}

}

type Book2 struct {
	auth string
}

func main() {
	book := Book2{auth: "Tom"}
	myFunc(book)
	myFunc("I'm string")
	myFunc(12)
	myFunc(12.2222)
}
