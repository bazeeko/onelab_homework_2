package atoi

import (
	"errors"
	"fmt"
)

var ErrStringIsEmpty = errors.New("string passed as an argument is empty")
var ErrStringNonDigit = errors.New("string contains non-digital elements")

func Atoi(str string) (int, error) {
	if str == "" {
		return 0, fmt.Errorf("Atoi: %w", ErrStringIsEmpty)
	}

	var isNegative bool

	if str[0:1] == "-" {
		isNegative = true
		str = str[1:]
	}

	var result int

	for _, r := range str {
		if r < '0' || r > '9' {
			return 0, fmt.Errorf("Atoi: %w", ErrStringNonDigit)
		}
		result *= 10
		result += int(r) - '0'
	}

	if isNegative {
		result = -result
	}

	return result, nil
}
