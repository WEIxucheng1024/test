package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriterBook()
}

type Book struct {
	
}

func (this *Book) ReadBook()  {
	fmt.Println("read book")
}

func (this *Book) WriterBook()  {
	fmt.Printf("writer book")
}

func main() {
	b := &Book{}
	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriterBook()
}
