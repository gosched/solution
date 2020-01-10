package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var length = len(nums)

	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var result = nums[0] + nums[1] + nums[2]

	for i := 0; i < length; i++ {
		var start, end = i + 1, length - 1
		for start < end {
			var sum = nums[i] + nums[start] + nums[end]
			if abs(target-sum) < abs(target-result) {
				result = sum
			}
			if sum < target {
				start++
			} else if sum > target {
				end--
			} else {
				return result
			}
		}
	}

	return result
}
