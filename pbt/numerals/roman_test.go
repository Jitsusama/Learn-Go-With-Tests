package numerals

import "testing"

func TestRomainNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 = I", 1, "I"}, {"2 = II", 2, "II"},
		{"3 = III", 3, "III"}, {"4 = IV", 4, "IV"},
		{"5 = V", 5, "V"}, {"6 = VI", 6, "VI"},
		{"7 = VII", 7, "VII"}, {"8 = VIII", 8, "VIII"},
		{"9 = IX", 9, "IX"}, {"10 = X", 10, "X"},
		{"14 = XIV", 14, "XIV"}, {"18 = XVIII", 18, "XVIII"},
		{"20 = XX", 20, "XX"}, {"39 = XXXIX", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			if got := ConvertToRoman(test.Arabic); got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}
