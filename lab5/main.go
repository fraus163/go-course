package main

import (
	"fmt"
	"lab5/person"
	"lab5/shape"
	"lab5/stringer"
)

func main() {
	// Тест Person
	p := person.Person{
		Name: "name",
		Age:  15,
	}

	fmt.Println("Информация до увеличения возраста:")
	p.GetInfo()
	p.Birthday()
	fmt.Println("Информация после увеличения возраста:")
	p.GetInfo()

	// Тест Shape
	c := shape.Circle{
		Radius: 10,
	}
	r := shape.Rectangle{
		A: 5,
		B: 10,
	}
	s := []shape.Shape{c, r}

	shape.CalcAreas(s)

	// Тест Stringer
	var b stringer.Stringer
	b = stringer.Book{
		Name:   "Metro 2033",
		Author: "D.Glukhovsky",
	}
	b.GetInfo()
}
