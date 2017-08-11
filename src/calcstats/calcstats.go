package calcstats

// Your task is to process a sequence of integer numbers
// to determine the following statistics:

//     o) minimum value
//     o) maximum value
//     o) number of elements in the sequence
//     o) average value

// For example: [6, 9, 15, -2, 92, 11]

//     o) minimum value = -2
//     o) maximum value = 92
//     o) number of elements in the sequence = 6
//     o) average value = 21.833333

import (
	"fmt"
	"sort"
)

type CalcStats struct {
	minVal  int
	maxVal  int
	numElem int
	sumElem int
	aveVal  int
}

func NewCalcStats(seq []int) *CalcStats {
	cs := new(CalcStats)
	cs.minVal = findMinValue(seq)
	cs.maxVal = findMaxValue(seq)
	cs.numElem = countNumberOfElements(seq)
	cs.sumElem = calcSumOfElements(seq)
	cs.aveVal = calcAverageValue(seq)
	return cs
}

func (cs *CalcStats) PrintStats() {
	fmt.Printf(`
Minimum Value = %v
Maximum Value = %v
Number of elements in the sequence = %v
Average value = %v
`, cs.minVal, cs.maxVal, cs.numElem, cs.aveVal)
}

func findMinValue(num_seq []int) int {
	sort.Ints(num_seq)
	return num_seq[0]
}

func findMaxValue(num_seq []int) int {
	sort.Ints(num_seq)
	return num_seq[len(num_seq)-1]
}

func countNumberOfElements(num_seq []int) int {
	return len(num_seq)
}

func calcSumOfElements(num_seq []int) int {
	sum := 0
	for _, v := range num_seq {
		sum += v
	}
	return sum
}

func calcAverageValue(num_seq []int) int {
	return calcSumOfElements(num_seq) / len(num_seq)
}
