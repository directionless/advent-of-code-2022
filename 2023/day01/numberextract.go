package day01

import (
	"strings"
)

var (
	numbersInt = map[string]any{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	numbersStr = map[string]any{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	numbersCombined = map[string]any{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func mapKeys(m map[string]any, extras ...map[string]any) []string {
	l := len(m)
	for _, e := range extras {
		l += len(e)
	}

	keys := make([]string, 0, l)
	for k := range m {
		keys = append(keys, k)
	}
	for _, e := range extras {
		for k := range e {
			keys = append(keys, k)
		}
	}

	return keys
}

func valFor(k string, m map[string]any, extras ...map[string]any) any {
	if v, ok := m[k]; ok {
		return v
	}

	for _, e := range extras {
		if v, ok := e[k]; ok {
			return v
		}
	}

	return nil
}

func findFirst(s string, m map[string]any) (string, any) {
	firstIndex := -1
	var firstKey string
	var firstVal any

	for k, v := range m {
		i := strings.Index(s, k)
		switch {
		case i < 0:
			// not found
			continue
		case firstIndex < 0:
			firstIndex = i
			firstKey = k
			firstVal = v

		case i < firstIndex:
			firstIndex = i
			firstKey = k
			firstVal = v
		}

		if i == 0 {
			// Can't be anything smaller
			return k, v
		}
	}

	return firstKey, firstVal
}

func findLast(s string, m map[string]any) (string, any) {
	lastIndex := -1
	var lastKey string
	var lastVal any

	for k, v := range m {
		i := strings.LastIndex(s, k)
		switch {
		case i < 0:
			// not found
			continue
		case lastIndex < 0:
			lastIndex = i
			lastKey = k
			lastVal = v
		case i > lastIndex:
			lastIndex = i
			lastKey = k
			lastVal = v
		}

		if i == len(s)-1 {
			// Can't be anything larger
			return k, v
		}
	}

	return lastKey, lastVal
}
