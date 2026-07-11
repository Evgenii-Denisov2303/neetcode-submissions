func isValid(s string) bool {
    stack := []rune{}

	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, ch := range s {
		if _, ok := pairs[ch]; ok {
			stack = append(stack, ch)
			continue
		} 

		if len(stack) == 0 {
			return false
		}
		
		last := stack[len(stack)-1]

		if ch != pairs[last] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
