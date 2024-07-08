package main

import "fmt"

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Println("value = ", v)
	}

	arr[0] = 100
}

func main() {
	arr := []int{1, 2, 3, 4, 5} //动态数组，切片，slice
	fmt.Printf("arr type is %T\n", arr)

	printArr(arr) // 动态数组在传参时传引用值
	fmt.Println("------------")
	for _, v := range arr {
		fmt.Println("value = ", v)
	}

}
