package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	fmt.Printf("slice1切片长度：%d, 内容为：%v\n", len(slice1), slice1)

	var slice2 []int
	slice2 = make([]int,10)

	slice2 = append(slice2, 21)
	slice2[1] = 2121
	fmt.Printf("slice2切片长度：%d, 内容为：%v\n", len(slice2), slice2)

	var slice3 []int = make([]int,3)
	// slice4 := make([]int,5)
	fmt.Printf("slice3切片长度：%d, 内容为：%v\n", len(slice3), slice3)

	var slice5 []int
	if slice5 == nil {
		fmt.Printf("切片为空")
	}

}