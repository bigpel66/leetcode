// 416. Partition Equal Subset Sum
//
// https://leetcode.com/problems/partition-equal-subset-sum/
//
// 	When the numbers are given, make two partitions such that each summation is equal.
// 	Numbers will be given with the only positive integers, and non-empty array.

// canPartition function checks the given numbers can be grouped or not.
// When 2 groups which summations are equal, it implies that total summation is even.
// After that, this function get the truth of each number whether it can be made with the element of the given array or not.
// In the outer loop, the element of the array has been used in sequence.
// To get the truth whether the number can be made, decrement inner loop has been used.
//  -> when the given array is [ 1, 2, 3, 5, 11 ]
//  -> first : 1 by dp[0] (using 1)
//  -> second : 1, 2, 3 by dp[0], dp[1] (using 2)
//  -> third : 1, 2, 3, 4, 5, 6 by dp[0] to dp[3] (using 3)
//  -> ...
func canPartition(nums []int) bool {
	size := len(nums)
	acc := 0
	for i := 0; i < size; i++ {
		acc += nums[i]
	}
	if acc%2 == 1 {
		return false
	}
	goal := acc / 2
	dp := make([]bool, acc+1)
	dp[0] = true
	for i := 0; i < size; i++ {
		for j := goal; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[goal]
}
