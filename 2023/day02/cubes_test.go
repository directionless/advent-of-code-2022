package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PixelFromString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		in string
		r  int
		g  int
		b  int
	}{

		{"3 blue, 4 red", 4, 0, 3},
		{"1 red, 2 green, 6 blue", 1, 2, 6},
		{"2 green", 0, 2, 0},
		{"1 blue, 2 green", 0, 2, 1},
		{"3 green, 4 blue, 1 red", 1, 3, 4},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			pixel, err := PixelFromString(tt.in)
			require.NoError(t, err)

			assert.Equal(t, tt.r, pixel.Red(), "red")
			assert.Equal(t, tt.g, pixel.Green(), "green")
			assert.Equal(t, tt.b, pixel.Blue(), "blue")

		})
	}

}
