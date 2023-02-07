package main

import "fmt"

func main() {

	var myMap map[string]string
	if myMap == nil {
		fmt.Printf("为空\n")
	}

	myMap1 := make(map[string]string, 10)
	myMap1["ones"] = "111"
	myMap1["two"] = "222"
	myMap1["three"] = "333"
	fmt.Println(myMap1)

	myMap2 := make(map[int]int)
	myMap2[1] = 1
	myMap2[2] = 2
	myMap2[3] = 3
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"111": "111",
		"222": "222",
		"333": "333",
	}
	fmt.Println(myMap3)
	myMap3["444"] = "444"
	fmt.Println(myMap3)
}
