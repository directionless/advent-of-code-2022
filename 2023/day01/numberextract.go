package day01

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	numbersInt = map[string]any{
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

	numbersIntRE = regexp.MustCompile(fmt.Sprintf("(%s)", strings.Join(mapKeys(numbersInt), "|")))
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

func findFirst(s string, m map[string]any, extras ...map[string]any) (string, any) {
	re := regexp.MustCompile(fmt.Sprintf("(%s)", strings.Join(mapKeys(m, extras...), "|")))

	// Only need the first match
	if matches := re.FindAllString(s, 1); len(matches) > 0 {
		return matches[0], valFor(matches[0], m, extras...)
	}

	return "", nil
}

func findLast(s string, m map[string]any, extras ...map[string]any) (string, any) {

	expression := reverse(strings.Join(mapKeys(m, extras...), "|"))

	re := regexp.MustCompile(fmt.Sprintf("(%s)", expression))

	// Only need the first match
	if matches := re.FindAllString(reverse(s), 1); len(matches) > 0 {
		s := reverse(matches[0])
		return s, valFor(s, m, extras...)
	}

	return "", nil
}

// reverse reverses a string. From https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang
func reverse(s string) string {
	rns := []rune(s)
	// swap the letters of the string,
	// like first with last and so on.
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
