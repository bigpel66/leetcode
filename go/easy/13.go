// 13. Roman to Integer
//
// https://leetcode.com/problems/roman-to-integer/

// romanToInt function converts the roman number into the integer.
// Only valid roman number will be given. (IV is possible but IIV impossible.)
// To convert correctly, read from the last digit is required.
func romanToInt(s string) int {
	rtoi := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	val := rtoi[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if rtoi[s[i+1]] > rtoi[s[i]] {
			val -= rtoi[s[i]]
		} else {
			val += rtoi[s[i]]
		}
	}
	return val
}
