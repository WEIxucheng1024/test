package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID string
	Name string
	Age	int
}

func (this User) Call() {
	fmt.Println("User is called...")
	fmt.Println(this)
}

func main()  {
	user := User{ID:"adad",Name: "小明",Age: 12}
	DofiledAndMethod(user)

}

func DofiledAndMethod(i interface{}) {
	iType := reflect.TypeOf(i)
	fmt.Println("Type:",iType) 
	iValue := reflect.ValueOf(i)
	fmt.Println("value:",iValue)

	for i := 0; i < iType.NumField(); i++ {
		field := iType.Field(i)
		fmt.Println(field)
		value := iValue.Field(i)
		fmt.Println(value)
	}

	for i := 0; i <iType.NumMethod(); i++ {
		method := iType.Method(i)
		fmt.Printf("%s:%v\n",method.Name,method.Type)
	}
}
