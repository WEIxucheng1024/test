package main

import "fmt"

func main() {

	slice1 := make([]int,3,5)
	fmt.Printf("slice1 长度：%d, 容量为：%d，内容为：%v\n", len(slice1),cap(slice1),slice1)

	slice2 := append(slice1,1)
	fmt.Printf("slice2 长度：%d, 容量为：%d，内容为：%v\n", len(slice2),cap(slice2),slice2)

	slice3 := append(slice2,2)
	fmt.Printf("slice3 长度：%d, 容量为：%d，内容为：%v\n", len(slice3),cap(slice3),slice3)

	slice4 := append(slice3, 3)
	fmt.Printf("slice4 长度：%d, 容量为：%d，内容为：%v\n", len(slice4),cap(slice4),slice4)

	var numbers = make([]int,3)
	fmt.Printf("numbers 长度：%d, 容量为：%d，内容为：%v\n", len(numbers),cap(numbers),numbers)

	numbers1 := append(numbers,3)
	fmt.Printf("numbers1 长度：%d, 容量为：%d，内容为：%v\n", len(numbers1),cap(numbers1),numbers1)

}
