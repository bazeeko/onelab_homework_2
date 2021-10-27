package fibonacci

func Fibonacci() func() int {
	current, next := 0, 1

	return func() (r int) {
		defer func() {
			current, next = next, current+next
		}()
		return current
	}
}
