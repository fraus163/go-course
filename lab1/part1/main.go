package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Вывод текущей даты и времени
	now := time.Now()

	fmt.Println("Текущая дата: ", now.Format("2006-01-02 15:04:05"))

	// 2. Создание переменные различных типов (int, float64, string, bool)
	var intVar int
	var floatVar float64
	var strVar string = "String"
	var boolVar bool

	fmt.Println(
		"Int variable: ", intVar,
		", float64 variable: ", floatVar,
		", string variable: ", strVar,
		", bool variable: ", boolVar)

	// 3. Использование краткой формы объявления переменных для создания и вывода переменных
	intNum := 1
	floatNum := 1.0
	str := "String"
	isBool := true

	fmt.Println(
		"intNum: ", intNum,
		", floatNum: ", floatNum,
		", str: ", str,
		", isBool: ", isBool)
}
