package main

import "fmt"

func main() {
	a := 4
	b := 2

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)

	x := 20.5
	y := 17.1

	sumAndSub(x, y)
}

func sumAndSub(a float64, b float64) {
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
}
