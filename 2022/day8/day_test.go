package day8

import (
	"fmt"
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 21
	exampleAnswer2 = 8

	realAnswer1 = 1711
	realAnswer2 = 301392
)

func TestLook(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		rowIdx    int
		colIdx    int
		direction directionType
		expected  int
	}{
		// row, col
		{0, 0, right, 2},
		{0, 0, left, 0},
		{0, 0, down, 2},
		{0, 0, up, 0},

		// From the example text
		{1, 2, up, 1},
		{1, 2, left, 1},
		{1, 2, right, 2},
		{1, 2, down, 2},

		// Example 2
		{3, 2, up, 2},
		{3, 2, left, 2},
		{3, 2, down, 1},
		{3, 2, right, 2},
	}

	in, err := os.Open("example.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))
	require.NoError(t, day.Check())

	for _, tt := range tests {
		tt := tt

		name := fmt.Sprintf("%d,%d,%s", tt.rowIdx, tt.colIdx, tt.direction)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := day.Look(tt.rowIdx, tt.colIdx, tt.direction)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExample1(t *testing.T) {
	t.Parallel()

	in, err := os.Open("example.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))
	require.NoError(t, day.Check())

	assert.Equal(t, 4, day.ScenicScore(1, 2))
	assert.Equal(t, 8, day.ScenicScore(3, 2))

	assert.Equal(t, exampleAnswer1, day.AnswerPart1())
	assert.Equal(t, exampleAnswer2, day.AnswerPart2())
}

func TestReal(t *testing.T) {
	t.Parallel()

	in, err := os.Open("input.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))
	require.NoError(t, day.Check())

	//day.Dump()

	assert.Equal(t, realAnswer1, day.AnswerPart1())
	assert.Equal(t, realAnswer2, day.AnswerPart2())
}
