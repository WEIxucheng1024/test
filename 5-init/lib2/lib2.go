package lib2

import "fmt"

func Lib2API(){
	fmt.Println("Lib2方法被执行")
}

func init(){
	fmt.Println("lib2包init方法被执行")
}