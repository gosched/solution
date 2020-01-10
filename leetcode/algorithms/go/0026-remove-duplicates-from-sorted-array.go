package leetcode

func removeDuplicates(nums []int) int {
	var i, n = 0, len(nums)

	if n < 2 {
		return n
	}

	for j := 0; j < n; j++ {
		for j < n-1 && nums[j] == nums[j+1] {
			j++
		}
		nums[i] = nums[j]
		i++
	}

	return i
}
