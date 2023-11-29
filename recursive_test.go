package lcs_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecursive(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "should return empty string",
			a:    "",
			b:    "",
			want: "",
		},
		{
			name: "should return empty string",
			a:    "A",
			b:    "",
			want: "",
		},
		{
			name: "should return empty string",
			a:    "",
			b:    "A",
			want: "",
		},
		{
			name: "should return CE",
			a:    "COMPUTER",
			b:    "SCIENCE",
			want: "CE",
		},
		{
			name: "should return ACE",
			a:    "ABCDE",
			b:    "ACE",
			want: "ACE",
		},
		{
			name: "should return ABC",
			a:    "ABC",
			b:    "ABC",
			want: "ABC",
		},
		{
			name: "should return FSH",
			a:    "FOSH",
			b:    "FISH",
			want: "FSH",
		},
		{
			name: "should return FO",
			a:    "FOSH",
			b:    "FORT",
			want: "FO",
		},
		{
			name: "should return A CE",
			a:    "AB CE",
			b:    "A CE",
			want: "A CE",
		},
		{
			name: "should return MU",
			a:    "ILMU",
			b:    "KOMPUTER",
			want: "MU",
		},
		{
			name: "should return MU",
			a:    "ILMUO",
			b:    "KOMPUTER",
			want: "MU",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := recursive(test.a, test.b)
			t.Logf("LCS of %s and %s is %s\n", test.a, test.b, res)
			assert.Equal(t, test.want, res)
		})
	}
}
