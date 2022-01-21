// 142. Linked List Cycle II
//
// https://leetcode.com/problems/linked-list-cycle-ii/

// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// detectCylce function detects whether the linked list has a cycle or not.
// Golang supports a pointer operation and taking address, so a hash map for address has been used.
// Maybe slow-fast algorithm which technique to reach the middle of the linked list can be used to check the list cyclic or not.
func detectCycle(head *ListNode) *ListNode {
	addressMap := make(map[*ListNode]bool)
	for head != nil {
		if !addressMap[head] {
			addressMap[head] = true
			head = head.Next
		} else {
			return head
		}
	}
	return nil
}
