package main

import "fmt"

// 声明一种新的数据类型，是int的一种别名
type myint int

type Book struct {
	 title string
	 auth string
}

func printBook(b Book) {
	b.auth = "李四"
}

func printBook2(b *Book) {
	b.auth = "李四"
}

func main()  {
	/*
	var a myint
	a = 10
	fmt.Printf("Type of a = %T\n",a)

	 */

	var Book1 Book
	Book1.title = "golang"
	Book1.auth = "张三"

	fmt.Printf("%v\n", Book1)
	printBook(Book1)
	fmt.Printf("%v\n", Book1)

	printBook2(&Book1)
	fmt.Printf("%v\n", Book1)

}
