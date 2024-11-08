package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberOfAdjacent(t *testing.T) {
	grid := [][]bool{
		{
			false,
			false,
			true,
			false,
			false,
		},
		{
			false,
			true,
			true,
			false,
			false,
		},
		{
			false,
			false,
			true,
			false,
			false,
		},
	}

	numberOfAdjacent := numberOfAdjacent(grid, 1, 2)

	require.Equal(t, numberOfAdjacent, 3)
}

func TestNumberOfAdjacent1(t *testing.T) {
	grid := [][]bool{
		{
			false,
			true,
		},
		{
			true,
			true,
		},
	}

	numberOfAdjacent := numberOfAdjacent(grid, 1, 1)

	require.Equal(t, 2, numberOfAdjacent)
}

func TestIsAlive(t *testing.T) {
	grid := [][]bool{
		{
			false,
			true,
			false,
		},
		{
			true,
			false,
			true,
		},
	}

	result := BecomeAlive(grid, 1, 1)

	require.True(t, result)
}

func TestStayAlive(t *testing.T) {
	grid := [][]bool{
		{
			false,
			true,
		},
		{
			true,
			true,
		},
	}

	result := StayAlive(grid, 1, 1)

	require.True(t, result)
}

func TestGridRun(t *testing.T) {
	grid := [][]bool{
		{
			false,
			true,
			false,
		},
		{
			true,
			false,
			true,
		},
	}

	expectedGrid := [][]bool{
		{
			false,
			false,
			false,
		},
		{
			false,
			true,
			false,
		},
	}

	got := GridRun(grid)

	require.Equal(t, expectedGrid, got)
}

func TestGridRun1(t *testing.T) {
	grid := [][]bool{
		{
			false, false, true, false, false,
		},
		{
			false, true, true, true, false,
		},
		{
			true, false, true, false, true,
		},
	}

	expectedGrid := [][]bool{
		{
			false, false, false, false, false,
		},
		{
			false, false, false, false, false,
		},
		{
			false, true, false, true, false,
		},
	}

	got := GridRun(grid)

	require.Equal(t, expectedGrid, got)
}
