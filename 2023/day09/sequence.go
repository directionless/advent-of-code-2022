package day09

import (
	"errors"
	"fmt"
	"strings"

	"github.com/directionless/advent-of-code-2022/2023/util/extract"
)

type sequence struct {
	nums     []int
	depthMap []int
}

func sequenceFromLine(line []byte) (*sequence, error) {
	foundNums := extract.NumbersFromLineWithNegatives(line)
	nums := make([]int, len(foundNums))

	for i, f := range foundNums {
		nums[i] = f.Int
	}

	s := &sequence{
		nums: nums,
	}

	return s, nil
}

func (s *sequence) Solve() error {
	dm, err := findDifference(s.nums, 0)
	if err != nil {
		return fmt.Errorf("finding differences in %v: %w", s.nums, err)
	}

	s.depthMap = append(dm, s.nums[len(s.nums)-1])
	return nil
}

func (s *sequence) FindNext(n int) (int, error) {
	dm := make([]int, len(s.depthMap))
	copy(dm, s.depthMap)

	numStrs := make([]string, len(s.nums))
	for i, n := range s.nums {
		numStrs[i] = fmt.Sprintf("%d", n)
	}

	fmt.Printf("Finding: %s N=%d\n", strings.Join(numStrs, ", "), n)

	for i := 0; i < n; i++ {
		carry := 0
		for d, val := range dm {
			//fmt.Printf("  Iter: %d: At depth %d, lastval: %d\n", i, d, val)
			dm[d] = val + carry
			carry += val
		}
	}

	fmt.Printf("  Now we have %v\n", dm)

	return dm[len(dm)-1], nil
}

func findDifference(nums []int, depth int) ([]int, error) {
	if depth > 100 {
		return nil, errors.New("too deep")
	}

	fmt.Printf("Solving: %v\n", nums)

	diffs := make([]int, len(nums)-1)

	zeros := true
	for i := 0; i < (len(nums) - 1); i++ {
		d := nums[i+1] - nums[i]
		diffs[i] = d
		if d != 0 {
			zeros = false
		}
	}

	depthMap := make([]int, 0)
	if len(diffs) > 0 {
		depthMap = append(depthMap, diffs[len(diffs)-1])
	}

	if !zeros {
		dm, err := findDifference(diffs, depth+1)
		if err != nil {
			return nil, err
		}

		depthMap = append(dm, depthMap...)
	}

	return depthMap, nil
}
