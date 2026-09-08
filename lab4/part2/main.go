package main

import "fmt"

func main() {
	users := map[int]string{
		34:    "Name1",
		592:   "Name2",
		94724: "Name3",
	}

	fmt.Println("Карта до удаления:", users)

	delete(users, 34)

	fmt.Println("Карата после удаления:", users)
}
