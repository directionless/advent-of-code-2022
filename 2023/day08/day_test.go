package day08

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/2023/runner"
	"github.com/stretchr/testify/require"
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
		{"Example", "example.txt", ExampleAnswer1, ExampleAnswer2, "example2.txt"},
		{"Real", "input.txt", RealAnswer1, RealAnswer2, ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			in, err := os.Open(tt.input)
			require.NoError(t, err)
			defer in.Close()

			// Ingest and solve
			day := New()
			require.NoError(t, runner.ScanToHandler(day, in))
			require.NoError(t, day.Solve())

			t.Run("part1", func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tt.part1, day.AnswerPart1())
			})

			t.Run("part2", func(t *testing.T) {
				t.Parallel()

				// part2 _occasionally_ needs a new input file. In that case, regenerate
				// `day`. Otherwise reuse it
				if tt.p2inputOverride != "" {
					in, err := os.Open(tt.p2inputOverride)
					require.NoError(t, err)
					defer in.Close()

					day = New()
					require.NoError(t, runner.ScanToHandler(day, in))
					require.NoError(t, day.Solve())
				}

				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}
