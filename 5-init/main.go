package main

import (
	// _ 代表匿名，此时golang可以让不调用的方式导入该包
	// 这种方式导入只有执行init()方法
	_ "awesomeProject/5-init/lib1"
	//mylib "awesomeProject/5-init/lib2"
	// . 代表可省略包名直接使用包内方法
	. "awesomeProject/5-init/lib2"
)

func main() {
	//lib1.Lib1Test()
	Lib2Test()
}
