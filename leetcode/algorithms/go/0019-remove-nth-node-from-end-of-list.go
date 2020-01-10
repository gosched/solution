package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	c := 0
	t := head
	for t != nil {
		t = t.Next
		c++
	}

	if n == c {
		backup := head.Next
		head.Next = nil
		return backup
	}

	t = head
	for i := 0; i < c-n-1; i++ {
		t = t.Next
	}
	if t.Next != nil {
		t.Next = t.Next.Next
	} else {
		t.Next = nil
	}

	return head
}
