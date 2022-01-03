// 1010. Pairs of Songs With Total Durations Divisible by 60
//
// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/

// numPairsDivisibleBy60 function finds the number of pairs that makes divisible by 60.
// First, check 60 - (one of time t % 60) exists. This element is the count that sum of two elements divisible by 60.
// Above value of the element is 0 by default, so just addition is fine.
// Second, one of time t % 60 should be incremented that a half of a pair is already exists.
// Exceptionally, the number which is divisible by 60 would be index 60 but there's no index 60 on the first process.
// Though index 60 exsists, index 0 will be updated in the second process.
// Thus, specific case should be process when the number is divisible by 60.
// When the number is divisible by 60, count should be updated by the index 0.
func numPairsDivisibleBy60(time []int) int {
	count := 0
	pair := make([]int, 60)
	for _, t := range time {
		if val := t % 60; val == 0 {
			count += pair[0]
		} else {
			count += pair[60-val]
		}
		pair[t%60]++
	}
	return count
}
