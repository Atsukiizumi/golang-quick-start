package main

import "fmt"

func main() {
	//声明slice1是一个切片，并且初始化
	slice1 := []int{1, 2, 3}

	//声明slice2是一个切片，但是并没有给slice2分配空间
	var slice2 []int
	slice2 = make([]int, 3)
	slice2[0] = 1

	//声明slice3是一个切片，同时给slice3分配空间，3个空间，初始化值是0
	var slice3 = make([]int, 3)

	//声明slice4是一个切片，同时给slice4分配空间，3个空间，初始化值是0，通过:=推导出slice4是一个切片
	slice4 := make([]int, 3)

	fmt.Println("len = %d, slice=%v\n", len(slice1), slice1)
	fmt.Println("len = %d, slice=%v\n", len(slice2), slice2)
	fmt.Println("len = %d, slice=%v\n", len(slice3), slice3)
	fmt.Println("len = %d, slice=%v\n", len(slice4), slice4)

	var slice5 []int
	//判断一个slice是否为0
	if slice5 == nil {
		fmt.Println("slice5 is nil")
	} else {
		fmt.Println("slice5 is not nil")
	}
}
