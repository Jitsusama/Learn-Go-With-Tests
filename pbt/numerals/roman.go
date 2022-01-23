package numerals

import "strings"

type converter struct {
	Value  int
	Symbol string
}

var conversions = []converter{
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func ConvertToRoman(arabic int) string {
	var roman strings.Builder
	for _, conversion := range conversions {
		for arabic >= conversion.Value {
			roman.WriteString(conversion.Symbol)
			arabic -= conversion.Value
		}
	}
	return roman.String()
}
