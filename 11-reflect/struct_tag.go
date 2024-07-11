package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `info:"name"`
	Age  int    `info:"age"`
	Sex  string `info:"sex"`
}

func FindTag(str interface{}) {
	//Elem()获取当前指针所指的结构体
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("info")
		fmt.Println("info:", tag)
	}
}

func main() {
	st1 := Student{"wo", 12, "man"}

	FindTag(&st1)
}
