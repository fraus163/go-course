package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	//var mtx sync.Mutex
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			//mtx.Lock()
			time.Sleep(1 * time.Second)
			count++
			fmt.Println(count)
			//mtx.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			//mtx.Lock()
			count++
			time.Sleep(1 * time.Second)
			fmt.Println(count)
			//mtx.Unlock()
		}
	}()

	wg.Wait()
}
