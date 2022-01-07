// 563. Binary Tree Tilt
//
// https://leetcode.com/problems/binary-tree-tilt/

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// search function search the absolute difference on global variable result.
// Each node find the difference of the left and right child, and add to the result.
// To record the difference correctly, every value of the left and right child should be return to the upward.
// When the recursive call of this function to the end, every tilt value can be accumulated.
func search(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left, right := search(root.Left, result), search(root.Right, result)
	diff := left - right
	if diff < 0 {
		diff *= -1
	}
	*result += diff
	return left + right + root.Val
}

func findTilt(root *TreeNode) int {
	result := 0
	_ = search(root, &result)
	return result
}
