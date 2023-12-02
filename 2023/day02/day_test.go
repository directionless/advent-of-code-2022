package day02

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/2023/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 8
	exampleAnswer2 = 2286

	realAnswer1 = 2416
	realAnswer2 = 63307
)

func Test(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name            string
		input           string
		part1           any
		part2           any
		p2inputOverride string
	}{
		{"Example", "example.txt", exampleAnswer1, exampleAnswer2, ""},
		{"Real", "input.txt", realAnswer1, realAnswer2, ""},
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

				day := New(pixelTyp{12, 13, 14})
				require.NoError(t, runner.ScanToHandler(day, in))

				require.Equal(t, tt.part1, day.AnswerPart1())
			})

			t.Run("part2", func(t *testing.T) {
				t.Parallel()

				input := tt.input
				if tt.p2inputOverride != "" {
					input = tt.p2inputOverride
				}
				in, err := os.Open(input)

				require.NoError(t, err)
				defer in.Close()

				day := New(pixelTyp{})
				require.NoError(t, runner.ScanToHandler(day, in))

				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}
