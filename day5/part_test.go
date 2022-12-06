package day5

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/directionless/advent-of-code-2022/runner"
	"github.com/stretchr/testify/require"
)

const (
	exampleData = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	exampleAnswer1 = 0 //"CMZ"
	exampleAnswer2 = 0

	realAnswer1 = 0 //"SBPQRSCDF"
	realAnswer2 = 0
)

func TestExamples(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in  string
		out int
	}{
		{in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", out: 7},
		{in: "bvwbjplbgvbhsrlpgdmjqwftvncz", out: 5},
		{in: "nppdvjthqldpwncqszvftbrmjlhg", out: 6},
		{in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", out: 10},
		{in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", out: 11},
	}

	t.Run("example", func(t *testing.T) {
		t.Parallel()

		in := strings.NewReader(exampleData)

		part1 := New()
		require.NoError(t, runner.ScanToHandler(part1, in))
		require.Equal(t, exampleAnswer1, part1.AnswerPart1())

	})

	t.Run("real", func(t *testing.T) {
		t.Parallel()

		file, err := os.Open("input.txt")
		if errors.Is(err, os.ErrNotExist) {
			t.Skip("no input")
		}
		require.NoError(t, err)
		defer file.Close()

		part1 := New()
		require.NoError(t, runner.ScanToHandler(part1, file))
		require.Equal(t, realAnswer1, part1.AnswerPart1())

		require.NoError(t, part1.PrintStacks())
		//spew.Dump(part1)
	})

}

/*
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
*/
