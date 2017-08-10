package main

import (
	"calcstats"
	"fmt"
)

func main() {
	fmt.Printf("** Calc Stats **\n")
	num_seq := []int{6, 9, 15, -2, 92, 11}
	fmt.Printf("Incoming Seq: %v\n", num_seq)

	cs := calcstats.NewCalcStats(num_seq)
	cs.PrintStats()
}
