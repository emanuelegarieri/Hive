package sprint

func StrConcatWith(strs []string, sep string) string {

	var result string

	for r := range strs {
		if result != "" {
			result = result + sep
		}
		result = result + strs[r]
	}
	return result
}