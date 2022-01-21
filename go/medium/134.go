// 134. Gas Station
//
// https://leetcode.com/problems/gas-station/

// canCompleteCircuit function finds the index to travers all stations successfully.
// gas has the ammount of gas to be gained.
// cost has the amount of gas to be spent for i to i + 1 station.
// Find all of the diffs to traverse (gas - cost) and total sum of the diffs.
// If the sum of diffs is negative, it means traversing all nodes is impossible.
// Else it is absolutely possible to traverse all nodes.
// Thus, just iterate all of the nodes and check the accumulation of diffs from index 0.
// Check whether the accumulated value is negative or not.
// If the accumulated value is negative, it means that point cannot be the starting point.
// In that case, update the starting point and accumulated value. Because the index cannot be the starting point. Starting point should be current index + 1 and accumulated value should be 0 value.
// This iterating process means it is possible to traverse all nodes doubtlessly, so the answer of this case is the updated starting point index.
func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	diff := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		diff[i] = gas[i] - cost[i]
		sum += diff[i]
	}
	if sum < 0 {
		return -1
	}
	startingPoint := 0
	leftover := 0
	for i := 0; i < len(gas); i++ {
		leftover += diff[i]
		if leftover < 0 {
			leftover = 0
			startingPoint = i + 1
		}
	}
	return startingPoint
}
