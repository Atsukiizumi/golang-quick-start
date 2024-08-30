package main

import "fmt"

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

	var book1 Book
	book1.title = "Go Program"
	book1.author = "www.baidu.com"
	printBook(book1)

	updateBook(book1, "Golang's roadmap", "N/A")
	printBook(book1)

	changeBook(&book1, "Golang's roadmap", "N/A")
	printBook(book1)
}
