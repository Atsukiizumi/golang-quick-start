package main

import "fmt"

func main() {

	c := make(chan int, 3) //带有缓冲的channel

	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < cap(c); i++ {
			c <- i
			fmt.Println("子go正在运行，发送的元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	for i := 0; i < cap(c); i++ {
		num := <-c //从c中接收数据，并赋予给num
		fmt.Println("num=", num)
	}

	fmt.Println("main goroutine 结束")
}
