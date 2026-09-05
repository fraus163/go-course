package main

import "fmt"

type Rectangle struct {
	a int
	b int
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	str := "String"

	len := strLength(str)
	fmt.Println("Длина строки", str, ":", len, "байт")

	rectangle := Rectangle{a: 5, b: 10}

	area := rectangle.calcArea()
	fmt.Println("Площадь прямоугольника:", area)

	a := 45
	b := 22

	calcAvg(a, b)
}

func strLength(str string) int {
	return len(str)
}

func (r Rectangle) calcArea() int {
	return r.a * r.b
}

func calcAvg(a int, b int) {
	fmt.Println("(a + b) / 2 =", (a+b)/2)
}
