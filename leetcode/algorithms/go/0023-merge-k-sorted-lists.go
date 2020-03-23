package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	m := len(lists) / 2
	return marge(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
}

func marge(l1, l2 *ListNode) *ListNode {
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
			mark.Next = l1
			l1 = l1.Next
		} else {
			mark.Next = l2
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

// Heap
// import "container/heap"

// type NodeHeap []*ListNode

// func (h NodeHeap) Len() int { return len(h) }
// func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
// func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// func (h *NodeHeap) Push(x interface{}) {
//     *h = append(*h, x.(*ListNode))
// }

// func (h *NodeHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[0:n-1]
//     return x
// }

// func mergeKLists(lists []*ListNode) *ListNode {
//     if (len(lists) == 0) {
//         return nil
//     }

//     if (len(lists) == 1) {
//         return lists[0]
//     }

//     dummy := &ListNode{}
//     mark := dummy

//     h := &NodeHeap{}
//     heap.Init(h)

//     for _, list := range lists {
//         if list != nil {
//             heap.Push(h, list)
//         }
//     }

//     for h.Len() > 0 {
//         node := heap.Pop(h).(*ListNode)
//         if node.Next != nil {
//             heap.Push(h, node.Next)
//         }
//         mark.Next = node
//         mark = mark.Next
//     }

//     return dummy.Next
// }

// PriorityQueue
// import (
//     "container/heap"
// )

// type Item struct {
//     value *ListNode
//     priority int
//     index int
// }

// type NodeHeap []*Item

// func (h NodeHeap) Len() int { return len(h) }
// func (h NodeHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
// func (h NodeHeap) Swap(i, j int) {
//     h[i], h[j] =  h[j], h[i]
//     h[i].index = i
//     h[j].index = j
// }

// func (h *NodeHeap) Push(x interface{}) {
//     n := len(*h)
//     item := &Item {
//         value : x.(*ListNode),
//         index : n,
//     }
//     item.priority = item.value.Val
//     *h = append(*h, item)
// }

// func (h *NodeHeap) Pop() interface{} {
//     old := *h
//     n := len(old)

//     item := old[n-1]
//     item.index = -1

//     *h = old[0 : n-1]

//     return item
// }

// func mergeKLists(lists []*ListNode) *ListNode {
//     dummy := &ListNode{}
//     mark := dummy

//     h := &NodeHeap{}
//     heap.Init(h)

//     for _, list := range lists {
//         if list != nil {
//             heap.Push(h, list)
//         }
//     }

//     for h.Len() > 0 {
//         node := heap.Pop(h).(*Item).value
//         if node.Next != nil {
//             heap.Push(h, node.Next)
//         }
//         mark.Next = node
//         mark = mark.Next
//     }

//     return dummy.Next
// }

// import "sort"

// func mergeKLists(lists []*ListNode) *ListNode {
// 	var x = []int{}
// 	for i := 0; i < len(lists); i++ {
// 		mark := lists[i]
// 		for mark != nil {
// 			x = append(x, (*mark).Val)
// 			mark = (*mark).Next
// 		}
// 	}
// 	sort.Ints(x)
// 	var dummy = &ListNode{}
// 	var mark = dummy
// 	for i := 0; i < len(x); i++ {
// 		mark.Next = &ListNode{Val: x[i]}
// 		mark = mark.Next
// 	}
// 	return dummy.Next
// }
