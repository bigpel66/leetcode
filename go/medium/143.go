// 143. Reorder List
//
// https://leetcode.com/problems/reorder-list/

// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// reorderList function performs just as its name implies.
// When list exists such as 1 -> 2 -> 3 -> 4, the result of this function becomes 1 -> 4-> 2 -> 3.
// Remember that this function use the ListNode as singly-linked list.
// To reorder the list, the first thing to do is find the half point of the list.
// Then, flip the order of the list from the half point of the list.
// When the first half of the list is forward direction and the second half of the list is reverse direction, then merge then into the desired order.
// This function does not require additional memory and its time complexity is only O(n).
func reorderList(head *ListNode) {
	flipPos, controlPos := head, head
	for controlPos != nil && controlPos.Next != nil && controlPos.Next.Next != nil {
		flipPos, controlPos = flipPos.Next, controlPos.Next.Next
	}
	var prevPos *ListNode
	for flipPos != nil {
		flipPos.Next, prevPos, flipPos = prevPos, flipPos, flipPos.Next
	}
	temp := head
	for temp != nil && prevPos != nil {
		temp.Next, temp = prevPos, temp.Next
		prevPos.Next, prevPos = temp, prevPos.Next
	}
}
