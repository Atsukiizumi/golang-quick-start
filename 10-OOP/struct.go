package main

import (
	"fmt"
)

// 声明一种新的数据类型myint，是int的一个别名
type myint int

type Book struct {
	title  string
	author string
}

func printBook(book Book) {
	fmt.Printf("Book's title is %s\n", book.title)
	fmt.Printf("Book's author is %s\n", book.author)
}

func updateBook(book Book, title string, author string) {
	//传递一个副本，非引用
	book.title = title
	book.author = author
}

func changeBook(book *Book, title string, author string) {
	//传递一个副本，引用
	book.title = title
	book.author = author
}

func main() {
	var a myint = 1
	fmt.Println("a = ", a)
	fmt.Printf("typeof %T\n", a)

	// 结构体实例化三种方式
	// 1.var声明
	var book1 Book
	book1.title = "Go Program"
	book1.author = "www.baidu.com"
	printBook(book1)

	// 2.new/&引用分配内存
	book2 := new(Book)
	book2.title = "Python"
	book2.author = "www.baidu.com"

	book2 = &Book{
		"Python Program",
		"www.baidu.com",
	}
	printBook(*book2)

	// 3.赋值初始化
	book3 := Book{
		"Java",
		"www.baidu.com",
	}
	
	book3 = Book{
		title:  "Java",
		author: "www.baidu.com",
	}
	printBook(book3)

	updateBook(book1, "Golang's roadmap", "N/A")
	printBook(book1)

	changeBook(&book1, "Golang's roadmap", "N/A")
	printBook(book1)
}
