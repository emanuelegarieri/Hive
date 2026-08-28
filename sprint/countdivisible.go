package sprint

// Write a Go function that takes four integers as input: from, to, step, and divisor. 
// The function should check for exceptions and return 0 if any of the following conditions are met:

// - step is negative or zero.
// - divisor is zero.

// Otherwise
// It should loop through the range of numbers from from (inclusive) to to (exclusive), 
// checking every step-th element, and return the count of elements among these that are divisible by the divisor.

func CountDivisible(from, to, step, divisor int) int {
	if step <= 0 || divisor == 0 {
		return 0
	}

	count := 0

	for i := from; i < to; i += step {
		if i%divisor == 0 {
			count += 1
		}
	}
	return count
}
