package day03

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

var (
	spaceChars  = []rune{'.'}
	symbolChars = []rune{'#', '$', '%', '&', '*', '+', '-', '/', '=', '@'}
	numRE       = regexp.MustCompile(`([0-9]+)`)
)

type lineInfo struct {
	line            string
	symbolPositions []int
	numbers         []int
	numberIndexes   [][2]int
}

var emptyLine = &lineInfo{}

func parseLine(line string) (*lineInfo, error) {
	if line == "" || line == "\n" {
		return emptyLine, nil
	}

	li := &lineInfo{
		line:            line,
		symbolPositions: []int{},
		numbers:         []int{},
	}

	for i, r := range line {
		switch {
		case slices.Contains(spaceChars, r):
			continue
		case slices.Contains(symbolChars, r):
			li.symbolPositions = append(li.symbolPositions, i)
		case '0' <= r && r <= '9':
			// We've found a number. Skip it, because we're going to regex it later. (lazy!)
			continue
		default:
			return emptyLine, fmt.Errorf(`unknown character "%s" on line "%s"`, string(r), line)
		}
	}

	// using regex to extract the numbers feels lazy, I ought be able to do it in the line iteration
	// above. But, well, this is good enough.
	numIndexes := numRE.FindAllStringIndex(line, -1)
	li.numberIndexes = make([][2]int, len(numIndexes))
	for i, idx := range numIndexes {
		li.numberIndexes[i] = [2]int{idx[0], idx[1]}
	}

	nums := numRE.FindAllString(line, -1)
	if len(nums) != len(numIndexes) {
		return emptyLine, fmt.Errorf(`mismatch in RE parsing %d indexes, but %d numbers. Line "%s"`, len(numIndexes), len(nums), line)
	}

	li.numbers = make([]int, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return emptyLine, fmt.Errorf(`unable to parse "%s" into number from line "%s"`, num, line)
		}
		li.numbers[i] = n
	}

	return li, nil
}
