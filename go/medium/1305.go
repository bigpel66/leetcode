// 1305. All Elements in Two Binary Search Trees
//
// https://leetcode.com/problems/all-elements-in-two-binary-search-trees/

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// getAllElements function returns the all elements in ascending order by using 2 Binary Tree.
// Logic to solve this problem is just traverse both trees each and memoize into the slice.
// Finally, two slices can be returned.
// To return the answer, two slices should be integrated into a sole slice.
// This can be done by the core logic of merge sort.
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	list1, list2 := &[]int{}, &[]int{}
	var depthSearch func(node *TreeNode, list *[]int)
	depthSearch = func(node *TreeNode, list *[]int) {
		if node == nil {
			return
		}
		depthSearch(node.Left, list)
		*list = append(*list, node.Val)
		depthSearch(node.Right, list)
	}
	var merge func(list1 []int, list2 []int) []int
	merge = func(list1 []int, list2 []int) []int {
		result := make([]int, 0, len(list1)+len(list2))
		i1, i2 := 0, 0
		for i1 < len(list1) && i2 < len(list2) {
			if list1[i1] < list2[i2] {
				result = append(result, list1[i1])
				i1++
			} else {
				result = append(result, list2[i2])
				i2++
			}
		}
		for i1 < len(list1) {
			result = append(result, list1[i1])
			i1++
		}
		for i2 < len(list2) {
			result = append(result, list2[i2])
			i2++
		}
		return result
	}
	depthSearch(root1, list1)
	depthSearch(root2, list2)
	return merge(*list1, *list2)
}
