package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var largest = nums[0]
	var sum = 0
	for _, num := range nums {
		if sum+num < num {
			sum = num
		} else {
			sum += num
		}
		if largest < sum {
			largest = sum
		}
	}
	return largest
}

// func maxSubArray(nums []int) int {
//     var n = len(nums)
//     var dp = make([]int, n, n)
//     dp[0] = nums[0]
//     for i := 1; i < n; i++ {
//         dp[i] = max(dp[i-1]+nums[i], nums[i])
//     }
//     return largest(dp)
// }

// func max(x, y int) int {
//     if x < y {
//         return y
//     }
//     return x
// }

// func largest(items []int) int {
//     if len(items) == 0 {
//         return 0
//     }
//     var max = items[0]
//     for _, item := range items {
//         if max < item {
//             max = item
//         }
//     }
//     return max
// }
