// 67. Add Binary
//
// https://leetcode.com/problems/add-binary/

// addBinary function adds the a, b given numbers by string.
// The length of the a or b is quite big (onto the 10^4), so byte slice is required.
// Digit and Carry are calculated by the bit manipulation because only 0 or 1 appeared.
func addBinary(a string, b string) string {
	size := len(a)
	if size < len(b) {
		size = len(b)
	}
	number := make([]byte, size+1)
	var carry byte
	for i, j, k := len(a)-1, len(b)-1, size; i >= 0 || j >= 0; {
		var digitA, digitB, digit byte
		if i >= 0 {
			digitA = a[i] - '0'
		}
		if j >= 0 {
			digitB = b[j] - '0'
		}
		digit = digitA ^ digitB ^ carry
		carry = digitA&digitB | ((digitA | digitB) & carry)
		number[k] = digit + '0'
		i, j, k = i-1, j-1, k-1
	}
	if carry == 1 {
		number[0] = '1'
		return string(number)
	}
	return string(number[1:])
}
