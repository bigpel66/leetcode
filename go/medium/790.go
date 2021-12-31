// 790. Domino and Tromino Tiling
//
// https://leetcode.com/problems/domino-and-tromino-tiling/
//
// 	Tiling can be done with Domino and Tromino
//	Finding the number of cases to tile with the size n
const mod = 1e9 + 7

// Size Bigger than 3
// Using Domino -> i - 1 case & i - 2 case
// Using Tromino -> (i - 3 ... 0) * 2
// dp[i] = dp[i-1] + dp[i-2] + sigma(0 to i - 3) * 2
// dp[i] = dp[i-1] + dp[i-3] + dp[i-1]
func numTilings(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % mod
	}
	return dp[n]
}
