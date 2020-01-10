package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var mark = dummy
	var carry = 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		mark.Next = &ListNode{Val: carry % 10}
		mark = mark.Next
		carry /= 10
	}
	if carry != 0 {
		mark.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
