package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleData = `A Y
B X
C Z`

func TestCalculateScore(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(exampleData)
	score, err := calculateScore(in)

	require.NoError(t, err)

	require.Equal(t, 15, score)

}
