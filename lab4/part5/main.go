package main

import "fmt"

func main() {
	arr := []int{4, 8, 12, 45, 77, 40}

	fmt.Println("Массив:", arr)

	reverseArr := reverse(arr)
	fmt.Println("Обратный массив:", reverseArr)
}

func reverse(arr []int) []int {
	len := len(arr)

	for i := 0; i < len/2; i++ {
		temp := arr[i]
		arr[i] = arr[len-1-i]
		arr[len-1-i] = temp
	}

	return arr
}
