package lcs_go

func dp(a, b string) string {
	x := len(a)
	y := len(b)
	aRune := []rune(a)
	bRune := []rune(b)

	var result string
	table := make([][]int, x+1)
	for i := range table {
		table[i] = make([]int, y+1)
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if aRune[i-1] == bRune[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	m := x
	n := y

	for m > 0 && n > 0 {
		if aRune[m-1] == bRune[n-1] {
			result = string(aRune[m-1]) + result
			m--
			n--
		} else if table[m-1][n] > table[m][n-1] {
			m--
		} else {
			n--
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
