package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	//基本反射用法
	fmt.Println("type of arg:", reflect.TypeOf(arg))
	fmt.Println("value of arg:", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("call user:", this.Name)
	fmt.Printf("%v\n", this)
}

func main() {
	var num float32 = 3.1415926
	reflectNum(num)

	user := User{1, "Aceld", 12}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type:", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value:", inputValue)

	//通过type 获取里面的字段
	//1. 获取interface{}的reflect.Type，通过Type得到NumField，进行遍历
	//2. 得到每个field，数据类型
	//3. 通过filed有一个interface{}方法得到 对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v %v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s:%v\n", method.Name, method.Type)
	}
}
