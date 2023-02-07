package main

import "fmt"



func main() {

	defer fmt.Println("999")

	defer fmt.Println("333")

	defer fmt.Println("555")

	defer func() {
		fmt.Println("aaaaaaaaaa")
	}()

	fmt.Println("111")
	fmt.Println("222")
}
