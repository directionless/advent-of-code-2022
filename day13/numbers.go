package day13

import (
	"encoding/json"
	"fmt"
)

type listNumber []any

func ParseNumber(raw []byte) (listNumber, error) {
	var num listNumber
	if err := json.Unmarshal(raw, &num); err != nil {
		return nil, fmt.Errorf("unable to parse %s: %w", raw, err)
	}

	return num, nil
}

// Compare returns -1 if left is lower, 0 if the same, and 1 if right is lower. Rules:
//
// If both values are integers, the lower integer should come first. If the left integer is lower than the right integer, the inputs are in the right order. If the left integer is higher than the right integer, the inputs are not in the right order. Otherwise, the inputs are the same integer; continue checking the next part of the input.
//
// If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
//
// If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].

func CompareNumbers(left, right listNumber) (int, error) {
	if left == nil && right == nil {
		return 0, nil
	}

	for i := 0; i < len(left); i++ {
		if i > len(right)-1 {
			// left is bigger
			return 1, nil
		}

		leftVal := left[i]
		rightVal := right[i]

		switch leftVal := leftVal.(type) {
		case float64:
			switch rightVal := rightVal.(type) {
			case float64:
				if leftVal < rightVal {
					return -1, nil
				}
				if leftVal > rightVal {
					return 1, nil
				}
			case []any:
				ret, err := CompareNumbers(listNumber{leftVal}, rightVal)
				if err != nil {
					return 0, err
				}
				if ret != 0 {
					return ret, nil
				}
			default:
				return 0, fmt.Errorf("unknown right type: %T", rightVal)
			}
		case []any:
			switch rightVal := rightVal.(type) {
			case float64:
				ret, err := CompareNumbers(leftVal, listNumber{rightVal})
				if err != nil {
					return 0, err
				}
				if ret != 0 {
					return ret, nil
				}

			case []any:
				ret, err := CompareNumbers(leftVal, rightVal)
				if err != nil {
					return 0, err
				}
				if ret != 0 {
					return ret, nil
				}

			default:
				return 0, fmt.Errorf("unknown right type: %T", rightVal)
			}
		default:
			return 0, fmt.Errorf("unknown left type: %T", leftVal)
		}
	}

	// Check to see if right is longer, if we got here, it's bigger.
	if len(right) > len(left) {
		return -1, nil
	}

	return 0, nil

}
