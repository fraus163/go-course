package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

	response := checkNum(num)

	fmt.Println(response)
}

func checkNum(num int) string {
	if num == 0 {
		return "Zero"
	} else if num > 0 {
		return "Positive"
	} else {
		return "Negative"
	}
}
