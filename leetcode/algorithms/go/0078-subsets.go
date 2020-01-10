package leetcode

func subsets(nums []int) [][]int {
	var result = [][]int{}
	result = append(result, []int{}) // empty set
	for _, num := range nums {
		size := len(result)
		for j := 0; j < size; j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			temp = append(temp, num)
			result = append(result, temp)
		}
	}
	return result
}

// import "sort"

// func subsets(nums []int) [][]int {
// 	var result = [][]int{}
// 	result = append(result, []int{}) // empty set
// 	sort.Ints(nums)
// 	for _, num := range nums {
// 		size := len(result)
// 		for j := 0; j < size; j++ {
// 			temp := make([]int, len(result[j]))
// 			copy(temp, result[j])
// 			temp = append(temp, num)
// 			result = append(result, temp)
// 		}
// 	}
// 	return result
// }

// []
// [], [1]
// [], [1], [2], [1,2]
// [], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]

// func subsets(nums []int) [][]int {
// 	var result = [][]int{}
// 	var helper func(mark int, temp []int)
// 	helper = func(mark int, temp []int) {
// 		set := make([]int, len(temp), len(temp))
// 		copy(set, temp)
// 		result = append(result, set)
// 		for i := mark; i < len(nums); i++ {
// 			temp = append(temp, nums[i])
// 			helper(i+1, temp)
// 			temp = temp[:len(temp)-1]
// 		}
// 	}
// 	sort.Ints(nums)
// 	helper(0, []int{})
// 	return result
// }

// []
// [] [1]
// [] [1] [1 2]
// [] [1] [1 2] [1 2 3]
// [] [1] [1 2] [1 2 3] [1 3]
// [] [1] [1 2] [1 2 3] [1 3] [2]
// [] [1] [1 2] [1 2 3] [1 3] [2] [2 3]
// [] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]
