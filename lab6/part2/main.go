package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fibonacci(ch, wg)

	wg.Add(1)
	go readChan(ch, wg)

	wg.Wait()
}

func fibonacci(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i + sum
		ch <- sum
		time.Sleep(500 * time.Millisecond)
	}

	close(ch)
}

func readChan(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Println(val)
		time.Sleep(500 * time.Millisecond)
	}
}
