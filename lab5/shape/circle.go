package shape

import "math"

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pow(float64(c.Radius), 2) * 3.14
}
