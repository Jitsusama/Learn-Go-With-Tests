package numerals

import "testing"

func TestRomainNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 = I", 1, "I"}, {"2 = II", 2, "II"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			if got := ConvertToRoman(test.Arabic); got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}
