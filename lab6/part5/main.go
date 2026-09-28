package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Response struct {
	Value float64
	Err   error
}

type Request struct {
	ID         int
	A          float64
	B          float64
	Op         string
	ResultChan chan Response
}

func executeOperation(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s", op)
	}
}

func worker(jobs <-chan Request, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range jobs {
		time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)

		res, err := executeOperation(req.A, req.B, req.Op)

		req.ResultChan <- Response{
			Value: res,
			Err:   err,
		}
	}
}

func client(id int, a, b float64, op string, jobs chan<- Request, wg *sync.WaitGroup) {
	defer wg.Done()

	resChan := make(chan Response, 1)

	req := Request{
		ID:         id,
		A:          a,
		B:          b,
		Op:         op,
		ResultChan: resChan,
	}

	jobs <- req

	resp := <-resChan

	if resp.Err != nil {
		fmt.Printf("[Клиент #%02d] Ошибка при расчете %.2f %s %.2f: %v\n", req.ID, req.A, req.Op, req.B, resp.Err)
	} else {
		fmt.Printf("[Клиент #%02d] Успех: %.2f %s %.2f = %.2f\n", req.ID, req.A, req.Op, req.B, resp.Value)
	}
}

func main() {
	const (
		numWorkers  = 3
		numRequests = 10
	)

	jobs := make(chan Request, numRequests)

	var workerWg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		workerWg.Add(1)
		go worker(jobs, &workerWg)
	}

	var clientWg sync.WaitGroup

	operations := []string{"+", "-", "*", "/", "%"}

	fmt.Println("=== Запуск клиентов и отправка запросов ===")

	for i := 1; i <= numRequests; i++ {
		clientWg.Add(1)

		a := float64(rand.Intn(50) + 1)
		b := float64(rand.Intn(10))
		op := operations[rand.Intn(len(operations))]

		go client(i, a, b, op, jobs, &clientWg)
	}

	clientWg.Wait()

	close(jobs)
	workerWg.Wait()

	fmt.Println("=== Все вычисления завершены, сервер остановлен ===")
}
