package main

import (
	"fmt"
	"github.com/ccrook/paired_programming/katas/src/fizzbuzzplus/orig"
)

func main() {
	for i := 0; i <= 100; i++ {
		if fizzbuzzplus.IsMultipleOrContains(i, 5) && fizzbuzzplus.IsMultipleOrContains(i, 3) {
			fmt.Println("FizzBuzz")
		} else if fizzbuzzplus.IsMultipleOrContains(i, 5) {
			fmt.Println("Buzz")
		} else if fizzbuzzplus.IsMultipleOrContains(i, 3) {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}
