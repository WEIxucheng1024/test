package main

import "fmt"

type Human struct {
	Name string
	Sex	string
}

func (this *Human) Eat()  {
	fmt.Println("Human.Eat...")
}

func (this Human) Walk() {
	fmt.Println("Human.Walk...")

}

type SuperMan struct {
	Human
	Level int
}

func main() {
	human := Human{
		Name: "张三",
		Sex: "男",
	}

	human.Eat()
	human.Walk()

	sp := SuperMan{Human:Human{Name: "张三", Sex: "男",}, Level: 12,}
	fmt.Println(sp)



	//sp2 := SuperMan{Human{"张三","李四"},Level: 12}
	sp.Walk()

}