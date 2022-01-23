package numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var roman strings.Builder
	for i := arabic; i > 0; i-- {
		if i == 4 {
			roman.WriteString("IV")
			break
		}
		roman.WriteString("I")
	}
	return roman.String()
}
