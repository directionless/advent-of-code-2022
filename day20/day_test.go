package day20

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 3
	exampleAnswer2 = 0

	realAnswer1 = -1 // 13343 too high
	realAnswer2 = 0
)

func (h *dayHandler) MoveFirstElementTesting() {
	h.zero = h.FindZero()
	h.MoveElement(h.ring.Front())
}

func TestExtraExamples(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in   []int
		ring string
	}{
		{

			in:   []int{1, 0, 1},
			ring: "0, 1, 1, ",
		},
		{
			in:   []int{2, 0, 1},
			ring: "0, 1, 2, ",
		},

		{
			in:   []int{3, 0, 1},
			ring: "0, 3, 1, ",
		},

		{
			in:   []int{6, 0, 1},
			ring: "0, 1, 6, ",
		},
		{
			in:   []int{7, 0, 1},
			ring: "0, 7, ",
		},
		{
			in:   []int{-6, 0, 1},
			ring: "-6, 0, ",
		},
		{
			in:   []int{07, 0, 1},
			ring: "0, -7, ",
		},
		{
			in:   []int{-1, 0, 1},
			ring: "0, -1, ",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			var sb strings.Builder
			for _, v := range tt.in {
				fmt.Fprintf(&sb, "%d\n", v)
			}
			day := New()
			require.NoError(t, runner.ScanToHandler(day, strings.NewReader(sb.String())))

			day.MoveFirstElementTesting()

			if tt.ring != "" {
				require.Equal(t, tt.ring, day.StringRing())
			}

		})
	}
}

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

			t.Run("part1", func(t *testing.T) {
				t.Parallel()

				in, err := os.Open(tt.input)
				require.NoError(t, err)
				defer in.Close()

				day := New()
				require.NoError(t, runner.ScanToHandler(day, in))

				require.NoError(t, day.Solve())
				require.Equal(t, tt.part1, day.AnswerPart1())
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
