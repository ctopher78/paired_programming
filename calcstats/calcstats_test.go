package calcstats

import "testing"

var calc = []struct {
	num_seq  []int
	expected int
}{
	{[]int{6, 9, 15, -2, 92, 11}, -2},
	{[]int{4, 144, 11, 0, 0, 7}, 0},
}

func Test_CalcMinValue(t *testing.T) {
	for _, tt := range calc {
		actual := CalcMinValue(tt.num_seq)
		if actual != tt.expected {
			t.Errorf("calc(%v): expected %v, actual %v", tt.num_seq, tt.expected, actual)
		}
	}
}
