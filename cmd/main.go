package main

import (
	"fmt"
	"os"

	"github.com/directionless/advent-of-code-2022/day1"
	"github.com/directionless/advent-of-code-2022/day2"
)

// At least as of day2, there's a lot of basic open file, parse, etc.
// So I'm going to _slightly_ abstract this. Go makes it a bit hard. But :shrug:

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("Usage: ./main.go <day> <part> [input_file]")
		os.Exit(1)
	}

	day := os.Args[1]
	part := os.Args[2]

	inputFile := fmt.Sprintf("day%s/input.txt", day)
	if len(os.Args) == 4 {
		inputFile = os.Args[3]
	}

	file, err := os.Open(inputFile)
	checkError(err)
	defer file.Close()

	switch {
	case day == "1" && part == "1":
		checkError(day1.Part1(file))
	case day == "1" && part == "2":
		checkError(day1.Part2(file))
	case day == "2" && part == "1":
		checkError(day2.Part1(file))
	case day == "2" && part == "2":
		checkError(day2.Part2(file))
	default:
		fmt.Println("Unknown day/part combination")
		os.Exit(1)
	}
}
