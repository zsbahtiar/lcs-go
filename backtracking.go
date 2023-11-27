package lcs_go

type stackNode struct {
	x, y    int
	current string
}

func backtracking(a, b string) string {
	aRune := []rune(a)
	bRune := []rune(b)

	stack := []stackNode{
		{
			x:       len(a),
			y:       len(b),
			current: "",
		},
	}

	var result string

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := top.x, top.y

		if x == 0 || y == 0 {
			if len(top.current) > len(result) {
				result = top.current
			}
			continue
		}

		if aRune[x-1] == bRune[y-1] {
			temp := string(aRune[x-1]) + top.current
			stack = append(stack, stackNode{
				x:       x - 1,
				y:       y - 1,
				current: temp,
			})
		} else {
			stack = append(stack, stackNode{
				x:       x - 1,
				y:       y,
				current: top.current,
			})
			stack = append(stack, stackNode{
				x:       x,
				y:       y - 1,
				current: top.current,
			})
		}
	}

	return result
}
