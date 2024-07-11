package main

import "fmt"

func main() {
	var a string
	// pair<statictype:string, value:"color">
	a = "color"

	// pair<type:interface{}, value:"color">
	var allType interface{}
	allType = a

	// "类型断言"的机制
	str, _ := allType.(string)
	fmt.Println(str)

}
