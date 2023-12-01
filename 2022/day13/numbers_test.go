package day13

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseNumber(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in       string
		expected listNumber
	}{
		{
			in: "[[1],4]",
		},
		{
			in: "[[]]",
		},
		{
			in: "[9]",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			actual, err := ParseNumber([]byte(tt.in))
			require.NoError(t, err)
			_ = actual
			//require.Equal(t, tt.expected, actual)
		})
	}
}

func TestCompareNumbers(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		left     string
		right    string
		expected int
	}{
		{
			left:     "[1]",
			right:    "[1]",
			expected: 0,
		},
		{
			left:     "[1]",
			right:    "[2]",
			expected: -1,
		},
		{
			left:     "[2]",
			right:    "[1]",
			expected: 1,
		},

		{
			left:     "[1,1,3,1,1]",
			right:    "[1,1,5,1,1]",
			expected: -1,
		},

		{
			left:     "[1,1]",
			right:    "[1,1,1]",
			expected: -1,
		},

		{
			left:     "[1,5]",
			right:    "[1,1,1]",
			expected: 1,
		},
		{
			left:     "[[1],[2,3,4]]",
			right:    "[[1],4]",
			expected: -1,
		},
		{
			left:     "[[1],[[4]]]",
			right:    "[[1],[[1],2]]",
			expected: 1,
		},
		{
			left:     "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			right:    "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			expected: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%s vs %s", tt.left, tt.right), func(t *testing.T) {
			t.Parallel()

			left, err := ParseNumber([]byte(tt.left))
			require.NoError(t, err)

			right, err := ParseNumber([]byte(tt.right))
			require.NoError(t, err)

			actual, err := CompareNumbers(left, right)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
