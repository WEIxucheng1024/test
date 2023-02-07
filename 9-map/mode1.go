package main

import "fmt"

func printMap(cityMap map[string]string)  {
	for s, s2 := range cityMap {
		fmt.Println(s)
		fmt.Println(s2)
	}
	cityMap["USA"] = "newYork"

}

func main() {
	cityMap := map[string]string{
		"China":"beijing",
		"Japan":"dongjing",
		"USA" : "newYork",
	}

	for s, s2 := range cityMap {
		fmt.Println(s)
		fmt.Println(s2)
	}

	delete(cityMap,"Japan")
	fmt.Println(cityMap)

	cityMap["USA"] = "华盛顿"
	fmt.Println(cityMap)

	printMap(cityMap)
	fmt.Println(cityMap)

}
