package day15

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	examplePart1Row = 10
	exampleAnswer1  = 26
	exampleAnswer2  = 0

	realPart1Row = 2000000
	realAnswer1  = 4961647
	realAnswer2  = 0
)

func Test(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		input string
		p1row int
		part1 int
		part2 int
	}{
		{"Example", "example.txt", examplePart1Row, exampleAnswer1, exampleAnswer2},
		{"Real", "input.txt", realPart1Row, realAnswer1, realAnswer2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Run("part1", func(t *testing.T) {
				t.Parallel()

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New(tt.p1row)
				require.NoError(t, runner.ScanToHandler(day, in))

				require.Equal(t, tt.part1, day.AnswerPart1())
			})

			t.Run("part2", func(t *testing.T) {
				t.Parallel()

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New(tt.p1row)
				require.NoError(t, runner.ScanToHandler(day, in))

				//.NoError(t, day.Run(10000))
				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}
