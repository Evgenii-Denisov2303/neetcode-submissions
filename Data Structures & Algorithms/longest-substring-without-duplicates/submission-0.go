func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLength := 0

	seen := make(map[byte]bool)

	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		seen[s[right]] = true

		length := right - left + 1

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
