// 1217. Minimum Cost to Move Chips to The Same Position
//
// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
//
//	Moving 2 steps == 0 cost
//	Moving 1 step == 1 cost
//	+, - direction possible
//	What is the minimum cost?

// minCostToMoveChips function finds the minimum cost to make the chips the same position.
// Main logic is based on the parity of the position due to same parity has no cost.
// Thus, smaller value of the count for the parity will be the answer.
func minCostToMoveChips(position []int) int {
	var cnt int
	for _, number := range position {
		if number%2 == 1 {
			cnt++
		}
	}
	if cnt > len(position)-cnt {
		return len(position) - cnt
	} else {
		return cnt
	}
}

