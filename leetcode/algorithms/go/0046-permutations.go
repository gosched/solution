package leetcode

func permute(nums []int) [][]int {
	var result = [][]int{}
	var helper func(nums []int, start int)
	helper = func(nums []int, start int) {
		var n = len(nums)
		if start == n {
			var items = make([]int, n)
			copy(items, nums)
			result = append(result, items)
			return
		}
		for i := start; i < n; i++ {
			nums[start], nums[i] = nums[i], nums[start]
			helper(nums, start+1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	helper(nums, 0)
	return result
}

// func permute(nums []int) [][]int {
// 	var result = [][]int{}
// 	helper(nums, 0, &result)
// 	return result
// }

// func helper(nums []int, start int, result *[][]int) {
// 	var n = len(nums)
// 	if start == n {
// 		var items = make([]int, n, n)
// 		copy(items, nums)
// 		*result = append(*result, items)
// 	}
// 	for i := start; i < n; i++ {
// 		nums[start], nums[i] = nums[i], nums[start]
// 		helper(nums, start+1, result)
// 		nums[start], nums[i] = nums[i], nums[start]
// 	}
// }

// 1,2,3,4,5  // 1,2,3,4,5  // 1,2,3,4,5  // 1,2,3,4,5  // 1,2,3,4,5
// 1,2,3,4,5  // 2,1,3,4,5  // 3,2,1,4,5  // 4,2,3,1,5  // 5,2,3,4,1
// 1,2,3,4,5  // 1,2,3,4,5  // 1,2,3,4,5  // 1,2,3,4,5  // 1,2,3,4,5

// 1;2,3,4,5  // 1;2,3,4,5  // 1;2,3,4,5 // 1;2,3,4,5
// 1;2,3,4,5  // 1;3,2,4,5  // 1;4,3,2,5 // 1;5,3,4,2
// 1;2,3,4,5  // 1;2,3,4,5  // 1;2,3,4,5 // 1;2,3,4,5

// 1,2;3,4,5  // 1,2;3,4,5  // 1,2;3,4,5
// 1,2;3,4,5  // 1,2;4,3,5  // 1,2;5,4,3
// 1,2;3,4,5  // 1,2;3,4,5  // 1,2;3,4,5

// 1,2,3;4,5  // 1,2,3;4,5
// 1,2,3;4,5  // 1,2,3;5,4
// 1,2,3;4,5  // 1,2,3;4,5

// 1,2,3,4;5
// 1,2,3,4;5
// 1,2,3,4;5

// ........

// [[1 2 3]]
// [[1 2 3] [1 3 2]]
// [[1 2 3] [1 3 2] [2 1 3]]
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1]]
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1]]
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]

// func permute(nums []int) [][]int {
// 	var result = [][]int{}
// 	var n = len(nums)
// 	var limit = factorial(n)
// 	for i := 0; i < limit; i++ {
// 		result = append(result, nums)
// 		var items = make([]int, n, n)
// 		copy(items, nums)
// 		nums = nextPermutation(items)
// 	}
// 	return result
// }

// func factorial(n int) int {
// 	if n > 0 {
// 		return n * factorial(n-1)
// 	}
// 	return 1
// }

// func nextPermutation(nums []int) []int {
// 	var n = len(nums)
// 	var i = n - 2
// 	for i > -1 && nums[i+1] <= nums[i] {
// 		i--
// 	}
// 	if i != -1 {
// 		var j = n - 1
// 		for j > 0 && nums[j] <= nums[i] {
// 			j--
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	reverse(nums, i+1, n-1)
// 	return nums
// }

// func reverse(nums []int, i, j int) {
// 	for i < j {
// 		nums[i], nums[j] = nums[j], nums[i]
// 		i++
// 		j--
// 	}
// }

// len(nums) == n
// len(result) == n!

// [[1 2 3]]
// [[1 2 3] [1 3 2]]
// [[1 2 3] [1 3 2] [2 1 3]]
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1]]
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2]]
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
