// golang编译顺序
// 1.寻找main包
package main //包名，只有包名为main才能作为启动包

// 2.寻找其他包，如果有包则进入该包中，重复步骤2
// 导入包
import (
	"fmt"
	"time"
)

// 3.初始化const

// 4.初始化var

// 5.初始化init()函数，非main包则返回上一级继续初始化

// 6.执行main()函数
// main函数
func main() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
	fmt.Println("I sleep 1 sec.")
}
