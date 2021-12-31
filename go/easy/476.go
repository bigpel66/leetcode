// 476. Number Complement
//
// https://leetcode.com/problems/number-complement/

// findComplement function finds the complement number from input number.
// Binary of 5 is 101, and complement of 101 is 010 which is 2.
// To find the complement number, must consider the type size.
// 5 in 64-bit is 0000....101, and its complement is 1111....010 which is not 2.
// The complement number 1111....010 should be masked by 0000...111.
func findComplement(num int) int {
	mask := ^int(0)
	for mask&num > 0 {
		mask <<= 1
	}
	return ^num ^ mask
}
