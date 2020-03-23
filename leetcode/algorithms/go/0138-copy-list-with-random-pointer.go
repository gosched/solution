package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	var m = make(map[*Node]*Node)
	var mark = head
	for mark != nil {
		cloned := &Node{Val: mark.Val}
		m[mark] = cloned
		mark = mark.Next
	}
	mark = head
	for mark != nil {
		m[mark].Next = m[mark.Next]
		m[mark].Random = m[mark.Random]
		mark = mark.Next
	}
	return m[head]
}

// 跑 loop 遍歷 original
// 創造 cloned, 備份 Val
// 透過 map 令 original 與 cloned 對應

// 跑 loop 遍歷 original
// 透過 map 操作 cloned, original
// 為 cloned 補上 Next 與 Random
