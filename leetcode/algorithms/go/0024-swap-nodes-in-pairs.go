package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head

	var prev = dummy
	var current = head
	var nextPair *ListNode

	for current != nil && current.Next != nil {
		nextPair = current.Next.Next

		current.Next.Next = current
		prev.Next = current.Next

		current.Next = nextPair

		prev = current
		current = nextPair
	}

	return dummy.Next
}

// 1->2->3->4 ---> 2->1->3->4 ---> 2->1->4->3
