package lcs_go

import "testing"

func BenchmarkDp(t *testing.B) {
	for i := 0; i < t.N; i++ {
		dp("COMPUTER", "SCIENCE")
	}

}

func BenchmarkRecursive(t *testing.B) {
	for i := 0; i < t.N; i++ {
		recursive("COMPUTER", "SCIENCE")
	}

}

func BenchmarkBacktracking(t *testing.B) {
	for i := 0; i < t.N; i++ {
		backtracking("COMPUTER", "SCIENCE")
	}

}
