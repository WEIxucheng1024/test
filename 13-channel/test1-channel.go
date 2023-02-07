package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")

		return
		c <- 666
	}()

	//num := <- c
	//
	//fmt.Println(num)
	time.Sleep(1*time.Second)
	fmt.Println("main结束")
}
