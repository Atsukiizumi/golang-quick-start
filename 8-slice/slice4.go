package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len=3, cap=3
	fmt.Println("len = %d, cap = %d, slice=%v\n", len(s), cap(s), s)

	//[0,2]取两个元素
	s1 := s[0:2]
	fmt.Println("s1 = %v", s1)
}
