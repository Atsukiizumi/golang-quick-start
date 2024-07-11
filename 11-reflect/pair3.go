package main

import "fmt"

type reader interface {
	readBook()
}

type writer interface {
	writeBook()
}

type book struct {
	name   string
	author string
}

func (this *book) readBook() {
	fmt.Println("read a Book ", this.name)
}

func (this *book) writeBook() {
	fmt.Println("write a Book ", this.name)
}

func main() {
	//b1: pair<type:book, value:book{"Golang", "N/A"}>
	b1 := book{"Golang", "N/A"}

	//r: pair<type: ,value: >
	var r reader
	//r: pair<type:book, value:book{"Golang", "N/A"}>
	r = &b1
	r.readBook()

	//w: pair<type: ,value: >
	var w writer
	//w: pair<type:book, value:book{"Golang", "N/A"}>
	w = r.(writer) // w 和 r 的类型断言为什么能成功？因为他们的具体类型是一致的
	w.writeBook()

}
