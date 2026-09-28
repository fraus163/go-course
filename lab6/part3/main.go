package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go random(ch1)
	go isEven(ch1, ch2)

	for {
		select {
		case num := <-ch1:
			fmt.Println("Получено число:", num)
		case isEven := <-ch2:
			if isEven {
				fmt.Println("Число четное")
			} else {
				fmt.Println("Число нечетное")
			}
		}
	}
}

func random(ch1 chan int) {
	for {
		time.Sleep(1 * time.Second)
		num := rand.Int()
		ch1 <- num
	}
}

func isEven(ch1 chan int, ch2 chan bool) {
	for {
		time.Sleep(1 * time.Second)
		num := <-ch1

		if num%2 != 0 {
			ch2 <- false
		} else {
			ch2 <- true
		}
	}
}
