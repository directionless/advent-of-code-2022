package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	elfNum, calories := mostCalories(elves)
	fmt.Printf("Elf %d (0 indexed) has the most calories with %d\n", elfNum, calories)
}

func mostCalories(elves []int) (int, int) {
	var elfNumber, heaviest int

	for i, calories := range elves {
		if calories > heaviest {
			heaviest = calories
			elfNumber = i
		}
	}

	return elfNumber, heaviest
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
