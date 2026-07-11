func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxArea := 0

	for left < right {
		width := right - left

		height := min(heights[left], heights[right])

		area := width * height

		if area > maxArea {
			maxArea = area
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
