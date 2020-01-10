package leetcode

func nextPermutation(nums []int) {
	var n = len(nums)
	var i = n - 2
	for -1 < i && nums[i+1] <= nums[i] {
		i--
	}
	if i != -1 {
		var j = n - 1
		for 0 < j && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	var l, r = i + 1, n - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
