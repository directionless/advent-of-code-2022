package day01

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/2023/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 142
	exampleAnswer2 = 281

	realAnswer1 = 54968
	realAnswer2 = 0 // 54110 is too high
)

func Test_wordToNum(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"two1nine", "219"},
		{"eightwothree", "8wo3"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			linep2 := findNumberWords.ReplaceAllFunc([]byte(tt.input), wordToNum)

			require.Equal(t, []byte(tt.expected), linep2)
		})
	}

}

func Test(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name            string
		input           string
		part1           any
		part2           any
		p2inputOverride string
	}{
		{"Example", "example.txt", exampleAnswer1, exampleAnswer2, "example2.txt"},
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

				day := New()
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

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				//.NoError(t, day.Run(10000))
				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}
