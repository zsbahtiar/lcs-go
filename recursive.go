package lcs_go

func recursive(a, b string) string {
	x := len(a)
	y := len(b)

	if x == 0 || y == 0 {
		return ""
	}

	aRune := []rune(a)
	bRune := []rune(b)

	if aRune[x-1] == bRune[y-1] {
		return recursive(a[:x-1], b[:y-1]) + string(aRune[x-1])
	}

	r1 := recursive(a[:x-1], b)
	r2 := recursive(a, b[:y-1])

	if len(r1) > len(r2) {
		return r1
	}

	return r2
}
