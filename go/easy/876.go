// 876. Middle of the Linked List
//
// https://leetcode.com/problems/middle-of-the-linked-list/

// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// middleNode function finds the Middle Node from the given linked list.
// The specified fast node is incremented by 2.
// When the fast node cannot be incremented, then head node is the answer.
func middleNode(head *ListNode) *ListNode {
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		head = head.Next
	}
	return head
}
