package main

import (
	"fmt"
	"time"
)

const NUMBER_OF_ROWS int = 25
const NUMBER_OF_COLUMNS int = 50

func main() {
	var grid [NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int
	populateSeed(&grid)

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				printGrid(grid)
				nextGeneration(&grid)
			}
		}
	}()

	time.Sleep(3600 * time.Second)
	ticker.Stop()
	done <- true
}

func populateSeed(grid *[NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int) *[NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int {
	centerX := len(grid) / 2
	centerY := len(grid[0]) / 2

	// grid[centerX-1][centerY-1] = 1
	grid[centerX-1][centerY] = 1
	// grid[centerX-1][centerY+1] = 1
	grid[centerX][centerY-1] = 1
	grid[centerX][centerY] = 1
	// grid[centerX][centerY+1] = 1
	// grid[centerX+1][centerY-1] = 1
	grid[centerX+1][centerY] = 1
	grid[centerX+1][centerY+1] = 1

	return grid
}

func nextGeneration(grid *[NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int) *[NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int {
	var nextGenerationGrid [NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int

	for i := 0; i < len(grid); i++ {
		rowBefore := (i - 1 + len(grid)) % len(grid)
		rowAfter := (i + 1) % len(grid)

		for j := 0; j < len(grid[i]); j++ {
			columnBefore := (j - 1 + len(grid[i])) % len(grid[i])
			columnAfter := (j + 1) % len(grid[i])

			neighbours := []int{
				grid[rowBefore][columnBefore],
				grid[rowBefore][j],
				grid[rowBefore][columnAfter],

				grid[i][columnBefore],
				grid[i][columnAfter],

				grid[rowAfter][columnBefore],
				grid[rowAfter][j],
				grid[rowAfter][columnAfter],
			}

			var deadNeighbours int
			var aliveNeighbours int

			for n := 0; n < len(neighbours); n++ {
				if neighbours[n] == 0 {
					deadNeighbours++
				} else {
					aliveNeighbours++
				}

			}

			if grid[i][j] == 0 {
				if aliveNeighbours == 3 {
					nextGenerationGrid[i][j] = 1
					continue
				}
			} else {
				if aliveNeighbours == 2 || aliveNeighbours == 3 {
					nextGenerationGrid[i][j] = 1
					continue
				}

				if aliveNeighbours < 2 {
					nextGenerationGrid[i][j] = 0
					continue
				}

				if aliveNeighbours > 3 {
					nextGenerationGrid[i][j] = 0
					continue
				}
			}
		}
	}

	*grid = nextGenerationGrid
	return grid
}

func printGrid(grid [NUMBER_OF_ROWS][NUMBER_OF_COLUMNS]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var print_square string
			if grid[i][j] < 1 {
				print_square = " "
			} else {
				print_square = "*"
			}

			fmt.Printf("%s ", print_square)
		}
		fmt.Println()
	}
}
