package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число:")
	fmt.Scan(&b)

	sum := a + b

	fmt.Println("a + b =", sum)
}
