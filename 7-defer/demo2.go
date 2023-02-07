package main

import "fmt"

func func1(){
	fmt.Println("aaaa")
}

func func2() int {
	fmt.Println("bbbb")
	return 1
}

func deferFunc() int{
	defer func1()
	return func2()
}

func main()  {
	deferFunc()
}
