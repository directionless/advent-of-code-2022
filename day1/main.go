package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: day1 <input>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	elves, err := elfDivider(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n := 3
	sum := 0
	for _, c := range topCalories(n, elves) {
		sum += c
	}
	fmt.Printf("Top %d elves are carrying %d\n", n, sum)
}

func topCalories(n int, elves []int) []int {
	// This only works for small values of N. Too large and the perf is going to be horrible.
	biggest := make([]int, n+1)

	for _, calories := range elves {
		biggest[0] = calories
		sort.Sort(sort.IntSlice(biggest))
	}

	return biggest[1:]
}

func elfDivider(rd io.Reader) ([]int, error) {
	elves := []int{}

	scanner := bufio.NewScanner(rd)

	carrying := 0
	for scanner.Scan() {
		line := scanner.Text()

		// New Elf
		if line == "" {
			elves = append(elves, carrying)
			carrying = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", line, err)
		}

		carrying += num
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	// Remember the last elf!
	elves = append(elves, carrying)

	return elves, nil
}
