package day01

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findFirstLast(t *testing.T) {
	var tests = []struct {
		in           string
		firstK       string
		firstV       int
		lastK        string
		lastV        int
		includeWords bool
	}{
		{
			in:           "eightwothree",
			firstK:       "eight",
			firstV:       8,
			lastK:        "three",
			lastV:        3,
			includeWords: true,
		},
		{
			in:           "oneightwo",
			firstK:       "one",
			firstV:       1,
			lastK:        "two",
			lastV:        2,
			includeWords: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			m := numbersInt
			if tt.includeWords {
				m = numbersCombined
			}

			{
				firstK, firstV := findFirst(tt.in, m)
				require.Equal(t, tt.firstK, firstK, "first key")
				require.Equal(t, tt.firstV, firstV, "first value")
			}

			{
				lastK, lastV := findLast(tt.in, m)
				require.Equal(t, tt.lastK, lastK, "last key")
				require.Equal(t, tt.lastV, lastV, "last value")
			}

		})
	}
}

func Benchmark_findFirst(b *testing.B) {
	// load data
	data, err := os.ReadFile("input.txt")
	require.NoError(b, err)
	lines := strings.Split(string(data), "\n")

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		line := lines[n%len(lines)]
		findFirst(line, numbersCombined)
	}
}
