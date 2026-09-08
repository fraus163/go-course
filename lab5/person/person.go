package person

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) GetInfo() {
	fmt.Println("Имя:", p.Name)
	fmt.Println("Возраст:", p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}
