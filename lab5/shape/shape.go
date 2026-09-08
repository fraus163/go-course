package shape

import "fmt"

type Shape interface {
	Area() float64
}

func CalcAreas(s []Shape) {
	for k, v := range s {
		fmt.Println("Площадь фигуры под индексом", k, "равна", v.Area())
	}
}
