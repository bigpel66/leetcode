// 421. Maximum XOR of Two Numbers in an Array
//
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/

// findMaximumXOR function finds the maximum value of operation xor-ed by the elements in the given array.
// By thinking a greedy way of Bit Manipulation, if xor-ed result of the higher digit is marked as 1, then it would be the candiate of the answer value.
// As an example, 100 is bigger than any of 2 digits numbers of xor-ed result.
// The next one to solve the problem as a greedy way is that reflect the higher bit result to the lower bit.
// If the result of xor with several numbers has been made the 100 as a result, then the answer could be made by the two numbers which yield the xor-ed value 100.
// To reflect and ensure the greedy answer, if the xor-ed bit can be yield, memoize the value and use it on the next operations.
// This can be done by the hash map instance which type is map[int]bool and bit shifting.
// The expression of the hash map memoization is that if a^b=c, then a^c=b can be done.
func findMaximumXOR(nums []int) int {
	xor := 0
	for i := 31; i >= 0; i-- {
		hash := make(map[int]bool)
		for _, n := range nums {
			hash[n>>i<<i] = true
		}
		trial := xor | 1<<i
		for key := range hash {
			if hash[trial^key] {
				xor = trial
				break
			}
		}
	}
	return xor
}
