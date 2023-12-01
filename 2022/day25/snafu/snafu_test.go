package snafu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnafu(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		decimal int
		snafu   string
	}{
		{0, "0"},
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{15, "1=0"},
		{20, "1-0"},
		{2022, "1=11-2"},
		{12345, "1-0---0"},
		{314159265, "1121-1110-1=0"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("i%d to snafu", tt.decimal), func(t *testing.T) {
			t.Parallel()

			actual, err := FromInt(tt.decimal)
			require.NoError(t, err)
			require.Equal(t, tt.snafu, actual)
		})

		t.Run(fmt.Sprintf("%s to int", tt.snafu), func(t *testing.T) {
			t.Parallel()

			actual, err := ToInt(tt.snafu)
			require.NoError(t, err)
			require.Equal(t, tt.decimal, actual)
		})

	}
}
