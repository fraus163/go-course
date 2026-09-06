package main

import (
	"fmt"
	"part2/stringutils"
)

func main() {
	var str string

	fmt.Print("Введите строку: ")
	fmt.Scan(&str)

	result := stringutils.Reverse(str)
	fmt.Println("Обратная строка:", result)
}
