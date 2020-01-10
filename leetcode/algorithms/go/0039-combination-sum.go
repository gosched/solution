package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var result = [][]int{}
	var helper func(candidates []int, mark, target int, temp []int)
	helper = func(candidates []int, mark, target int, temp []int) {
		if target < 0 {
			return
		}
		if target == 0 {
			var items = make([]int, len(temp))
			copy(items, temp)
			result = append(result, items)
			return
		}
		for i := mark; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			helper(candidates, i, target-candidates[i], temp)
			temp = temp[:len(temp)-1]
		}
	}
	helper(candidates, 0, target, []int{})
	return result
}

// func combinationSum(candidates []int, target int) [][]int {
// 	var result = [][]int{}
// 	sort.Ints(candidates)
// 	var helper func(candidates []int, mark, target int, temp []int)
// 	helper = func(candidates []int, mark, target int, temp []int) {
// 		if target < 0 {
// 			return
// 		}
// 		if target == 0 {
// 			var items = make([]int, len(temp))
// 			copy(items, temp)
// 			result = append(result, items)
// 		}
// 		for i := mark; i < len(candidates); i++ {
// 			if target < candidates[i] {
// 				break
// 			}
// 			temp = append(temp, candidates[i])
// 			helper(candidates, i, target-candidates[i], temp)
// 			temp = temp[:len(temp)-1]
// 		}
// 	}
// 	helper(candidates, 0, target, []int{})
// 	return result
// }

// Depth-first search

// []
// [2], [3], [6], [7!]
// [2,2], [2,3], [2,6], [2,7], [3,3], [3,6], [3,7], [6,6], [6,7], [7,7]
// [2,2,2], [2,2,3!], [2,2,6], [2,2,7], [2,3,3], [2,3,6], [2,3,7], [3,3,3], [3,3,6], [3,3,7]
// [2,2,2,2], [2,2,2,3], [2,2,2,6], [2,2,2,7], [2,2,3,3], [2,2,3,6], [2,2,3,7]

// []
// [2]
// [2 2]
// [2 2 2]
// [2 2 2 2]
// [2 2 2 3]
// [2 2 2 6]
// [2 2 2 7]
// [2 2 3]
// [2 2 3 3]
// [2 2 3 6]
// [2 2 3 7]
// [2 2 6]
// [2 2 7]
// [2 3]
// [2 3 3]
// [2 3 6]
// [2 3 7]
// [2 6]
// [2 7]
// [3]
// [3 3]
// [3 3 3]
// [3 3 6]
// [3 3 7]
// [3 6]
// [3 7]
// [6]
// [6 6]
// [6 7]
// [7]
// [7 7]
