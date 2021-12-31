// 394. Decode String
//
// https://leetcode.com/problems/decode-string/

// strings package is imported to use the Builder
import (
	"strings"
)

// The map named digit for byte and int has been used for decoding the numeric value.
var digit = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

// decodeString function decode the encoded string.
// Encoded string is formed with number + [string].
// This function returns the decoded string.
// The form can be embedded such as "3[a2[c]]", so recursion solution is necessary.
// To perform the logic effective way, strings.Builder and byte slices are used.
func decodeString(s string) string {
	bytes := []byte(s)
	var logic func() string
	logic = func() string {
		sb := strings.Builder{}
		nb := 0
		for len(bytes) > 0 {
			c := bytes[0]
			bytes = bytes[1:]
			if d, ok := digit[c]; ok {
				nb = nb*10 + d
			} else if c == '[' {
				sb.WriteString(strings.Repeat(logic(), nb))
				nb = 0
			} else if c == ']' {
				break
			} else {
				sb.WriteByte(c)
			}
		}
		return sb.String()
	}
	return logic()
}
