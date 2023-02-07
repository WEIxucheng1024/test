package main

import "fmt"

var a int

var (
	b int
	c int
)

func main()  {
	a = 2
	fmt.Println("a = ", a)
	fmt.Printf("type fo a = %T\n", a)

	b =3
	fmt.Println("b = ", b)
	fmt.Printf("type fo b = %T\n", b)

	var c = 20
	fmt.Println("c = ", c)
	fmt.Printf("type fo c = %T\n", c)

	var d = "abcd"
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	e := 222
	fmt.Printf("e = %v type of e = %T\n", e, e)

	f := 1.234
	fmt.Printf("f = %v type of f = %T\n", f, f)

}