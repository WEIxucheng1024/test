package main

import "fmt"

// const 来定义枚举类型
const (
	// 可以在const()添加关键字iota，每行的iota都会累加1，第一行的iota默认值为0
	Beijing = 5
	Shanghai = iota
	Shenzhen
)

const (
	a, b = iota + 1, iota + 2		// iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d							// iota = 1，c = iota + 1, d = iota + 2, c = 2 , d = 3
	e, f							// iota = 2, e = iota + 1, f = iota + 2， e = 3, f = 4

	g, h = iota * 2, iota * 3		// iota = 3, g = iota * 2, g = iota * 3, g = 6, h = 9
	i, k							// iota = 4, i = iota * 2, k = iota * 3, i = 8, k = 12
)



func main() {
	const length int = 10

	fmt.Println(length)

	fmt.Println(Beijing,Shanghai,Shenzhen)

	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(i, k)

}
