package reverse

import "strings"

func Reverse(str string) string {
	var b strings.Builder

	sliceRune := []rune(str)
	for i := len(sliceRune) - 1; i >= 0; i-- {
		b.WriteRune(sliceRune[i])
	}

	return b.String()
}
