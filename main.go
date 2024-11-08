package main

import "fmt"

func main() {
	grid := [][]bool{}

	for range 10 {
		line := []bool{}
		for range 10 {
			line = append(line, false)
		}
		grid = append(grid, line)
	}

	fmt.Println(grid)
}

func numberOfAdjacent(grid [][]bool, lineIndex, columnIndex int) int {
	numberOfAdjacent := 0

	if lineIndex-1 >= 0 {
		if grid[lineIndex-1][columnIndex] {
			numberOfAdjacent++
		}
	}

	if lineIndex+1 < len(grid) {
		if grid[lineIndex+1][columnIndex] {
			numberOfAdjacent++
		}
	}

	if columnIndex-1 >= 0 {
		if grid[lineIndex][columnIndex-1] {
			numberOfAdjacent++
		}
	}

	if columnIndex+1 < len(grid[lineIndex]) {
		if grid[lineIndex][columnIndex+1] {
			numberOfAdjacent++
		}
	}

	return numberOfAdjacent
}

func BecomeAlive(grid [][]bool, lineNumber, columnNumber int) bool {
	numberOfAdjacent := numberOfAdjacent(grid, lineNumber, columnNumber)

	return numberOfAdjacent == 3
}

func StayAlive(grid [][]bool, lineNumber, columnNumber int) bool {
	numberOfAdjacent := numberOfAdjacent(grid, lineNumber, columnNumber)

	return (numberOfAdjacent == 2) || (numberOfAdjacent == 3)
}

func GridRun(grid [][]bool) [][]bool {
	nextGrid := make([][]bool, len(grid))
	for lineNumber, line := range grid {
		nextGrid[lineNumber] = make([]bool, len(line))
		for columnNumber, isAlive := range line {
			if !isAlive {
				nextGrid[lineNumber][columnNumber] = BecomeAlive(grid, lineNumber, columnNumber)
			} else {
				nextGrid[lineNumber][columnNumber] = StayAlive(grid, lineNumber, columnNumber)
			}
		}
	}

	return nextGrid
}

/*
func firstRule(grid [][]bool) {
	for line := range grid {
		for element := range line {

		}
	}
}
*/
