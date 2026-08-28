package sprint


func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}

	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}