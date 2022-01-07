// 147. Insertion Sort List
//
// https://leetcode.com/problems/insertion-sort-list/description/

// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// insertionSortList function sort the element of ListNode with ascending order.
// ListNode will be sorted by insertion sort as its function name.
// To find the place of the element, cursor moves from the start of the dummy to the end.
// Dummy holds the empty ListNode, so the answer is the Next of Dummy.
func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		cur := dummy
		for cur.Next != nil && cur.Next.Val < head.Val {
			cur = cur.Next
		}
		cur.Next, head.Next, head = head, cur.Next, head.Next
	}
	return dummy.Next
}
