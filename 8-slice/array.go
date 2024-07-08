package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println("arr:", arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", arr2)

	arr3 := [10]int{1}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for index, value := range arr2 {
		fmt.Println(index, value)
	}
}
