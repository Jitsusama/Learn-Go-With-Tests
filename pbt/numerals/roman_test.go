package numerals

import (
	"fmt"
	"testing"
)

var conversionTests = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"}, {2, "II"}, {3, "III"}, {4, "IV"},
	{5, "V"}, {6, "VI"}, {7, "VII"}, {8, "VIII"},
	{9, "IX"}, {10, "X"}, {14, "XIV"}, {18, "XVIII"},
	{20, "XX"}, {39, "XXXIX"}, {40, "XL"}, {47, "XLVII"},
	{49, "XLIX"}, {50, "L"}, {90, "LC"}, {100, "C"},
	{400, "CD"}, {500, "D"}, {900, "CM"}, {1000, "M"},
	{1984, "MCMLXXXIV"},
}

func TestArabicToRomanConversions(t *testing.T) {
	for _, c := range conversionTests {
		d := fmt.Sprintf("%04d equals %v", c.Arabic, c.Roman)
		t.Run(d, func(t *testing.T) {
			if got := ConvertToRoman(c.Arabic); got != c.Roman {
				t.Errorf("got %v want %v", got, c.Roman)
			}
		})
	}
}

func TestRomanToArabicConversions(t *testing.T) {
	for _, c := range conversionTests[:2] {
		d := fmt.Sprintf("%v equals %04d", c.Roman, c.Arabic)
		t.Run(d, func(t *testing.T) {
			if got := ConvertToArabic(c.Roman); got != c.Arabic {
				t.Errorf("got %d want %d", got, c.Arabic)
			}
		})
	}
}
