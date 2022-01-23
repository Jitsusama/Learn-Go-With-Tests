package numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var roman strings.Builder
	for i := 0; i < arabic; i++ {
		roman.WriteString("I")
	}
	return roman.String()
}
