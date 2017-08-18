package calcstats

import "testing"

var calc = []struct {
	num_seq  []int
	expected int
}{
	{[]int{6, 9, 15, -2, 92, 11}, -2},
	{[]int{4, 144, 11, 0, 0, 7}, 0},
}

func TestCalcMinValue(t *testing.T) {
	for _, tt := range calc {
		actual := CalcMinValue(tt.num_seq)
		if actual != tt.expected {
			t.Errorf("calc(%v): expected %v, actual %v", tt.num_seq, tt.expected, actual)
		}
	}
}

func TestCalcMaxValue(t *testing.T) {
	num_seq := []int{6, 9, 15, -2, 92, 11}
	expected := 92
	actual := CalcMaxValue(num_seq)
	if actual != expected {
		t.Errorf("calc(%v): expected %v, actual %v", num_seq, expected, actual)
	}
}

func TestNumOfElements(t *testing.T) {
	num_seq := []int{6, 9, 15, -2, 92, 11}
	num := NumOfElements(num_seq)
	if num != 6 {
		t.Errorf("count(%v): expected %v, actual %v", num_seq, 6, num)
	}
}

func TestAverageValue(t *testing.T) {
	num_seq := []int{10, 8}
	expected := 9
	actual := AverageValue(num_seq)

	if actual != expected {
		t.Errorf("count(%v): expected %v, actual %v", num_seq, expected, actual)
	}
}
