package snafu

import (
	"errors"
	"fmt"
	"strings"
)

// Snafu Numbers
//
// "SNAFU works the same way, except it uses powers of five instead of ten. Starting from the right, you have a ones place, a fives place, a twenty-fives place, a one-hundred-and-twenty-fives place, and so on. It's that easy!"
//
// You ask why some of the digits look like - or = instead of "digits".
//
// "You know, I never did ask the engineers why they did that. Instead of using digits four through zero, the digits are 2, 1, 0, minus (written -), and double-minus (written =). Minus is worth -1, and double-minus is worth -2."

const (
	radix           = 5
	maxExponentSize = 20
)

var segmentMap = map[int]string{
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
}

func FromInt(num int) (string, error) {
	//maxE, err := maxExp(num)
	//if err != nil {
	//	return "", err
	//}

	if num < 0 {
		return "", errors.New("negative numbers not supported")
	}

	if num == 0 {
		return "0", nil
	}

	var sbr strings.Builder

	// Becasue the snafu system uses a funny off-by-two, we do some add/substracts
	// here so the modulus works

	for num > 0 {
		remainder := ((num + 2) % radix) - 2
		segment, ok := segmentMap[remainder]
		if !ok {
			return "", fmt.Errorf("invalid segment %d", remainder)
		}
		sbr.WriteString(segment)

		//fmt.Printf("Looking at %d, remainder: %d, segment %s\n", num, remainder, segment)
		num = (num + 2) / 5
	}

	// Since we build the string backwards, we need to reverse it
	var sb strings.Builder
	for i := sbr.Len() - 1; i >= 0; i-- {
		sb.WriteByte(sbr.String()[i])
	}

	return sb.String(), nil
}

func ToInt(s string) (int, error) {
	pos := 0
	tot := 0
	for i := len(s) - 1; i >= 0; i-- {
		val := intPow(5, pos)
		pos++
		switch s[i] {
		case '2':
			tot += 2 * val
		case '1':
			tot += val
		case '0':
			continue
		case '-':
			tot -= val
		case '=':
			tot -= 2 * val
		default:
			panic(fmt.Sprintf("Unknown character %c", s[i]))
		}
	}
	return tot, nil
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
