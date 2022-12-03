package main

import (
	"fmt"
	"os"
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

	score, err := calculateScorePart2(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(score)

}
