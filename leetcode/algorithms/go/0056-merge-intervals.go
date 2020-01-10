package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result = [][]int{}
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
		} else if result[len(result)-1][1] < intervals[i][1] {
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	return result
}
