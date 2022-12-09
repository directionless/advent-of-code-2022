package day9

import (
	"os"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleAnswer1 = 13
	exampleAnswer2 = 0

	realAnswer1 = 5878
	realAnswer2 = 0
)

func TestExample1(t *testing.T) {
	t.Parallel()

	in, err := os.Open("example.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))

	require.Equal(t, exampleAnswer1, day.AnswerPart1())
	require.Equal(t, exampleAnswer2, day.AnswerPart2())

}

func TestReal(t *testing.T) {
	t.Parallel()

	in, err := os.Open("input.txt")
	require.NoError(t, err)
	defer in.Close()

	day := New()
	require.NoError(t, runner.ScanToHandler(day, in))

	require.Equal(t, realAnswer1, day.AnswerPart1())
	require.Equal(t, realAnswer2, day.AnswerPart2())
}
