package sprint

func StrLength(s string) []int {

	var len_byte int
	var list_or_runes []rune
	var list_of_int []int

	len_byte = len(string(s))
	for _, r := range s {
		list_or_runes = append(list_or_runes, r)
		}	
	
	list_of_int = append(list_of_int, len(list_or_runes), len_byte)

	return list_of_int

}