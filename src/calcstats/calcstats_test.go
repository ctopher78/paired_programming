package calcstats

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestFindMinValue(t *testing.T) {
	seq := []int{6, 9, 15, -2, 92, 11}
	expected := -2
	actual := FindMinValue(seq)

	if actual != expected {
		t.Errorf("calc(%v): expected %v, actual %v", seq, expected, actual)
	}
}

func TestFindMaxValue(t *testing.T) {
	seq := []int{6, 9, 15, -2, 92, 11}
	expected := 92
	actual := FindMaxValue(seq)

	if actual != expected {
		t.Errorf("calc(%v): expected %v, actual %v", seq, expected, actual)
	}
}

func TestCountNumberOfElements(t *testing.T) {
	seq := []int{6, 9, 15, -2, 92, 11}
	expected := 6
	actual := CountNumberOfElements(seq)

	if actual != expected {
		t.Errorf("calc(%v): expected %v, actual %v", seq, expected, actual)
	}
}

func TestCalcSumOfElements(t *testing.T) {
	seq := []int{6, 9, 1}
	expected := 16
	actual := CalcSumOfElements(seq)

	if actual != expected {
		t.Errorf("calc(%v): expected %v, actual %v", seq, expected, actual)
	}
}

func TestCalcAverageValue(t *testing.T) {
	seq := []int{6, 9, 1}
	expected := 5
	actual := CalcAverageValue(seq)

	if actual != expected {
		t.Errorf("calc(%v): expected %v, actual %v", seq, expected, actual)
	}
}

var calcStats = []struct {
	num_seq  []int
	expected *CalcStats
}{
	{[]int{6, 9, 15, -2, 92, 11},
		&CalcStats{
			minVal:  -2,
			maxVal:  92,
			numElem: 6,
			sumElem: 131,
			aveVal:  21,
		},
	},
	{[]int{4, 144, 11, 0, 0, 7},
		&CalcStats{
			minVal:  0,
			maxVal:  144,
			numElem: 6,
			sumElem: 166,
			aveVal:  27,
		},
	},
}

func TestNewCalcStats(t *testing.T) {
	for _, tt := range calcStats {
		actual := NewCalcStats(tt.num_seq)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("calc(%v): expected %v, actual %v", tt.num_seq, tt.expected, actual)
		}
	}
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	os.Stdout = old

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestPrintStats(t *testing.T) {
	seq := []int{6, 9, 1}
	expected := `
Minimum Value = 1
Maximum Value = 9
Number of elements in the sequence = 3
Average value = 5
`
	cs := NewCalcStats(seq)

	actual := captureStdout(cs.PrintStats)
	if expected != actual {
		t.Errorf("\nexpected %v\n, actual %v\n", expected, actual)
	}
}
