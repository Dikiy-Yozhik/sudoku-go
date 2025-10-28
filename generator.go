package main

import (
	"math/rand"
	"time"
)

func NewRandomGame(diff int) *Grid {
	rand.Seed(time.Now().UnixNano())

	grid := generateFullGrid()

	cells := getCellsToRemove(diff)
	removeCells(&grid, cells)

	return &grid
}

func generateFullGrid() Grid {
	var grid Grid

	numbers := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	for i := 0; i < 9; i++ {
		grid[0][i].digit = numbers[i]
		grid[0][i].fixed = true
	}

	solveGrid(&grid)

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			grid[r][c].fixed = true
		}
	}

	return grid
}

func solveGrid(g *Grid) bool {
	for r := 1; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if g[r][c].digit == empty {
				for d := int8(1); d <= 9; d++ {
					err := g.Set(r, c, d)
					if err == nil {
						if solveGrid(g) {
							return true
						}
						g[r][c].digit = empty
					}
				}
				return false
			}
		}
	}
	return true
}

func getCellsToRemove(diff int) int {
	switch diff {
	case 1:
		return 30
	case 2:
		return 40
	case 3:
		return 50
	default:
		return 40
	}
}

func removeCells(g *Grid, cells int) {
	removedCells := 0

	for removedCells < cells {
		row := rand.Intn(rows)
		col := rand.Intn(columns)

		if g[row][col].digit != empty {
			backup := g[row][col].digit
			g[row][col].digit = empty
			g[row][col].fixed = false

			if hasUniqueSolution(g) {
				removedCells++
			} else {
				g[row][col].digit = backup
				g[row][col].fixed = true
			}
		}
	}
}

func hasUniqueSolution(g *Grid) bool {
	var tempg Grid = *g
	solutions := countSolutions(&tempg, 0)
	return solutions == 1
}

func countSolutions(g *Grid, count int) int {
	if count > 1 {
		return count
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if g[r][c].digit == empty {
				for d := int8(1); d <= 9; d++ {
					err := g.Set(r, c, d)
					if err == nil {
						count = countSolutions(g, count)
						g[r][c].digit = empty

						if count > 1 {
							return count
						}
					}
				}
				return count
			}
		}
	}
	return count + 1
}
