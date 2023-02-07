package lib1

import "fmt"

func Lib1API(){
	fmt.Println("Lib1方法被执行")
}

func init(){
	fmt.Println("lib1包init方法被执行")
}
