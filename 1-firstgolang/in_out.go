package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 不使用缓存包读
	var s, s2, s3 string
	//fmt.Scan(&s, &s2)
	//fmt.Scanln(&s, &s2)
	scanf, err := fmt.Scanf("%s %s \n %s", &s, &s2, &s3)
	if err != nil {
		fmt.Println(scanf, err)
	}
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

	// 读
	scanner := bufio.NewScanner(os.Stdin) // 使用缓存包
	scanner.Scan()
	fmt.Println(scanner.Text())

	// 写
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello world!\n")
	writer.Flush() // 清空
	fmt.Println(writer.Buffered())
}
