package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valueFromHand(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in        string
		expected1 handValue
		expected2 handValue
	}{
		{"23456", highCard, highCard},
		{"23345", onePair, onePair},
		{"22TT4", twoPair, twoPair},
		{"JJJ34", threeOfAKind, fourOfAKind},
		{"22QQQ", fullHouse, fullHouse},
		{"AAAAK", fourOfAKind, fourOfAKind},
		{"KKKKK", fiveOfAKind, fiveOfAKind},
		{"JJJ22", fullHouse, fiveOfAKind},
		{"22J44", twoPair, fullHouse},
		{"AAJJJ", fullHouse, fiveOfAKind},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			actual1 := valueFromHand([]byte(tt.in), false)
			assert.Equal(t, tt.expected1, actual1)

			actual2 := valueFromHand([]byte(tt.in), true)
			assert.Equal(t, tt.expected2, actual2)

		})
	}

}
