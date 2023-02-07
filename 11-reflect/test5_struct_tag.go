package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name 	string		`info:"name" emm:"ahahah"`
	Sex		string		`info:"sex"`
}

func Findtag(str interface{}){
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		get := t.Field(i).Tag.Get("info")
		fmt.Println("info:",get)
	}
}

func main() {
	re := &resume{Name: "emm",Sex: "232"}
	Findtag(re)
}
