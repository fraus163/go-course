package main

import "fmt"

func main() {
	slice := []string{
		"computer",
		"cpu",
		"videocard",
		"ram",
	}

	var maxLen int

	for _, v := range slice {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	fmt.Println("Максимальная длина строки:", maxLen)
}
