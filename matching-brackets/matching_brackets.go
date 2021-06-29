package brackets

var openningFor map[byte]byte = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func Bracket(input string) (ret bool) {
	var stack []byte
	for _, c := range input {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, byte(c))
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) != 0 && stack[len(stack)-1] == openningFor[byte(c)] {
				stack = stack[:len(stack)-1]
			} else {
				return
			}
		}
	}
	if len(stack) == 0 {
		ret = true
	}

	return
}