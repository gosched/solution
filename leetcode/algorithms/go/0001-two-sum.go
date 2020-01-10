package leetcode

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	var m = make(map[int]int, len(nums))
// 	for i, num := range nums {
// 		if j, ok := m[target-num]; ok && j != i {
// 			return []int{j, i}
// 		}
// 		m[num] = i
// 	}
// 	return nil
// }
