package stringutils

func Reverse(str string) string {
	var bytes []byte
	len := len(str)

	for i := len - 1; i >= 0; i-- {
		bytes = append(bytes, str[i])
	}

	return string(bytes)
}
