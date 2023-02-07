package main

import "fmt"

func printArray(myArray []int){
	for _, i := range myArray {
		fmt.Println("value = ", i)
	}
	myArray[0] = 100
}

func main() {
	// []内不含数字的时候，标识为动态数组
	myArray := []int{1,2,3,4}

	fmt.Printf("myArray Type = %T\n", myArray)

	printArray(myArray)
	// sliec为引用传递，方法内修改数据原始数据会跟这变
	fmt.Println(myArray)

}
