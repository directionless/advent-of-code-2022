package runner

import (
	"bufio"
	"fmt"
	"io"
)

type challengerHandler interface {
	Consume([]byte) error
	AnswerPart1() int
	AnswerPart2() int
	Print()
}

func Run(handler challengerHandler, rd io.Reader) error {
	if err := ScanToHandler(handler, rd); err != nil {
		return err
	}

	handler.Print()

	return nil
}

func ScanToHandler(handler challengerHandler, rd io.Reader) error {
	scanner := bufio.NewScanner(rd)
	lineNum := 0
	for scanner.Scan() {
		lineNum += 1

		// scanner.Bytes() would end up being a pointer, so we need to copy
		// it. else, we run into weird reuse errors.
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())

		if err := handler.Consume(line); err != nil {
			return fmt.Errorf("error handling line number %d: %w", lineNum, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}

	return nil
}
