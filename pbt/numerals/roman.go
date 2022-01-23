package numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var roman strings.Builder
	for arabic > 0 {
		switch {
		case arabic > 4:
			roman.WriteString("V")
			arabic -= 5
		case arabic > 3:
			roman.WriteString("IV")
			arabic -= 4
		default:
			roman.WriteString("I")
			arabic--
		}
	}
	return roman.String()
}
