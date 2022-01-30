// 189. Rotate Array
//
// https://leetcode.com/problems/rotate-array/

// rotate function rotates the nums slice by iterating the number with k.
// Time complexity O(n) & non-additional memory is the way with reserve the slice several times.
// For about k, only k mod the length of the nums slice is required to rotate the slice.
// Reverse index 0 to len(nums) - 1.
// Reverse index 0 to k - 1.
// Reverse index k to len(nums) - 1.
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
