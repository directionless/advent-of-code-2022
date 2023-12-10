package extract

import (
	"strconv"
)

const (
	ascii0   = 48
	ascii9   = 57
	asciiNeg = 45
)

type foundNumber struct {
	Int      int
	Str      string
	StartIdx int
	EndIdx   int
}

// NumbersFromLine extracts the ascii numbers from a line. It ignores all non-number characters.
// It should produce identical results as two regex calls, but faster.
//
//	re := regexp.MustCompile(`(\d+)`)
//	fmt.Printf("%v\n", re.FindAllString(s, -1))
//	fmt.Printf("%v\n", re.FindAllStringIndex(s, -1))
func NumbersFromLine(line []byte) []foundNumber {
	return findNumbers(line, false)
}

func NumbersFromLineWithNegatives(line []byte) []foundNumber {
	return findNumbers(line, true)
}

func findNumbers(line []byte, includeNegatives bool) []foundNumber {
	if line == nil || len(line) == 0 {
		return nil
	}

	numbers := make([]foundNumber, 0)

	cur := foundNumber{}
	found := false
	for i, b := range line {
		if ascii0 <= b && b <= ascii9 || (includeNegatives && b == asciiNeg) {
			cur.Str += string(b)

			if !found {
				found = true
				cur.StartIdx = i
			}
		} else {
			if found {
				found = false
				cur.EndIdx = i
			}

			if cur.EndIdx != 0 {
				cur.Int = mustAtoi(cur.Str)
				numbers = append(numbers, cur)
				cur = foundNumber{}
			}
		}
	}

	// handle a number at the end of the string
	if found {
		cur.EndIdx = len(line)
		cur.Int = mustAtoi(cur.Str)
		numbers = append(numbers, cur)
	}

	//spew.Dump(numbers)
	return numbers
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
