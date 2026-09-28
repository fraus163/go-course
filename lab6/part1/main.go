package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go isEven(wg, 5)

	wg.Add(1)
	go randomNums(wg, 1, 10)

	wg.Add(1)
	go reverse(wg, "string")

	wg.Wait()
}

func isEven(wg *sync.WaitGroup, num int) {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func randomNums(wg *sync.WaitGroup, from, to int) {
	defer wg.Done()

	time.Sleep(3 * time.Second)

	randomNum := from + rand.Intn(to)
	fmt.Println("Радномное число:", randomNum)
}

func reverse(wg *sync.WaitGroup, str string) {
	defer wg.Done()

	time.Sleep(5 * time.Second)

	var bytes []byte
	len := len(str)

	for i := len - 1; i >= 0; i-- {
		bytes = append(bytes, str[i])
	}

	result := string(bytes)
	fmt.Println("Перевернутая строка:", result)
}
