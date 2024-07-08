package main

import "fmt"

func main() {
	// 第一种声明方式
	// 声明map1是一种map类型，key是string，value是int
	var map1 map[string]int
	if map1 == nil {
		fmt.Println("map1 is nil")
	}

	//在使用map，前首先用make给map分配空间
	map1 = make(map[string]int, 5)

	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	//可以直接打印，也可以使用索引
	fmt.Println(map1)

	// 第二种声明方式
	map2 := make(map[string]string)
	map2["one"] = ".net"
	map2["two"] = "c++"
	map2["three"] = "c"
	fmt.Println(map2)

	// 第三种声明方式
	map3 := map[string]int{
		"one":   4,
		"two":   5,
		"three": 6,
	}
	fmt.Println(map3)
}
