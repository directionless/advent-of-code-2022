package day7

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 95437
	exampleAnswer2 = 24933642

	// 1577489 is too high
	realAnswer1 = 1517599
	realAnswer2 = 2481982
)

func TestExample1(t *testing.T) {
	t.Parallel()

	in, err := os.Open("example.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))

	day.DumpTree()

	require.Equal(t, exampleAnswer1, day.AnswerPart1())
	require.Equal(t, exampleAnswer2, day.AnswerPart2())

	var tests = []struct {
		path string
		size int
	}{
		{path: "/a/e", size: 584},
		{path: "/a", size: 94853},
		{path: "/d", size: 24933642},
		{path: "/", size: 48381165},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.path, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.size, day.DirSize(tt.path))
		})
	}
}

func TestExample2(t *testing.T) {
	t.Parallel()

	in, err := os.Open("example2.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))

	day.DumpTree()

	require.Equal(t, 33, day.AnswerPart1())
}

func TestReal(t *testing.T) {
	t.Parallel()

	in, err := os.Open("input.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))

	day.DumpTree()

	require.Equal(t, realAnswer1, day.AnswerPart1())
	require.Equal(t, realAnswer2, day.AnswerPart2())
}
