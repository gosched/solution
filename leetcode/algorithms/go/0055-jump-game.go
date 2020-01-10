package leetcode

func canJump(nums []int) bool {
	var reach int
	var last = len(nums) - 1
	for i, step := range nums {
		if reach < i || last <= reach {
			break
		}
		if reach < i+step {
			reach = i + step
		}
	}
	return last <= reach
}

// Greedy algorithm
// reach 代表最遠能到的位置, 我們只關心最遠能到的位置

// 到此為止無需再試
// 最遠能到的位置小於當前位置, 無法前進, 可能到終點, 也可能沒到終點
// 最遠能到的位置大於等於終點, 無需前進, 可抵達終點

// 調整最遠能到的位置
// 我不走這 or 我走這, 我在 i, 我能再走 step 步, 我能到 i + step

// 判斷是否能抵達終點
// 最遠能到的位置是否超過或等於終點的位置
