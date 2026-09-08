package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scan(&str)

	strInUppercase := strings.ToUpper(str)

	fmt.Println("Строка в верхнем регистре:", strInUppercase)
}
