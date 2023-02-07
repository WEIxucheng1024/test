package main

import "fmt"

func myFunc(arg interface{}){
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	s, ok := arg.(book)
	if !ok {
		fmt.Println("arg不是book")
	}else{
		fmt.Println("arg为book类型")
		fmt.Printf("类型为：%T\n", s)
	}
}

type book struct {
	auth	string

}

func main() {
	b := book{"Golang"}
	myFunc(b)

	myFunc("emmmm")
	myFunc(121)
}
