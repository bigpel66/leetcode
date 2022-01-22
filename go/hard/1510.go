// 1510. Stone Game IV
//
// https://leetcode.com/problems/stone-game-iv/

// winnerSquareGame function returns true if Alice wins, false if Bob wins.
// About turn to play the game, Alice always starts first.
// Once Alice wins, it should be makred false to true.
// For example to make number 7, Alice can choose number 1 or 4.
// Then, the result depends on dp[7 - 1] or dp[7 - 4].
// If dp[7 - 1] was false (Alice loses), then dp[7] can be marked as true.
// Also if dp[7 - 4] was false (Alice loses), then dp[7] can be marked as true.
// To get dp[7], dp[7 - 1] and dp[7 - 4] can be done as same way. Thus, every number until input n should be calculated.
// Iteratings on outer loop for 1 to n, and inner loop for square numbers until outer one should be done.
func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if !dp[i] {
				dp[i] = !dp[i-j*j]
			}
		}
	}
	return dp[n]
}
