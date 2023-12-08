package day07

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_valueFromHand(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in       string
		expected handValue
	}{
		{"23456", highCard},
		{"23345", onePair},
		{"22TT4", twoPair},
		{"JJJ34", threeOfAKind},
		{"22QQQ", fullHouse},
		{"AAAAK", fourOfAKind},
		{"KKKKK", fiveOfAKind},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			actual := valueFromHand([]byte(tt.in))
			require.Equal(t, tt.expected, actual)
		})
	}

}
