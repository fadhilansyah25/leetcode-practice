func romanToInt(s string) int {
    	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s = strings.Replace(s, "CM", "CCCCCCCCC", -1) // 900
	s = strings.Replace(s, "CD", "CCCC", -1)      // 400
	s = strings.Replace(s, "XC", "XXXXXXXXX", -1) // 90
	s = strings.Replace(s, "XL", "XXXX", -1)      // 40
	s = strings.Replace(s, "IX", "IIIIIIIII", -1) // 9
	s = strings.Replace(s, "IV", "IIII", -1)      // 4

	var sum int
	for _, roman := range s {
		sum += romans[roman]
	}
	return sum
}