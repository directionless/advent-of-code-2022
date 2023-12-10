package day09

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sequence(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in   string
		next int
	}{
		{"0 3 6 9 12 15", 18},
		{"1 3 6 10 15 21", 28},
		{"10 13 16 21 30 45", 68},
		{"1 16 46 89 145 214 296 391 499 620 754 901 1061 1234 1420 1619 1831 2056 2294 2545 2809", 3088},
		{"6 6 6 6 6 6 6 6 6 6 6 6 6 6 6 6 6 20 300 3242 24839", 149169},
		{"-5 -10 -15 -2 67 250 630 1350 2723 5496 11405 24252 51884 109706 226850 457196 899887 1739416 3326840 6346565 12156399", 23466486},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()

			s, err := sequenceFromLine([]byte(tt.in))
			require.NoError(t, err)

			require.NoError(t, s.Solve())

			n, err := s.FindNext(1)
			require.NoError(t, err)
			require.Equal(t, tt.next, n)
		})
	}

}
