package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	length := len(nums)
	limit := length - 2
	sort.Ints(nums)
	for i := 0; i < limit; i++ {
		if 0 < i && nums[i] == nums[i-1] {
			continue
		}
		indexB := i + 1
		indexC := length - 1
		for indexB < indexC {
			a := nums[i]
			b := nums[indexB]
			c := nums[indexC]
			if a+b+c < 0 {
				indexB++
			} else if a+b+c > 0 {
				indexC--
			} else {
				result = append(result, []int{a, b, c})
				for indexB < indexC && nums[indexB] == nums[indexB+1] {
					indexB++
				}
				for indexB < indexC && nums[indexC] == nums[indexC-1] {
					indexC--
				}
				indexB++
				indexC--
			}
		}
	}
	return result
}
