package main

import "fmt"

func main() {

	//group.Add(10)
	// map
	mp := make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			// 写操作
			for i := 0; i < 100; i++ {
				mp["helloworld"] = 1
			}
			// 读操作
			for i := 0; i < 10; i++ {
				fmt.Println(mp["helloworld"])
			}
			//group.Done()
		}()
	}
	//group.Wait()
}
