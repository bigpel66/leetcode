// 1022. Sum of Root To Leaf Binary Numbers
//
// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// sumRootToLeaf function accumulates all of the numbers root to leaf as a binary.
// Preorder traverse is required to calculate the root to leaf number.
func sumRootToLeaf(root *TreeNode) int {
	var accumulate func(int, *TreeNode) int
	accumulate = func(bin int, node *TreeNode) int {
		if node == nil {
			return 0
		}
		bin = bin<<1 | node.Val
		if node.Left == nil && node.Right == nil {
			return bin
		}
		return accumulate(bin, node.Left) + accumulate(bin, node.Right)
	}
	return accumulate(0, root)
}
