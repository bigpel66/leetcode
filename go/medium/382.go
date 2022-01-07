// 382. Linked List Random Node
//
// https://leetcode.com/problems/linked-list-random-node/

// math/rand package has been used to call Intn.
import (
	"math/rand"
)

// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// Solution has the length of input list, and its contents.
type Solution struct {
	length  int
	content []int
}

// Constructor function initialize the Solution structure with the given list.
func Constructor(head *ListNode) Solution {
	sol := Solution{}
	for head != nil {
		sol.length++
		sol.content = append(sol.content, head.Val)
		head = head.Next
	}
	return sol
}

// GetRandom function returns the value of the random index node.
// This is done by calling rand package funciton.
// A index range is the length which is held by Solution structure.
// Solution structure already knows the content of the each index, that is the answer.
func (this *Solution) GetRandom() int {
	index := rand.Intn(this.length)
	return this.content[index]
}

// Your Solution object will be instantiated and called as such:
// obj := Constructor(head);
// param_1 := obj.GetRandom();
