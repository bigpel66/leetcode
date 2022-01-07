// 116. Populating Next Right Pointers in Each Node
//
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

// Definition for a Node.
// type Node struct {
//     Val int
//     Left *Node
//     Right *Node
//     Next *Node
// }

// connect function connects current node to the next node.
// Recursion has been used to connect the next node.
// Also, connects if left node has right child and right node has left child.
//            1                     1                    1
//        2   -   3             2   -   3            2   -   3
//      4   5   6   7         4   5 -  6  7        4 - 5 - 6 - 7
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left != nil && current.Right != nil {
			current.Left.Next = current.Right
			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}
			queue = append(queue, current.Left)
			queue = append(queue, current.Right)
		}
	}
	return (root)
}
