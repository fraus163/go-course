package stringer

import "fmt"

type Book struct {
	Name   string
	Author string
}

func (b Book) GetInfo() {
	fmt.Println("Название книги:", b.Name)
	fmt.Println("Автор книги:", b.Author)
}
