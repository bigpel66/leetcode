// 1672. Richest Customer Wealth
//
// https://leetcode.com/problems/richest-customer-wealth/

// maximumWealth function returns the maximum value of the person's wealth.
func maximumWealth(accounts [][]int) int {
	answer := 0
	for _, person := range accounts {
		acc := 0
		for _, account := range person {
			acc += account
		}
		if answer < acc {
			answer = acc
		}
	}
	return answer
}
