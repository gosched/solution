package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var result = []int{}
	var stack = []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}

// func inorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	if root == nil {
// 		return result
// 	}
// 	var helper func(node *TreeNode)
// 	helper = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		helper(node.Left)
// 		result = append(result, node.Val)
// 		helper(node.Right)
// 	}
// 	helper(root)
// 	return result
// }

// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	result := []int{}
//     result = append(result, inorderTraversal(root.Left)...)
//     result = append(result, root.Val)
//     result = append(result, inorderTraversal(root.Right)...)
//     return result
// }

// slice = append(slice, nil...)
// 是合法的 不會影響 slice
