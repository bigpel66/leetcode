// 1009. Complement of Base 10 Integer
//
// https://leetcode.com/problems/complement-of-base-10-integer/

// bits package has been used to call Len function.
import (
	"math/bits"
)

// bitwiseComplement function finds the number which is the complement of given input.
// Goal is to make the number 0000...101 to 0000...010.
// Find the length of the n's bit and shift the all bits 1 with the length.
// Then 1111...111 will be 1111...000 which is the mask.
// The answer can be made with that mask by xor operation.
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	return (1<<bits.Len(uint(n)) - 1) ^ n
}
