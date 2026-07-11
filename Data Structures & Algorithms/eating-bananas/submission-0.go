func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0

	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left < right {
		mid := left + (right-left)/2

		var hours int64

		for _, pile := range piles {
			hours += int64((pile + mid - 1) / mid)
		}

		if hours > int64(h) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
