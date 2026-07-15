func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)
	maxLength := 0

	for _, num := range nums {
		set[num] = true
	}

	for num := range set {

		if !set[num-1] {
			currentNumber := num
			currentLength := 1

			for set[currentNumber+1] {
				currentNumber++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}
