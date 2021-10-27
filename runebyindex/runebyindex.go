package runebyindex

import "fmt"

func RuneByIndex(str *string, i *int) (r rune, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("RuneByIndex: %s", rec)
		}
	}()

	arrRune := []rune(*str)

	return arrRune[*i], nil
}
