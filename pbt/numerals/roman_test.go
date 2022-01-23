package numerals

import (
	"fmt"
	"testing"
)

func TestArabicToRomainNumeralConversions(t *testing.T) {
	conversions := []struct {
		Arabic int
		Roman  string
	}{
		{1, "I"}, {2, "II"}, {3, "III"}, {4, "IV"},
		{5, "V"}, {6, "VI"}, {7, "VII"}, {8, "VIII"},
		{9, "IX"}, {10, "X"}, {14, "XIV"}, {18, "XVIII"},
		{20, "XX"}, {39, "XXXIX"}, {40, "XL"}, {47, "XLVII"},
		{49, "XLIX"}, {50, "L"},
	}

	for _, c := range conversions {
		d := fmt.Sprintf("%02d equals %v", c.Arabic, c.Roman)
		t.Run(d, func(t *testing.T) {
			if got := ConvertToRoman(c.Arabic); got != c.Roman {
				t.Errorf("got %q want %q", got, c.Roman)
			}
		})
	}
}
