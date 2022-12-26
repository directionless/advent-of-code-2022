package day25

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1      = 4890
	exampleAnswer1Snafu = "2=-1=0"
	exampleAnswer2      = 0

	realAnswer1      = 37512839082437
	realAnswer1Snafu = "20-1-11==0-=0112-222"
	realAnswer2      = 0
)

func Test(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name   string
		input  string
		part1  int
		part1S string
		part2  int
	}{
		{"Example", "example.txt", exampleAnswer1, exampleAnswer1Snafu, exampleAnswer2},
		{"Real", "input.txt", realAnswer1, realAnswer1Snafu, realAnswer2},
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

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				assert.Equal(t, tt.part1, day.AnswerPart1())
				sn, err := day.AnswerPart1Snafu()
				require.NoError(t, err)
				assert.Equal(t, tt.part1S, sn)
			})

			t.Run("part2", func(t *testing.T) {
				t.Parallel()

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				//.NoError(t, day.Run(10000))
				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}
