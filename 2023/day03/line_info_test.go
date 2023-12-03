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

func Test_parseLineTouching(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		line      string
		symbolIdx []int
		expected  []int
	}{
		{line: "", symbolIdx: []int{5}},
		{line: "..........", symbolIdx: []int{5}},
		{line: "467.114...", symbolIdx: []int{0}, expected: []int{467}},
		{line: "467.114...", symbolIdx: []int{3}, expected: []int{467, 114}},
		{line: "467.114...", symbolIdx: []int{7}, expected: []int{114}},
		{line: "467.114...", symbolIdx: []int{10}},

		{line: "..2..", symbolIdx: []int{0}, expected: []int{}},
		{line: "..2..", symbolIdx: []int{1}, expected: []int{2}},
		{line: "..2..", symbolIdx: []int{2}, expected: []int{2}},
		{line: "..2..", symbolIdx: []int{3}, expected: []int{2}},
		{line: "..2..", symbolIdx: []int{4}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			actual, err := parseLine(tt.line)
			require.NoError(t, err)

			nums := actual.NumbersTouching(tt.symbolIdx)

			if tt.expected == nil {
				require.Empty(t, nums)
			} else {
				require.Equal(t, tt.expected, nums)
			}
		})
	}
}
