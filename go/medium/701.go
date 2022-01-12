// 701. Insert into a Binary Search Tree
//
// https://leetcode.com/problems/insert-into-a-binary-search-tree/

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// insertIntoBST function insert a new node with value val into the binary tree.
// Fixup way can be used, but inserting the node to the leaf have been used in this case.
// If there's no node on the tree, just return a new node.
// Else, the child of each node should be validated to insert a new node. This can be done by recursion and assignation.
// Anyway no rotation has been used, just return the root node is the answer after insertion.
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
