package main

import "fmt"

func main()  {
	var a string
	// pair<statictype:string,value:"aceld">
	a = "aceld"

	var allType interface{}
	// pair<type:string,value:"aceld">
	allType = a

	str,_ := allType.(string)
	fmt.Println(str)
}
