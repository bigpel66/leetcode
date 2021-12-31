// 1446. Consecutive Charaters
//
// https://leetcode.com/problems/consecutive-characters/

// maxPower function calculates the length of the longest consecutive character on the given string s.
// Just count the characters and update the answer with a single loop.
func maxPower(s string) int {
	temp, length := 0, 0
	previous := s[0]
	for i := 0; i < len(s); i++ {
		if previous == s[i] {
			temp++
		} else {
			temp = 1
		}
		if length < temp {
			length = temp
		}
		previous = s[i]
	}
	return length
}
