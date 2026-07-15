func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left := 0
	right := len(height) - 1

	leftMax := height[left]
	rightMax := height[right]

	totalWater := 0

	for left < right {
		if leftMax <= rightMax {
			left++

			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
		} else {
			right--

			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
		}
	}

	return totalWater
}