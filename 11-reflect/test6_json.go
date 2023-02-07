package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title 		string 			`json:"title"`
	Year 		int 			`json:"year"`
	Price 		int 			`json:"price"`
	Actors 		[]string 		`json:"actors"`
}

func main() {
	m := Movie{
		Title: "aaa",
		Year: 12,
		Price: 2,
		Actors: []string{"bbb","ccc"},
	}

	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(marshal)
	fmt.Printf("%s\n",marshal)

	var m1 Movie
	err = json.Unmarshal(marshal, &m1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(m1)
	fmt.Printf("type : %T\n",m1)
}
