package day4

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleData = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	exampleAnswer1 = 2
	exampleAnswer2 = 4

	realAnswer1 = 556
	realAnswer2 = 0
)

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		t.Parallel()

		in := strings.NewReader(exampleData)

		part1 := &Part1Handler{}
		require.NoError(t, runner.ScanToHandler(part1, in))

		t.Run("part1", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, exampleAnswer1, part1.Answer())
		})

		t.Run("part2", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, exampleAnswer2, part1.AnswerPart2())
		})
	})

	t.Run("real", func(t *testing.T) {
		t.Parallel()

		file, err := os.Open("input.txt")
		if errors.Is(err, os.ErrNotExist) {
			t.Skip("no input")
		}
		require.NoError(t, err)
		defer file.Close()

		part1 := &Part1Handler{}
		require.NoError(t, runner.ScanToHandler(part1, file))

		t.Run("part1", func(t *testing.T) {
			require.Equal(t, realAnswer1, part1.Answer())
		})
		t.Run("part2", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, realAnswer2, part1.AnswerPart2())
		})

	})

}

func Testpart2(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		t.Parallel()

		in := strings.NewReader(exampleData)

		part2 := &Part2Handler{}
		require.NoError(t, runner.ScanToHandler(part2, in))
		require.Equal(t, exampleAnswer1, part2.Answer())
	})

	t.Run("real", func(t *testing.T) {
		t.Parallel()

		file, err := os.Open("input.txt")
		if errors.Is(err, os.ErrNotExist) {
			t.Skip("no input")
		}
		require.NoError(t, err)
		defer file.Close()

		part2 := &Part2Handler{}
		require.NoError(t, runner.ScanToHandler(part2, file))
		require.Equal(t, realAnswer1, part2.Answer())
	})

}
