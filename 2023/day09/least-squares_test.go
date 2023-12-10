package day09

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leastSquaresMethod(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		yVals []int
		a     float64
		b     float64
		next  int
	}{
		{
			yVals: []int{0, 3, 6, 9, 12, 15},
			next:  18,
		},
		{
			yVals: []int{1, 3, 6, 10, 15, 21},
			next:  28,
		},
		{
			yVals: []int{10, 13, 16, 21, 30, 45},
			next:  68,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			points := make([]Point, len(tt.yVals))
			for i, y := range tt.yVals {
				points[i] = Point{float64(i), float64(y)}
			}
			a, b := leastSquaresMethod(&points)

			nextX := len(tt.yVals)
			next := a*float64(nextX) + b
			fmt.Println(next)

			assert.Equal(t, tt.a, a, "a")
			assert.Equal(t, tt.b, b, "b")

		})
	}
}
