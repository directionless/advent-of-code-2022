package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleData = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestFindTotalPriority(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(exampleData)
	tot, err := findTotalPriority(in)
	require.NoError(t, err)
	require.Equal(t, 157, tot)
}

func TestCompartmentsFromLine(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input        []byte
		compartment1 []byte
		compartment2 []byte
		misfiled     byte
	}{
		{
			input:        []byte("vJrwpWtwJgWrhcsFMMfFFhFp"),
			compartment1: []byte("vJrwpWtwJgWr"),
			compartment2: []byte("hcsFMMfFFhFp"),
			misfiled:     byte('p'),
		},
		{
			input:        []byte("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
			compartment1: []byte("jqHRNqRjqzjGDLGL"),
			compartment2: []byte("rsFMfFZSrLrFZsSL"),
			misfiled:     byte('L'),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			c1, c2 := compartmentsFromLine(tt.input)
			require.Equal(t, tt.compartment1, c1)
			require.Equal(t, tt.compartment2, c2)

			require.Equal(t, tt.misfiled, findMisFiledInCompartments(c1, c2))
			require.Equal(t, tt.misfiled, findMisFiledInSack(tt.input))

		})
	}
}

func TestItemPriority(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    byte
		priority int
	}{
		{
			input:    byte('a'),
			priority: 1,
		},
		{
			input:    byte('z'),
			priority: 26,
		},
		{
			input:    byte('A'),
			priority: 27,
		},
		{
			input:    byte('Z'),
			priority: 52,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.input), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.priority, itemPriority(tt.input))
		})
	}
}

func TestFindBadge(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		sack1 []byte
		sack2 []byte
		sack3 []byte
		badge byte
	}{
		{
			sack1: []byte("vJrwpWtwJgWrhcsFMMfFFhFp"),
			sack2: []byte("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
			sack3: []byte("PmmdzqPrVvPwwTWBwg"),
			badge: byte('r'),
		},
		{

			sack1: []byte("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
			sack2: []byte("ttgJtRGJQctTZtZT"),
			sack3: []byte("CrZsJsPPZsGzwwsLwLmpwMDw"),
			badge: byte('Z'),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.badge, findBadge(tt.sack1, tt.sack2, tt.sack3))
		})
	}
}
