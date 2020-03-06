package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	set := [][]int{}
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		a := nums[i]
		if 0 < a {
			break
		}
		if 0 < i && nums[i] == nums[i-1] {
			continue
		}
		indexB, indexC := i+1, length-1
		for indexB < indexC {
			b, c := nums[indexB], nums[indexC]
			if a+b+c < 0 {
				indexB++
			} else if a+b+c > 0 {
				indexC--
			} else {
				set = append(set, []int{a, b, c})
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
	return set
}
