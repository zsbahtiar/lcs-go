package lcs_go

import (
	"strings"
	"unicode/utf8"
)

type stackNode struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Current string `json:"current"`
}

func backtracking(a, b string) string {
	aRune := []rune(a)
	bRune := []rune(b)

	stack := []stackNode{
		{
			X:       0,
			Y:       0,
			Current: "",
		},
	}
	var result string

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := top.X, top.Y

		if x == utf8.RuneCountInString(a) || y == utf8.RuneCountInString(b) {
			if len(top.Current) > len(result) {
				result = top.Current
			}
			continue
		}

		if strings.ToLower(string(aRune[x])) == strings.ToLower(string(bRune[y])) {
			temp := top.Current + string(aRune[x])
			stack = append(stack, stackNode{
				X:       x + 1,
				Y:       y + 1,
				Current: temp,
			})
		} else {
			stack = append(stack, stackNode{
				X:       x + 1,
				Y:       y,
				Current: top.Current,
			})
			stack = append(stack, stackNode{
				X:       x,
				Y:       y + 1,
				Current: top.Current,
			})
		}
	}

	return result
}
