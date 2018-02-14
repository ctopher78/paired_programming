package main

import (
	"fmt"
	"strconv"
	fb "github.com/ccrook/paired_programming/katas/src/fizzbuzzplus/new"
)

//var i flag.StringVar
func main() {
	f := fb.GetFunc(fb.Fizzbuzz)
	fmt.Println(f(3))

	f = fb.GetFunc(Buzzfizz)
	fmt.Println(f(3))

	var fizzint fb.Fizint = 3
	
	fmt.Println(fizzint.Fizzbuzz())
	greet.Hello()
}

func Buzzfizz(i int) string {
	fb := ""
	if i%3 == 0 {
		fb += "buzz"
	}
	if i%5 == 0 {
		fb += "fizz"
	}
	if fb == "" {
		return strconv.Itoa(i)
	}
	return fb
}