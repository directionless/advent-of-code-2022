package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseLine(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		line      string
		symbolPos []int
		numbers   []int
		numIdx    [][2]int
	}{
		{line: ""},
		{line: ".........."},
		{line: "467..114..", numbers: []int{467, 114}, numIdx: [][2]int{{0, 3}, {5, 8}}},
		{line: "...*......", symbolPos: []int{3}},
		{line: "617*......", symbolPos: []int{3}, numbers: []int{617}, numIdx: [][2]int{{0, 3}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()

			actual, err := parseLine(tt.line)
			require.NoError(t, err)
			assert.Equal(t, tt.line, actual.line)

			if tt.symbolPos == nil {
				assert.Empty(t, actual.symbolPositions)
			} else {
				assert.Equal(t, tt.symbolPos, actual.symbolPositions, "symbol positions")
			}

			if tt.numbers == nil {
				assert.Empty(t, actual.numbers)
			} else {
				assert.Equal(t, tt.numbers, actual.numbers, "numbers")
			}

			if tt.numIdx == nil {
				assert.Empty(t, actual.numberIndexes)
			} else {
				assert.Equal(t, tt.numIdx, actual.numberIndexes, "number indexes")
			}

		})
	}
}
