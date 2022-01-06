// 131. Palindrome Partitioning
//
// https://leetcode.com/problems/palindrome-partitioning/

// palindrome function checks whether the input string is palindrome or not.
func palindrome(s string) bool {
	size := len(s)
	if size == 0 || size == 1 {
		return true
	}
	for i := 0; i < size; i++ {
		if s[i] != s[size-1] {
			return false
		}
		size--
	}
	return true
}

// partition function returns the all of the subsets that can be partitioned with the given input string s.
// Edge case is the length of string s is 0 or 1. (0 for Recursion case, 1 for Palindrome case)
// 1 / 2 3 4 5 6 ...
// 1 2 / 3 4 5 6 ...
// 1 2 3 / 4 5 6 ...
// 1 2 3 4 / 5 6 ...
// 1 2 3 4 5 / 6 ...
// ways has been used.
// If the left subset is palindrome, then find the subsets of right one must be found by recursion.
// After the recursion, every subset of the right one will be resolved. Then append the left one and every element of the right one should be appended to the answer.
// However, if the left one is not palindrome or i == len(s) need not to iterate to append subsets because there are nothing to append.
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	} else if len(s) == 1 {
		return [][]string{[]string{s}}
	}
	answer := [][]string{}
	for i := 1; i <= len(s); i++ {
		if !palindrome(s[:i]) {
			continue
		}
		if i == len(s) {
			answer = append(answer, []string{s[:i]})
			continue
		}
		temp := partition(s[i:])
		for _, t := range temp {
			set := append([]string{s[:i]}, t...)
			answer = append(answer, set)
		}
	}
	return answer
}
