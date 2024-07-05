package main

import "fmt"

// 全局变量，仅能使用方法一到三来进行声明
var aa = "我是全局变量"

func main() {
	/*
		局部声明的四种方式
	*/

	//方法一： 声明变量，初始值默认是该类型对应的空值
	var a int
	var b string = "你好，我是"
	fmt.Println(b, a)

	//方法二：声明一个变量，初始化一个值
	var c int = 100
	fmt.Printf("type of c: %T\n", c)
	fmt.Println("c=", c)

	//方法三：初始化时可以省去数据类型，编译器根据初始化值自动匹配数据类型
	var d = "你是？"
	fmt.Printf("type of d: %T\n", d)
	fmt.Printf("d= %s\n", d)

	//方法四：(常用方法)省略var关键字，直接自动匹配
	// := 只能够在函数体内声明
	e := 100
	fmt.Printf("e= %d,type of e: %T\n", e, e)

	g := 3.14159
	fmt.Printf("g= %f,type of g: %T\n", g, g)

	fmt.Println("aa= ", aa)

	// 声明多变量
	// 同类型赋值
	var xx, yy int = 100, 200
	fmt.Println(xx, yy)
	// 不同类型赋值，无需定义数据类型
	var zz, ww = 300, "sss"
	fmt.Println(zz, ww)

	//多行的多变量声明
	var (
		vv int  = 400
		jj bool = true
	)
	fmt.Println(vv, jj)

}
