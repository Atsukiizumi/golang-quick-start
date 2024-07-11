package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//关闭channel,如果不关go会报错死锁
		close(c)
	}()

	/*for {
		//ok如果为true表示channel没有关闭，如果为false表示已经关闭
		if data, ok := <-c; ok {
			fmt.Println("data =", data)
		} else {
			fmt.Println("channel closed")
			break
		}
	}*/

	//可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("end")
}
