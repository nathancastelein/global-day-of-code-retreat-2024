package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Boolean bool

type Grid [][]Boolean

type Position struct {
	x int
	y int
}

type Cell struct {
	position Position
	isAlive  Boolean
}

func TestGetNeighbours(t *testing.T) {
	grid := Grid{
		{false, false},
		{false, true},
	}
	position := Position{1, 1}
	expectedNeighbours := []Cell{
		{
			Position{0, 0},
			false,
		},
		{
			Position{0, 1},
			false,
		},
		{
			Position{1, 0},
			false,
		},
	}

	neighbours := getNeighbours(grid, position)

	require.Equal(t, expectedNeighbours, neighbours)
}

func TestOutOfBounds(t *testing.T) {
	grid := Grid{
		{false, false},
		{false, true},
	}

	type test struct {
		position       Position
		expectedResult Boolean
	}

	positionsToTest := []test{
		{
			Position{0, 0},
			false,
		},
		{
			Position{0, 1},
			false,
		},
		{
			Position{0, 2},
			true,
		},
		{
			Position{1, 1},
			false,
		},
		{
			Position{1, 2},
			true,
		},
		{
			Position{1, 0},
			false,
		},
		{
			Position{2, 0},
			true,
		},
	}

	for _, test := range positionsToTest {
		got := outOfBounds(grid, test.position)

		require.Equal(t, test.expectedResult, got)
	}
}

func outOfBounds(grid Grid, position Position) Boolean {
	if position.x < 0 || position.y < 0 {
		return true
	}

	if position.x > len(grid)-1 {
		return true
	}

	if position.y > len(grid[position.x])-1 {
		return true
	}

	return false
}

func getNeighbours(grid Grid, position Position) []Cell {
	neighbours := []Cell{}

	for xIteration := -1; xIteration <= 1; xIteration++ {
		for yIteration := -1; yIteration <= 1; yIteration++ {
			if xIteration == 0 && yIteration == 0 {
				continue
			}

			checkedPosition := Position{position.x + xIteration, position.y + yIteration}
			if !outOfBounds(grid, checkedPosition) {
				neighbours = append(neighbours, Cell{
					position: checkedPosition,
					isAlive:  grid[checkedPosition.x][checkedPosition.y],
				})
			}
		}
	}
	return neighbours
}

func numberOfAdjacentCellsAlive(grid Grid, position Position) int {
	neighbours := getNeighbours(grid, position)

	numberOfAliveNeighbours := 0
	for _, neighbour := range neighbours {
		if neighbour.isAlive {
			numberOfAliveNeighbours++
		}
	}

	return numberOfAliveNeighbours
}

func NextState(grid Grid, cell Cell) Cell {
	numberOfAliveNeighbours := numberOfAdjacentCellsAlive(grid, cell.position)

	nextCell := Cell{
		position: cell.position,
		isAlive:  false,
	}
	if !cell.isAlive {
		if numberOfAliveNeighbours == 3 {
			nextCell.isAlive = true
		}
	}
	return nextCell
}

func TestNextState(t *testing.T) {
	grid := Grid{
		{false, true},
		{false, true},
	}

	adjacentCellsAlive := numberOfAdjacentCellsAlive(grid, Position{1, 1})

	require.Equal(t, 1, adjacentCellsAlive)
}

func TestGrid(t *testing.T) {
	grid := Grid{
		{false, true},
		{false, true},
	}

	adjacentCellsAlive := numberOfAdjacentCellsAlive(grid, Position{1, 1})

	require.Equal(t, 1, adjacentCellsAlive)
}
