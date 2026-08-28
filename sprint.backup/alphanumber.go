package sprint

// If the number is negative, preserve the minus sign. 
// Replace the digits with lowercase letters, where:
// 0 becomes a, 1 becomes b, and so on up to 9, which becomes j.

func AlphaNumber(n int) string {

	// We make sure that if the integer is 0 we have the correspondent string as "a"
	if n == 0 {
		return "a"
	}

	result := ""
	sign := ""

	// If n is negative we create a variable and assign the string "-" to it, then we convert n to a positive number
	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string('a'+rune(digit)) + result
		n = n / 10
	}
	return sign + result
}