package mathutils

func Factorial(num int) int {
	if num == 0 {
		return 1
	}

	if num > 1 {
		return Factorial(num-1) * num
	}
	return 1
}
