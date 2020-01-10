package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack = []*TreeNode{}
	var prev *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root
		root = root.Right
	}

	return true
}

// import "math"

// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	min := int(math.MinInt64)
// 	max := int(math.MaxInt64)

// 	var helper func(root *TreeNode, min, max int) bool
// 	helper = func(root *TreeNode, min, max int) bool {
// 		if root == nil {
// 			return true
// 		}
// 		if root.Val <= min {
// 			return false
// 		}
// 		if root.Val >= max {
// 			return false
// 		}
// 		return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
// 	}

// 	return helper(root, min, max)
// }
