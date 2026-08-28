package sprint


func BulkAtoi(arr []string) []int {

	var result []int
	var integer int

	for s := range arr {
		integer = StrToInt(arr[s])
		result = append(result, integer)
	}
	return result
}

func StrToInt(s string) int {

	result:=0
	var sign rune
	var value int

	// if string is empty, give me 0
	if len(s) == 0 {
		return 0
	}

	// if the leght of the string is 1 
	if len(s) == 1 {
		// and is not a number, give me 0
		if s[0] < '0' || s[0] > '9' {
			return 0
		// Otherwise give me da number
		} else {
		result = int(rune(s[0] - '0'))
		return result
		}
	}	

	// if the first rune is negative symbol and the second rune is not a number, give me a 0
	if s[0] == '-' && (s[1] < '0' || s[1] > '9') {
		return 0
		// Otherwise, memorize that sign is negative
		} else {
		sign = '-'
		}

	// of each value starting from the second
	for n := 1; n < len(s); n++ {
		// if is not a number, give me 0
		if s[n] < '0' || s[n] > '9' {
			return 0
		}
		// Otherwise give me the number as an integer of the value
		value = int(rune(s[n] - '0'))
		// becomes the temporary result, added to a starting empty sum. 
		// The previous temporary result is multiplied by 10 and we add the last found value.
		result = result*10 + value
		}
		// If the sign was negative, give me back the number as negative
		if sign == '-' {
			result = -result
		}
	return result

}
