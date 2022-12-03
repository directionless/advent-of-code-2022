package day3

import (
	"bufio"
	"fmt"
	"io"
)

func Part1(rd io.Reader) error {
	tot, err := findTotalPriority(rd)
	if err != nil {
		return err
	}

	fmt.Printf("Total Priorty: %d\n", tot)

	return nil
}

func findTotalPriority(rd io.Reader) (int, error) {
	totalPriority := 0
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		// Split into two strings.
		line := scanner.Bytes()

		miss := findMisFiledInSack(line)
		if miss == 0 {
			return 0, fmt.Errorf("no miss")
		}

		totalPriority += itemPriority(miss)

	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to scan: %w", err)
	}

	return totalPriority, nil
}

func Part2(rd io.Reader) error {
	tot, err := findBadges(rd)
	if err != nil {
		return err
	}

	// 7276 is too high
	// 2521 is too high
	fmt.Printf("Total Badge Priority: %d\n", tot)

	return nil
}

func findBadges(rd io.Reader) (int, error) {
	totalPriority := 0

	// Gotta split into 3s. Quick and dirty style
	num := 0
	sack1 := []byte{}
	sack2 := []byte{}
	sack3 := []byte{}

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {

		// scanner.Bytes() would end up being a pointer, so we need to copy
		// it. else, we run into weird reuse errors.
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())

		// Split into two strings.
		switch num {
		case 0:
			sack1 = line
			num += 1
			continue
		case 1:
			sack2 = line
			num += 1
			continue
		case 2:
			sack3 = line
			num = 0
		}

		badge := findBadge(sack1, sack2, sack3)
		totalPriority += itemPriority(badge)

	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to scan: %w", err)
	}

	return totalPriority, nil

}

func findBadge(sack1, sack2, sack3 []byte) byte {
	seen1 := make(map[byte]bool)
	seen2 := make(map[byte]bool)

	for _, b := range sack1 {
		seen1[b] = true
	}

	for _, b := range sack2 {
		seen2[b] = true
	}

	for _, b := range sack3 {
		if seen1[b] && seen2[b] {
			return b
		}
	}
	return 0
}

func compartmentsFromLine(items []byte) ([]byte, []byte) {
	if len(items)%2 != 0 {
		panic("Odd number of characters in line")
	}

	compartmntSize := len(items) / 2
	return items[0:compartmntSize], items[compartmntSize:]
}

func findMisFiledInCompartments(comp1, comp2 []byte) byte {
	seen := make(map[byte]bool)

	for _, c := range comp1 {
		seen[c] = true
	}

	for _, c := range comp2 {
		if seen[c] {
			return c
		}
	}

	return 0
}

func findMisFiledInSack(rucksack []byte) byte {
	size := len(rucksack) / 2

	seen := make(map[byte]bool)

	for i, b := range rucksack {
		// first compartment. Note what we've seen.
		if i < size {
			seen[b] = true
			continue
		}

		// Second compartment, have we seen it?
		if seen[b] {
			return b
		}
	}

	return 0
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func itemPriority(b byte) int {
	switch {
	// lowercase
	case b >= 97 && b <= 122:
		return int(b) - 96
	// uppercase
	case b >= 65 && b <= 90:
		return int(b) - 64 + 26
	default:
		panic("unknown score")
		return 0
	}
}
