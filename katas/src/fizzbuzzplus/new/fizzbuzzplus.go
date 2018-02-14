package fizzbuzz

// Write a program that prints the numbers from 1 to 100, but...

// numbers that are exact multiples of 3,
// or that contain 3, must print a string containing "Fizz"
//   For example 9 -> "...Fizz..."
//   For example 31 -> "...Fizz..."

// numbers that are exact multiples of 5,
// or that contain 5, must print a string containing "Buzz"
//   For example 10 -> "...Buzz..."
//   For example 51 -> "...Buzz..."

import (
	"strconv"
)

type Fizzbuzzer interface {
	Fizzbuzz()
}

type Fizint int

func(fi Fizint) Fizzbuzz() string {
	fb := ""
	if fi%3 == 0 {
		fb += "fizz"
	}
	if fi%5 == 0 {
		fb += "buzz"
	}
	if fb == "" {
		return strconv.Itoa(int(fi))
	}
	return fb
}

func GetFunc(f func(int) string) func(int) string {
	return f
}

func Fizzbuzz(i int) string {
	fb := ""
	if i%3 == 0 {
		fb += "fizz"
	}
	if i%5 == 0 {
		fb += "buzz"
	}
	if fb == "" {
		return strconv.Itoa(i)
	}
	return fb
}
