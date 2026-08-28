package sprint


func FindDividend(from, to, divisor int) int {
	// Return the first number in that range that is divisible by the divisor.
	for i := from; i < to; i++ {
		if i%divisor == 0 {
			return i
		}
	}
	return -1
}