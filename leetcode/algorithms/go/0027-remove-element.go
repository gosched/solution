package leetcode

func removeElement(nums []int, val int) int {
	var i int
	for j, num := range nums {
		if num != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return i
}
