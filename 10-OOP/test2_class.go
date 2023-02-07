package main

import "fmt"

type Hero struct{
	Name string
	Ad int
	Level int
}

func (h Hero) GetName() {
	fmt.Println(h.Name)
}

func (h *Hero) SetName(newName string) {
	h.Name = newName
}

func main() {
	hero := Hero{
		Name: "张三",
		Ad: 12,
		Level: 22,
	}
	hero.GetName()
	hero.SetName("李四")
	fmt.Println(hero )
}