package main

import "fmt"

func main() {
	//len是尾部指针指向合法元素最后一个元素的长，而cap是指切片系统声明的长度
	//make(slice,len,cap)
	slice1 := make([]int, 3, 5)
	slice1 = append(slice1[:1], -1, 0)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(slice1), cap(slice1), slice1)

	//向一个变量cap已经满的slice，追加元素
	slice1 = append(slice1, 3)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(slice1), cap(slice1), slice1)

	slice2 := make([]int, 3)
	slice2 = append(slice2, 4)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(slice2), cap(slice2), slice2)

}
