// 8. String to Integer (atoi)
//
// https://leetcode.com/problems/string-to-integer-atoi/

// math package has been imported to use Min and Max value of int32.
// unicode package has been imported to check the specific rune.
import (
	"math"
	"unicode"
)

// myAtoi function converts string to int.
// Some points are different from Atoi in strconv.
// The logic of this function consists of skipping white spaces, checking sign, appending the runes to int.
// If the converted number is overflowed or underflowed more than the 4 bytes number, return Min or Max number of int32.
func myAtoi(s string) int {
	sign := 1
	val := 0
	pos := 0
	for pos < len(s) && unicode.IsSpace(rune(s[pos])) {
		pos++
	}
	if pos == len(s) {
		return 0
	}
	if s[pos] == '-' || s[pos] == '+' {
		sign *= (int(s[pos]) - ',') * -1
		pos++
	}
	for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
		val = val*10 + int(s[pos]-'0')
		if val*sign < math.MinInt32 {
			return math.MinInt32
		} else if val*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		pos++
	}
	return val * sign
}
