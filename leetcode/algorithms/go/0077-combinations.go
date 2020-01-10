package leetcode

func combine(n int, k int) [][]int {
	var result = [][]int{}
	helper(n, k, 1, []int{}, &result)
	return result
}

func helper(n, k, level int, output []int, result *[][]int) {
	if len(output) == k {
		items := make([]int, len(output))
		copy(items, output)
		*result = append(*result, items)
		return
	}
	for i := level; i <= n; i++ {
		output = append(output, i)
		helper(n, k, i+1, output, result)
		output = output[:len(output)-1]
	}
}

// [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]

// [] -> [1] -> [1,2],[1,3],[1,4] -> []
// [] -> [2] -> [2,3],[2,4]       -> []
// [] -> [3] -> [3,4]             -> []
// [] -> [4]                      -> []
