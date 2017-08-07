package main

import (
	"calcstats"
	"fmt"
)

func main() {
	fmt.Println("Calc Stats")
	numb_seq := []int{6, 9, 15, -2, 92, 11}
	fmt.Printf("%v, %T\n", numb_seq, numb_seq)
	fmt.Println("minimum value = ", calcstats.CalcMinValue(numb_seq))
}
