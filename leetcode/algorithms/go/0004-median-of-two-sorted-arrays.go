package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var len1 = len(nums1)
	var len2 = len(nums2)
	var length = len1 + len2

	var mid = length / 2
	var midLeft = mid - 1

	var last, left, right int

	var i, j, k int

	for k <= mid && i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			last = nums1[i]
			i++
		} else {
			last = nums2[j]
			j++
		}
		if k == midLeft {
			left = last
		}
		if k == mid {
			right = last
		}
		k++
	}

	for k <= mid && i < len1 {
		last = nums1[i]
		if k == midLeft {
			left = last
		}
		if k == mid {
			right = last
		}
		i++
		k++
	}

	for k <= mid && j < len2 {
		last = nums2[j]
		if k == midLeft {
			left = last
		}
		if k == mid {
			right = last
		}
		j++
		k++
	}

	if length%2 == 0 {
		return float64(left+right) / 2.0
	}

	return float64(right)
}

// import "sort"
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	var result float64
// 	nums1 = append(nums1, nums2...)
// 	sort.Ints(nums1)
// 	var length = len(nums1)
// 	if length%2 != 0 {
// 		result = float64(nums1[length/2])
// 	} else {
// 		right := length - 1
// 		result = float64(nums1[right/2]+nums1[right/2+1]) / 2.0
// 	}
// 	return result
// }

// odd
// mid == 5 / 2 == 2
// median == nums[mid]
// [1, 2, 3, 4, 5]
// [0, 1, 2, 3, 4]
// [0, 1, !, 3, 4]

// even
// mid == 6 / 2 == 3
// midLeft  == 3 - 1 == 2
// median == (nums[midLeft] + nums[mid]) / 2
// [1, 2, 3, 4, 5, 6]
// [0, 1, 2, 3, 4, 5]
// [0, 1, !, !, 4, 5]

// last
// last == sortedLast