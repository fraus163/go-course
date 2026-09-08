package main

import "fmt"

func main() {
	users := map[string]int{
		"User1": 18,
		"User2": 25,
		"User3": 38,
		"User4": 29,
	}

	users["User5"] = 41

	fmt.Println(users)

	avgAge := calcAvgAge(users)

	fmt.Println("Средний возраст:", avgAge)
}

func calcAvgAge(users map[string]int) float64 {
	var sumOfAges int
	var numOfPeaple int

	for _, v := range users {
		sumOfAges += v
		numOfPeaple++
	}

	avgAge := float64(sumOfAges) / float64(numOfPeaple)
	return avgAge
}
