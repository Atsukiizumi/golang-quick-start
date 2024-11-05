package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "你好，世界!"
	// 未使用rune类型
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d. %d,%x,%s\n\n", i, str[i], str[i], string(str[i]))
	}
	// 使用rune类型
	for i, r := range str {
		fmt.Printf("%d. %d,%x,%s\n", i+1, r, r, string(r))
	}

	// 字符串转换成[]rune
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%d. %d,%x,%s\n\n", i+1, runes[i], runes[i], string(runes[i]))
	}

	// 使用utf8包
	for a, w := 0, 0; a < len(str); a += w {
		r, width := utf8.DecodeRuneInString(str[a:])
		fmt.Println(string(r))
		w = width
	}
}
