package shape

type Rectangle struct {
	A int
	B int
}

func (r Rectangle) Area() float64 {
	return float64(r.A) * float64(r.B)
}
