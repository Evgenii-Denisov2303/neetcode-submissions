func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens  {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := 0

			switch token {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				result = left / right
			}

			stack = append(stack, result)

		} else {
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}

	return stack[len(stack)-1]
}
