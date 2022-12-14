package day12

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 31
	exampleAnswer2 = 0

	realAnswer1 = 440 // 320 is too low
	realAnswer2 = 0
)

func Test(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		input string
		part1 int
		part2 int
	}{
		{"Example", "example.txt", exampleAnswer1, exampleAnswer2},
		{"Real", "input.txt", realAnswer1, realAnswer2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Run("Grid", func(t *testing.T) {
				t.Parallel()

				if tt.name != "Example" {
					t.Skip()
				}

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				grid := day.Grid()

				t.Run("Neighbors", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, 2, len(grid.Neighbors(grid.Node(0, 0))))
					assert.Equal(t, 3, len(grid.Neighbors(grid.Node(2, 2))))
					assert.Equal(t, 4, len(grid.Neighbors(grid.Node(3, 3))))
				})

			})

			t.Run("Part1", func(t *testing.T) {
				t.Parallel()

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				require.Equal(t, tt.part1, day.AnswerPart1())
			})

			t.Run("Part2", func(t *testing.T) {
				t.Parallel()

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}
