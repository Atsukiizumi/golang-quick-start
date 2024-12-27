package main

import "fmt"

func main() {
	fib := Fib(10)

	for n, next := fib(); next; n, next = fib() {
		fmt.Println(n)
	}
}

func Fib(n int) func() (int, bool) {
	a, b, c := 1, 1, 2
	i := 0
	return func() (int, bool) {
		if i >= n {
			return 0, false
		} else if i < 2 {
			f := i
			i++
			return f, true
		}

		a, b = b, c
		c = a + b
		i++
		return a, true
	}

}
