package numerals

import "testing"

func TestRomainNumerals(t *testing.T) {
	got := ConvertToRoman(1)

	if got != "I" {
		t.Errorf("got %q want %q", got, "I")
	}
}
