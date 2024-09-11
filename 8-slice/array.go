package main

import "fmt"

// 数组在传参的时候，golang会分配新的内存存储数组的值
func main() {
	var arr [10]int
	fmt.Println("arr:", arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", arr2)

	arr3 := [10]int{1}
	fmt.Println("arr3's len:", len(arr3))
	fmt.Println("arr3's cap: ", cap(arr3))
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//golang编译器识别长度，声明后无法自动扩充
	arr4 := [...]string{"Text", "A"}

	fmt.Println("arr4:", arr4)

	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	arr5 := [5]int{1, 2, 3, 4, 5}
	s1 := arr5[0:3]
	fmt.Printf("%v", s1)
	fmt.Printf("%T\n", s1)

	//查看数组类型
	fmt.Printf("arr type is %T\n", arr)
	fmt.Printf("arr2 type is %T\n", arr2)
	fmt.Printf("arr3 type is %T\n", arr3)
	fmt.Printf("arr4 type is %T\n", arr4)
}
