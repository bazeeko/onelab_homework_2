package itoa

import (
	"strings"

	"github.com/bazeeko/onelab_homework_2/reverse"
)

func Itoa(num int) string {
	if num == 0 {
		return "0"
	}

	var isNegative bool = false
	if num < 0 {
		isNegative = true
		num = -num
	}

	var b strings.Builder
	for num > 0 {
		b.WriteRune('0' + rune(num%10))
		num /= 10
	}

	if isNegative {
		b.WriteRune('-')
	}

	return reverse.Reverse(b.String())
}
