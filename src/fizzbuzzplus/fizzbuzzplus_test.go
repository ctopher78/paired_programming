package fizzbuzzplus

import (
	"testing"
)

func Test_life_the_universe_and_everything(t *testing.T) {
	if answer() != 42 {
		t.Error("answer() != 42 as expected.")
	}
}

var fizzbuzz = []struct {
	n          int
	n_fizzbuzz int
	expected   bool
}{
	{1, 3, false},
	{3, 3, true},
	{15, 5, true},
	{77, 5, false},
	{31, 3, true},
}

func Test_isMultipleOrContains(t *testing.T) {
	for _, tt := range fizzbuzz {
		actual := IsMultipleOrContains(tt.n, tt.n_fizzbuzz)
		if actual != tt.expected {
			t.Errorf("fizzbuzz(%v %v): expected %v, actual %v", tt.n, tt.n_fizzbuzz, tt.expected, actual)
		}
	}
}
