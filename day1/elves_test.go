package day1

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleData = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestElfDivider(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in       io.Reader
		expected []int
	}{
		{
			in:       strings.NewReader(exampleData),
			expected: []int{6000, 4000, 11000, 24000, 10000},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			got, err := elfDivider(tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}
