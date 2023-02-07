package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	Color string
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	Color string
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animalIF AnimalIF) {
	animalIF.Sleep()
	animalIF.GetColor()
	animalIF.GetType()
}

func main() {
	//var animal AnimalIF
	//animal = &Cat{"白色"}
	//
	//animal.Sleep()
	//
	//animal = &Dog{"黄色"}
	cat := Cat{"黄色"}
	dog := Dog{"黑色"}
	ShowAnimal(&cat)
	ShowAnimal(&dog)

}
