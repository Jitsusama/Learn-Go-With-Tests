package numerals

import "strings"

type converter struct {
	Arabic int
	Roman  string
}

var conversions = []converter{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "LC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var roman strings.Builder
	for _, conversion := range conversions {
		for arabic >= conversion.Arabic {
			roman.WriteString(conversion.Roman)
			arabic -= conversion.Arabic
		}
	}
	return roman.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for range roman {
		total++
	}
	return total
}
