package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap，是一个引用传递
	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["Canada"] = "Toronto"
}

func main() {
	//声明并初始化
	cityMap := map[string]string{
		"China":             "Beijing",
		"Japan":             "Tokyo",
		"The United States": "Washington, D.C.",
	}

	//添加
	cityMap["The United Kingdom"] = "Ford"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "Japan")

	fmt.Println("================")
	printMap(cityMap)

	//修改
	cityMap["The United Kingdom"] = "London"

	fmt.Println("================")
	printMap(cityMap)

	ChangeValue(cityMap)
	fmt.Println("================")
	printMap(cityMap)
}
