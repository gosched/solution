package leetcode

func trap(height []int) int {
    var result, leftMax, rightMax int
    var i, j = 0, len(height) - 1
    for i < j {
        if height[i] < height[j] {
            if height[i] < leftMax {
                result += leftMax - height[i]
            } else {
                leftMax = height[i]
            }
            i++
        } else {
            if height[j] < rightMax {
                result += rightMax - height[j]
            } else {
                rightMax = height[j]
            }
            j--
        }
    }
    return result
}