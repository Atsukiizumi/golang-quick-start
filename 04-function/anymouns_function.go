package main

// 匿名函数实例
func main() {
	// 匿名函数
	func(a, b int) int {
		return a + b
	}(1, 2)

	// 危险写法
	//people := []Pesss{
	//	newPesss("A", 11, 123),
	//	newPesss("B", 12, 123),
	//	newPesss("C", 13, 123),
	//}

	//fmt.Println(people)
	//
	//slices.SortFunc(people, func(p1, p2 Pesss) int {
	//	if p1.Name > p2.Name {
	//		return 1
	//	} else if p1.Name < p2.Name {
	//		return -1
	//	}
	//	return 0
	//})
	//
	//fmt.Println(people)
}

type Pesss struct {
	Name   string
	Age    int
	Salary float64
}

func newPesss(name string, age int, salary float64) Pesss {
	return Pesss{
		Name:   name,
		Age:    age,
		Salary: salary,
	}
}
