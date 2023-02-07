package main

import "fmt"

func swap(a, b int) {
	temp := a
	a = b
	b = temp
}

func swap2(a,b *int){
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a,b  := 10, 20

	swap(a, b)

	fmt.Println("swap1")
	fmt.Println("a = ", a, "b = ", b)

	fmt.Println("swap2")
	swap2(&a, &b)
	fmt.Println("a = ", a, "b = ", b)


	var p *int
	p = &a
	fmt.Println(*p)

	q := &p
	fmt.Println(**q)
}
