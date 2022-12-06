package day6

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	realAnswer1 = 1702
	realAnswer2 = 3559
)

func TestExamples(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		in    io.Reader
		part1 int
		part2 int
	}{
		{name: "ex1", in: strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), part1: 7, part2: 19},
		{name: "ex2", in: strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz"), part1: 5, part2: 23},
		{name: "ex3", in: strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg"), part1: 6, part2: 23},
		{name: "ex4", in: strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), part1: 10, part2: 29},
		{name: "ex5", in: strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), part1: 11, part2: 26},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			day := New()

			require.NoError(t, day.Handle(tt.in))

			t.Run("part1", func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tt.part1, day.AnswerPart1())
			})

			t.Run("part2", func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tt.part2, day.AnswerPart2())
			})
		})
	}
}

func TestReal(t *testing.T) {
	t.Parallel()

	in, err := os.Open("input.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()

	require.NoError(t, day.Handle(in))

	t.Run("part1", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, realAnswer1, day.AnswerPart1())
	})

	t.Run("part2", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, realAnswer2, day.AnswerPart2())
	})

}
