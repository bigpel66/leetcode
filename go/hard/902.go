// 902. Numbers At Most N Given Digit Set
//
// https://leetcode.com/problems/numbers-at-most-n-given-digit-set/
//      Digits' maximum length is 9, and consists of 1 to 9 with non-decreasing order.

// strconv package is used for the Itoa.
import (
	"strconv"
)

// pow function calculates the power number based on the input arguments.
func pow(n, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= n
	}
	return result
}

// atMostNGivenDigitSet function calculates the number of set which is less than or equal to n.
// If (the length of n) is x and i is 1 to x - 1, then (the number of digits)^i should be accumulated.
// The important count is occurred when the length of the composed number with digits is equal to (the length of n), so comparing between i th element of n and  each digit could be considered to the accumulated result.
// If the i th element of n is greater than each digit, then (the number of digits)^(the length of n - i - 1) should be accumulated.
// Else if i th element of th n is equal to the specific digit, then i + 1 th element should be compared. Otherwise, every accumulation is calculated. Thus, no more iterating is valid.
func atMostNGivenDigitSet(digits []string, n int) int {
	nStr, acc := strconv.Itoa(n), 0
	nLength, nDigits := len(nStr), len(digits)
	for i := 1; i < nLength; i++ {
		acc += pow(nDigits, i)
	}
	for i := 0; i < nLength; i++ {
		catchSameDigit := false
		for j := 0; j < nDigits; j++ {
			if digits[j][0] < nStr[i] {
				acc += pow(nDigits, nLength-i-1)
			} else if digits[j][0] == nStr[i] {
				catchSameDigit = true
			}
		}
		if !catchSameDigit {
			return acc
		}
	}
	return acc + 1
}
