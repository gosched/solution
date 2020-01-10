package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var dummy = &ListNode{}
	var mark = dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			mark.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			mark.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		mark = mark.Next
	}

	if l1 != nil {
		mark.Next = l1
	}

	if l2 != nil {
		mark.Next = l2
	}

	return dummy.Next
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var result *ListNode

// 	if l1 == nil {
// 		return l2
// 	}

// 	if l2 == nil {
// 		return l1
// 	}

// 	if l1.Val < l2.Val {
// 		result = l1
// 		result.Next = mergeTwoLists(l1.Next, l2)
// 	} else {
// 		result = l2
// 		result.Next = mergeTwoLists(l2.Next, l1)
// 	}

// 	return result
// }
