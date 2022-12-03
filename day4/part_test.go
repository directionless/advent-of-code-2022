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
	exampleData    = ``
	exampleAnswer1 = 0
	exampleAnswer2 = 0

	realAnswer1 = 0
	realAnswer2 = 0
)

func TestPart1(t *testing.T) {
	t.Parallel()

	t.Run("example", func(t *testing.T) {
		t.Parallel()

		in := strings.NewReader(exampleData)

		part1 := &Part1Handler{}
		require.NoError(t, runner.ScanToHandler(part1, in))
		require.Equal(t, exampleAnswer1, part1.Answer())
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
		require.Equal(t, realAnswer1, part1.Answer())
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
