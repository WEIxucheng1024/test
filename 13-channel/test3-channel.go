package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close可以关闭channel
		close(c)
	}()

	//for{
	//	// ok如果为true，那么channel未关闭，为false则关闭
	//	if data, ok := <- c; ok {
	//		fmt.Println(data)
	//	}else {
	//		fmt.Println("data = ",data)
	//		break
	//	}
	//}

	// 可以使用range来迭代不断操作channel
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Main Finished...")
}
