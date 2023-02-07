package main

import "fmt"

func main() {
	// 固定长度的数组，传值时是值传递(修改方法里的数据，原始数据不受影响)
	var myArray1 [10]int

	for i := 0; i < 10; i++ {
		myArray1[i] = i
	}
	fmt.Println(myArray1)

	 myArray2 := [10]int{1,2,3,4}
	 fmt.Println(myArray2)

	 // 查看数组的数据类型
	 fmt.Printf("myArray1 Types = %T\n" , myArray1)
	 fmt.Printf("myArray2 Types = %T\n" , myArray2)

}
