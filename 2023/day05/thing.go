package day05

import (
	"fmt"

	"github.com/directionless/advent-of-code-2022/2023/util/extract"
)

// the destination range start, the source range start, and the range length.

type thingThatNeedsMap struct {
	class    string
	dstStart int
	srcStart int
	length   int
}

func thingFromLine(class string, line []byte) (thingThatNeedsMap, error) {
	nums := extract.NumbersFromLine(line)
	if len(nums) != 3 {
		return thingThatNeedsMap{}, fmt.Errorf("expected 3 numbers, got %d", len(nums))
	}

	return thingThatNeedsMap{
		class:    class,
		dstStart: nums[0].Int,
		srcStart: nums[1].Int,
		length:   nums[2].Int,
	}, nil
}

func (t thingThatNeedsMap) String() string {
	return fmt.Sprintf("%s-%d(%d:%d)", t.class, t.dstStart, t.srcStart, t.srcStart+t.length-1)
}

func (t thingThatNeedsMap) Contains(n int) int {
	//fmt.Printf("%s: %d <= %d && %d < %d\n", t, t.srcStart, n, n, t.srcStart+t.length)
	if t.srcStart <= n && n < t.srcStart+t.length {
		return t.dstStart + (n - t.srcStart)
	}

	return -1
}

func (t thingThatNeedsMap) DestContains(n int) bool {
	return t.dstStart <= n && n < t.dstStart+t.length
}

type bySrcStart []thingThatNeedsMap

// Len is part of sort.Interface.
func (s bySrcStart) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s bySrcStart) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s bySrcStart) Less(i, j int) bool {
	return s[i].srcStart < s[j].srcStart
}
