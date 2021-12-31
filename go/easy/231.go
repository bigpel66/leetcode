// 231. Power of Two
//
// https://leetcode.com/problems/power-of-two/

// isPowerOfTwo function returns the truth whether the given n is the number such that the power of two.
// Whether n is the power of two or not can be determined by only 64 iterating with Bit-Manipulation. (int type is 64-bit in Golang.)
// If n is the power of two, then the count of Bit-Manipulation with logical-and 1 should be 1.
func isPowerOfTwo(n int) bool {
	cnt := 0
	for i := 0; i < 64; i++ {
		if n&1 == 1 {
			cnt++
		}
		n = n >> 1
	}
	return cnt == 1
}
