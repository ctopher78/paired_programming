package calcstats

import "sort"

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

func CalcMinValue(num_seq []int) int {
	sort.Ints(num_seq)
	return num_seq[0]
}

func CalcMaxValue(num_seq []int) int {
	sort.Ints(num_seq)
	return num_seq[len(num_seq)-1]
}

func NumOfElements(seq []int) int {
	return len(seq)
}

func AverageValue(seq []int) int {
	total := 0
	for _, num := range seq {
		total += num
	}

	avg := total / len(seq)

	return avg
}
