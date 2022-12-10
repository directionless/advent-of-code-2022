package crt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPosFromCycle(t *testing.T) {
	t.Parallel()

	// pos is [2]int{c, r}

	var tests = []struct {
		cycle    int
		expected [2]int
	}{
		{1, [2]int{0, 0}},
		{2, [2]int{1, 0}},
		{40, [2]int{39, 0}},
		{41, [2]int{0, 1}},
		{80, [2]int{39, 1}},
		{81, [2]int{0, 2}},
		{120, [2]int{39, 2}},
		{121, [2]int{0, 3}},
		{160, [2]int{39, 3}},
		{161, [2]int{0, 4}},
		{200, [2]int{39, 4}},
		{201, [2]int{0, 5}},
		{240, [2]int{39, 5}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.cycle), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.expected, posFromCycle(tt.cycle))
		})
	}
}
