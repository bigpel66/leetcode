// 878. Nth Magical Number
//
// https://leetcode.com/problems/nth-magical-number/
const mod = 1e9 + 7

// gcd function finds the Greatest Common Divisor.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// lcm function finds the Least Common Multiple.
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// nthMagicalNumber function finds the number which is fit to the argument n.
// n is quite big (1 <= n <= 10^9).
// To get the rank of the number, O(N) or O(logN) would be appropriate.
// Searching numbers are ascending orders, so Binary Search is a good strategy.
// if n = 4, a = 2, b = 3, answer is 6.
//  -> ( rank of a + rank of b ) - ( common number rank ) = n th number
//  -> ( answer / a + answer / b) - ( lcm rank) = n th number
func nthMagicalNumber(n int, a int, b int) int {
	var div int = lcm(a, b)
	var low, high int = 0, int(^uint(0) >> 1)
	for low < high {
		mid := (low + high) / 2
		if mid/a+mid/b-mid/div < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low % mod
}
