package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 有缓冲channel

	fmt.Println("len(c) = ",len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子goroutine结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子goroutine正在运行，发送的元素为：i = ",i ,"len(c) = ",len(c)," cap(c) = ",cap(c))
		}
	}()

	time.Sleep(2*time.Second)

	for i := 0; i < 4; i++ {
		num := <- c // 从C中取元素
		fmt.Println("num = ", num)
	}

	//time.Sleep(1*time.Second)

	fmt.Println("main 结束")
}
