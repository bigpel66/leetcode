// 312. Burst Ballons
//
// https://leetcode.com/problems/burst-balloons/

// maxCoins function returns the maximum values when all balloons bursted in some sequences.
// When i th balloon bursted, then values of  i - 1 * i * i + 1 will be producted.
// And then 0 ... i - 1 and i + 1 ... length should be conducted the same logic.
// Picking i th balloon could be divided into the subproblem.
// In the reverse concept, pretend that picking i th ballon is the last balloon to be bursted.
// [ 3 , 1 , 5 , 8 ] ----------> pick 5 as last element to be bursted.
//                               [ 3 , 1 ] [ 8 ] should be calculated.
// [ 3, 1 ]          ----------> same logic on the first case. (pretend x)
// [ 8 ]             ----------> same logic on the first case. (pretend y)
// [ 3 , 1 , 5 , 8 ] ----------> x + y + 1 * 5 * 1 will be the coins that this sequence earned.
// Maybe before the last sequence could be optimal, but not in the last sequence.
// Thus, picking the very first element to be last sequence should be searched too.
// n times loop * n times loop for the left range * n times loop for the right range becomes O(n^3).
// Each element for the last sequence should be maintained the n number of elements, so O(n^2) memory required.
func maxCoins(nums []int) int {
	vals := append([]int{1}, append(nums, 1)...)
	dp := make(map[[2]int]int, 0)
	var rangeSearch func(int, int) int
	rangeSearch = func(left, right int) int {
		if coins, ok := dp[[2]int{left, right}]; ok {
			return coins
		}
		coins := 0
		for i := left; i <= right; i++ {
			temp := vals[left-1] * vals[i] * vals[right+1]
			temp += rangeSearch(left, i-1) + rangeSearch(i+1, right)
			if coins < temp {
				coins = temp
			}
		}
		dp[[2]int{left, right}] = coins
		return coins
	}
	return rangeSearch(1, len(vals)-2)
}
