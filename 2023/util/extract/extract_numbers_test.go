package extract

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NumbersFromLine(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input   string
		numbers []int
		indexes [][2]int
	}{
		{"0", []int{0}, [][2]int{{0, 1}}},
		{".1.", []int{1}, [][2]int{{1, 2}}},
		{".1.34", []int{1, 34}, [][2]int{{1, 2}, {3, 5}}},
		{"-1-34", []int{1, 34}, [][2]int{{1, 2}, {3, 5}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			require.Equal(t, len(tt.numbers), len(tt.indexes), "Invalid Test")

			found := NumbersFromLine([]byte(tt.input))

			require.Equal(t, len(tt.numbers), len(found), "Len mismatch")

			for i, n := range found {
				assert.Equal(t, fmt.Sprintf("%d", tt.numbers[i]), n.Str, "string")
				assert.Equal(t, tt.numbers[i], n.Int, "number")
				assert.Equal(t, tt.indexes[i][0], n.StartIdx, "start index")
				assert.Equal(t, tt.indexes[i][1], n.EndIdx, "end index")
			}
		})
	}
}

func Test_NegativeNumbers(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input     string
		numbers   []int
		negatives []int
	}{
		{"0", []int{0}, []int{0}},
		{".1.", []int{1}, []int{1}},
		{".1.34", []int{1, 34}, []int{1, 34}},
		{"-1  34", []int{1, 34}, []int{-1, 34}},
		{"-0  -3..-6", []int{0, 3, 6}, []int{0, -3, -6}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			require.Equal(t, len(tt.numbers), len(tt.negatives), "Invalid Test")

			found := NumbersFromLine([]byte(tt.input))
			foundNegatives := NumbersFromLineWithNegatives([]byte(tt.input))

			require.Equal(t, len(tt.numbers), len(found), "Len mismatch")
			require.Equal(t, len(tt.numbers), len(foundNegatives), "Len mismatch")

			for i, n := range found {
				assert.Equal(t, fmt.Sprintf("%d", tt.numbers[i]), n.Str, "string")
				assert.Equal(t, tt.numbers[i], n.Int, "number")
				assert.Equal(t, tt.negatives[i], foundNegatives[i].Int, "negative")
			}
		})
	}
}
