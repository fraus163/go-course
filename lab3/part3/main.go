package main

import (
	"fmt"
	"slices"
)

func main() {
	//  Создание массива из 5 целых чисел
	arr := [5]int{
		10,
		2,
		457,
		332,
		61,
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Создание среза из массива
	slice := make([]int, 0)

	// Добавление элементов в слайс
	slice = append(slice, 88)
	slice = append(slice, 75)

	fmt.Println("После добавления:", slice)

	// Удаление последнего элемента
	slice = slices.Delete(slice, 1, 2)
	fmt.Println("После удаления:", slice)
}
