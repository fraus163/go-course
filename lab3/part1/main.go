package main

import (
	"fmt"
	"part1/mathutils"
)

func main() {
	var num int

	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Число не должно быть отрицательным")
		return
	}

	result := mathutils.Factorial(num)
	fmt.Println("Результат:", result)
}
