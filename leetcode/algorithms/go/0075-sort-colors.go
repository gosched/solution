package leetcode

func sortColors(nums []int) {
	r, w, b := 0, 0, len(nums)-1
	for w <= b {
		if nums[w] == 0 {
			nums[w], nums[r] = nums[r], nums[w]
			r++
			w++
		} else if nums[w] == 2 {
			nums[w], nums[b] = nums[b], nums[w]
			b--
		} else {
			w++
		}
	}
}

// func sortColors(nums []int) {
// 	red, white, blue := 0, 0, len(nums)-1
// 	for white <= blue {
// 		if nums[white] == 0 {
// 			nums[white], nums[red] = nums[red], nums[white]
// 			red++
// 			white++
// 		} else if nums[white] == 2 {
// 			nums[white], nums[blue] = nums[blue], nums[white]
// 			blue--
// 		} else {
// 			white++
// 		}
// 	}
// }

// https://en.wikipedia.org/wiki/Dutch_national_flag_problem
