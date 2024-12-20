package main

import "fmt"

type MyInt int

func (i *MyInt) Set(val int) {
	*i = MyInt(val) // 修改了，但是不会造成任何影响
}

func main() {
	myInt := MyInt(1)
	myInt.Set(2)
	fmt.Println(myInt)
}
