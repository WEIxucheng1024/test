package main

import "fmt"

func foo1(a, b int) int {
	fmt.Println(a)
	fmt.Println(b)

	c := 100

	return c
}

func foo2(a string, b int ) (int, int){
	fmt.Println(a)
	fmt.Println(b)

	return 222, b
}

func foo3(a string, b int) (i1 int, i2 int){
	fmt.Println(a)
	i2 = b
	i1 = 123
	return
}

func foo4(a string, b int)(i1, i2 int){
	fmt.Println(a)
	fmt.Println(i1, i2)
	i2 = 111
	i1 = b
	return
}

func main() {
	c := foo1(11,22)
	fmt.Println(c)

	i, i2 := foo2("emmm", 333)
	fmt.Println(i,i2)

	a, b := foo3("aaaa", 321)
	fmt.Println(a, b)

	a, b = foo4("bbb", 999)
	fmt.Println(a, b)
}


