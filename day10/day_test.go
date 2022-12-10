package day10

import (
	"os"
	"strings"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 13140
	exampleAnswer2 = 0

	realAnswer1 = 13440
	realAnswer2 = 0 //PBZGRAZA

	crtExample = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`

	crtReal = `###..###..####..##..###...##..####..##..
#..#.#..#....#.#..#.#..#.#..#....#.#..#.
#..#.###....#..#....#..#.#..#...#..#..#.
###..#..#..#...#.##.###..####..#...####.
#....#..#.#....#..#.#.#..#..#.#....#..#.
#....###..####..###.#..#.#..#.####.#..#.
`
)

func Test(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		input string
		part1 int
		part2 string
	}{
		{"Example", "example.txt", exampleAnswer1, crtExample},
		{"Real", "input.txt", realAnswer1, crtReal},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			in, err := os.Open(tt.input)
			require.NoError(t, err)
			defer in.Close()

			day := New()
			require.NoError(t, runner.ScanToHandler(day, in))

			require.Equal(t, tt.part1, day.AnswerPart1())

			var crt strings.Builder
			day.GetCRT(&crt)
			require.Equal(t, tt.part2, crt.String())
		})
	}
}
