// 1463. Cherry Pickup II
//
// https://leetcode.com/problems/cherry-pickup-ii/
// 	There are two robots which are collecting cherries from the given grid.
// 	A robot is on the (0, 0), and another robot is on the (0, cols - 1).
// 	Robot can move toward the left bottom, toward bottom, right bottom.

// max function finds the maximum value among the inputs.
func max(input ...int) int {
	val := -int(^uint(0)>>1) - 1
	for _, i := range input {
		if val < i {
			val = i
		}
	}
	return val
}

// cherryPickup function calculates the maximum cherries that can be collected.
// Let's define the coordinate of rows is x, the coordinate of cols is y.
// Pretend that the 3 Dimensional Slices DP[a][b][c] is the memoization of robot's move.
// Robot 1's coordinate is (a, b) and Robot 2's coordinate is (a, c).
// Then the DP[a][b][c] is the maximum value of collected cherries on that coordinates.
// All of the value in the DP are -1 when they are initialized.
// The spaces for the y have been created 2 tiles more than orignal because of robots move. Thus, the range of y is 1 to cols. (The range of x are as same as the original, so 0 to rows - 1.)
// The first rows of DP can be calculated with (0, 0) and (0, cols - 1). Thus the iterations for the x starts from 1 to rows - 1.
// Originally, y iterates 0 to cols - 1. However, y will be iterated by 1 to cols because tiles have been added.
// There are two ys for the robot1 and robot2. Thus, DP for the current position is DP[i][y1][y2] and it will be accumulated by previous one.
// Because DP solution is optimal, so the maximum value should be found before accumulated. To do this, there are 9 ways by considering the robots move.
// When it gets the maximum value of previous one, the number of cherry should be appened. But if y1 and y2 are the same position, only a sole value should be accumulated. Or, both value should be.
// The answer can be got on the last rows of DP, thus comparing each element should be done.
func cherryPickup(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([][]int, cols+2)
		for j := 0; j < cols+2; j++ {
			dp[i][j] = make([]int, cols+2)
			for k := 0; k < cols+2; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][1][cols] = grid[0][0] + grid[0][cols-1]
	answer := 0
	for i := 1; i < rows; i++ {
		for r1 := 1; r1 < cols+1; r1++ {
			for r2 := 1; r2 < cols+1; r2++ {
				dp[i][r1][r2] = max(
					dp[i-1][r1-1][r2-1],
					dp[i-1][r1-1][r2],
					dp[i-1][r1-1][r2+1],
					dp[i-1][r1][r2-1],
					dp[i-1][r1][r2],
					dp[i-1][r1][r2+1],
					dp[i-1][r1+1][r2-1],
					dp[i-1][r1+1][r2],
					dp[i-1][r1+1][r2+1],
				)
				if dp[i][r1][r2] == -1 {
					continue
				} else {
					if r1 == r2 {
						dp[i][r1][r2] += grid[i][r1-1]
					} else {
						dp[i][r1][r2] += grid[i][r1-1] + grid[i][r2-1]
					}
					if i == rows-1 {
						answer = max(answer, dp[i][r1][r2])
					}
				}
			}
		}
	}
	return answer
}
