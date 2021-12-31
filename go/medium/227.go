// 227. Basic Calculator II
//
// https://leetcode.com/problems/basic-calculator-ii/

// strings package is used for replacing spaces to nothing.
// unicode package is used for checking the rune is digit or not.
import (
	"strings"
	"unicode"
)

// calculate function calculates the given string expression.
// Only "+", "-", "*", "/" operators are allowed.
// As many spaces as possible between operands and operators.
// Such a simple calculating does not require the stack operation.
func calculate(s string) int {
	num, temp, res, lastOperator := 0, 0, 0, '+'
	strings.Replace(s, " ", "", -1)
	for i, char := range s {
		if unicode.IsDigit(char) {
			num = num*10 + (int(char - '0'))
		}
		if i == len(s)-1 || char == '+' || char == '-' || char == '*' || char == '/' {
			switch lastOperator {
			case '+':
				res += temp
				temp = num
			case '-':
				res += temp
				temp = -num
			case '*':
				temp *= num
			case '/':
				temp /= num
			}
			num = 0
			lastOperator = char
		}
	}
	res += temp
	return res
}
