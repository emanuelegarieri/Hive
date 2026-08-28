package sprint

func BetweenLimits(from, to rune) string {
	if from > to {
		from, to = to, from
	}

	digits := string("")

	for r := from + 1; r < to; r++ {
		digits += string(r)
	}
	return digits
}