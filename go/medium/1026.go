// 1026. Maximum Difference Between Node and AncestorA
//
// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// diff function returns the absolute value of difference between a and b.
func diff(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

// min function finds the minimum value among the elements in the slice.
func min(num []int) int {
	val := int(^uint(0) >> 1)
	for _, n := range num {
		if val > n {
			val = n
		}
	}
	return val
}

// maxAncestorDiff function finds the maximum difference between any ancestor and its decendants.
// Divide and Conquer from the leaf node finds the minimum value under current node.
// The Difference between the found minimum value and the value of current node updates the value to be answer.
func maxAncestorDiff(root *TreeNode) int {
	answer := -1
	var recursion func(*TreeNode, int) int
	recursion = func(node *TreeNode, val int) int {
		if node == nil {
			return val
		}
		val = min([]int{val, node.Val})
		l, r := recursion(node.Left, val), recursion(node.Right, val)
		val = min([]int{val, l, r})
		if d := diff(node.Val, val); answer < d {
			answer = d
		}
		return val
	}
	recursion(root, int(^uint(0)>>1))
	return answer
}
