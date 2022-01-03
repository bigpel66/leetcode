// 997. Fidn the Town Judge
//
// https://leetcode.com/problems/find-the-town-judge/

// There is a sole 'town judge' in the town. findJudge function finds the judge.
// The judge does not trust anyone, and others trust the judge.
// By using two maps, record the person who is trusting other and the person who is trusted.
// The answer is that the person who is trusted n - 1 (except the self) and does not trust anyone.
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	person := make([]bool, n+1)
	trusted := make(map[int]int)
	for _, r := range trust {
		person[r[0]] = true
		trusted[r[1]]++
	}
	for i, bit := range person {
		if bit == false && trusted[i] == n-1 {
			return i
		}
	}
	return -1
}
