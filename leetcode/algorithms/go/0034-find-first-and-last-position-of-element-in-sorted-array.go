package leetcode

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m
		} else {
			i, j := m, m
			for i > 0 && nums[i-1] == target {
				i--
			}
			for j < len(nums)-1 && nums[j+1] == target {
				j++
			}
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}

// func searchRange(nums []int, target int) []int {
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if nums[mid] < target {
// 			l = mid + 1
// 		} else if nums[mid] > target {
// 			r = mid - 1
// 		} else {
// 			i, j := mid, mid
// 			for i > 0 && nums[i-1] == target {
// 				i--
// 			}
// 			for j < len(nums)-1 && nums[j+1] == target {
// 				j++
// 			}
// 			return []int{i, j}
// 		}
// 	}
// 	return []int{-1, -1}
// }
