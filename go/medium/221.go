// 221. Maximal Square
//
// https://leetcode.com/problems/maximal-square/description/

// min function get the value which is smaller than another.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max function get the value which is bigger than another.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maximalSquare function finds the largest square from the given matrix
// Look up and update the side length on the tile.
// Position with (i, j) will be updated by (i - 1, j), (i, j - 1), (i - 1, j - 1).
func maximalSquare(matrix [][]byte) int {
	ans := 0
	row, col := len(matrix), len(matrix[0])
	tile := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		tile[i] = make([]int, col+1)
	}
	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			tile[i][j] = int(matrix[i-1][j-1] - '0')
			if tile[i][j] == 1 {
				tile[i][j] = min(tile[i-1][j-1], min(tile[i-1][j], tile[i][j-1])) + 1
			}
			ans = max(ans, tile[i][j])
		}
	}
	return ans * ans
}
