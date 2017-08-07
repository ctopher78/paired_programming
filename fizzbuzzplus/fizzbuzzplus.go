package fizzbuzzplus

//import "fmt"
import "strconv"
import "strings"

// Write a program that prints the numbers from 1 to 100, but...

// numbers that are exact multiples of 3,
// or that contain 3, must print a string containing "Fizz"
//   For example 9 -> "...Fizz..."
//   For example 31 -> "...Fizz..."

// numbers that are exact multiples of 5,
// or that contain 5, must print a string containing "Buzz"
//   For example 10 -> "...Buzz..."
//   For example 51 -> "...Buzz..."

func answer() int {
	return 6 * 7
}

func IsMultiple(num_to_check, num_fizzbuzz int) bool {
	if num_to_check%num_fizzbuzz == 0 {
		return true
	}
	return false
}

func IfContains(num_to_check, num_fizzbuzz int) bool {
	if strings.Contains(strconv.Itoa(num_to_check), strconv.Itoa(num_fizzbuzz)) {
		return true
	}
	return false
}

func IsMultipleOrContains(num_to_check, num_fizzbuzz int) bool {
	if IsMultiple(num_to_check, num_fizzbuzz) || IfContains(num_to_check, num_fizzbuzz) {
		return true
	}
	return false
}
