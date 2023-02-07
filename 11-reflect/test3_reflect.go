package main

import (
	"fmt"
	"reflect"
)

func reflectNum(num interface{}) {
	fmt.Println("type = ",reflect.TypeOf(num))
	fmt.Println("value = ",reflect.ValueOf(num))

}

func main() {
	var num float64 = 3.123123
	reflectNum(num)
}
