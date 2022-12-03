package main

import (
	"fmt"
	"os"

	"github.com/directionless/advent-of-code-2022/day2"
)

// At least as of day2, there's a lot of basic open file, parse, etc.
// So I'm going to _slightly_ abstract this. Go makes it a bit hard. But :shrug:
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./main.go <day> <part> <input_file>")
		os.Exit(1)
	}

	day := os.Args[1]
	part := os.Args[2]
	inputFile := os.Args[3]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	switch {
	case day == "1" && part == "1":
	case day == "1" && part == "2":
	case day == "2" && part == "1":
		score, err := day2.CalculateScorePart1(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(score)

	case day == "2" && part == "2":
		score, err := day2.CalculateScorePart2(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(score)

	default:
		fmt.Println("Unknown day/part combination")
		os.Exit(1)
	}
}
