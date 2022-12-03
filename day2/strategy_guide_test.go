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
	score, err := calculateScorePart1(in)

	require.NoError(t, err)

	require.Equal(t, 15, score)
}

func TestCalculateScore2(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(exampleData)
	score, err := calculateScorePart2(in)

	require.NoError(t, err)

	require.Equal(t, 12, score)
}
