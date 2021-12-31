// 1015. Smallest Integer Divisible by K
//
// https://leetcode.com/problems/smallest-integer-divisible-by-k/
//	k is a positive number
//	n is the number that is divisible by k
//	n only contains the digit 1
//	answer is the smallest number n's length (if no n, -1 is the answer)

//	smallestRepunitDivByK function finds the number n which satisfies the conditions above.
//	Under 10 for the 0-9 digits, 2 and 5 cannot make the last digit as 1.
//	Thus, answer is -1 if k is divisible by 2 and 5.
func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	length, current := 0, 0
	for {
		length++
		current = (current*10 + 1) % k
		if current == 0 {
			return length
		}
	}
}
