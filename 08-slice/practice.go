package main

import "fmt"

func arryQuestion() {
	arr := []int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[i]+arr[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}

func selectQuestion() {
	arr := []int{9, 8, 7, 6, 5, 4}
	smaller := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if smaller > arr[j] {
				smaller = arr[j]
				index = j
			}
		}
		arr[index] = arr[0]
		arr[0] = smaller
	}

}

func main() {
	arryQuestion()
	selectQuestion()
}
