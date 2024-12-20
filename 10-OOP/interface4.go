package main

import "fmt"

type Person interface {
	Say(s string) string
	Walk(i int)
}

type Func func()

func (f Func) Say(s string) string {
	f()
	return "bipabipa"
}

func (f Func) Walk(i int) {
	f()
	fmt.Println("can not walk")
}

func main() {
	var function Func
	function = func() {
		fmt.Println("do something")
	}

	fmt.Println(function.Say("what?"))
	function.Walk(1)

	function()
}
