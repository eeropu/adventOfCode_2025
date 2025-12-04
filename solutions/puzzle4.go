package solutions

import (
	"fmt"
	"strings"
	"github.com/eeropu/adventOfCode_2025/assets"
)

var xAxisLength int
var yAxisLength int
var numOfOccupiedNeighbours [][]int
var grid [][]byte


func SolvePuzzle4() {
	lines := strings.Split(strings.TrimSpace(assets.Puzzle4Input), "\n")

	xAxisLength = len(lines)
	yAxisLength = len(lines[0])

	numOfOccupiedNeighbours = make([][]int, xAxisLength)
	for i := range numOfOccupiedNeighbours {
			numOfOccupiedNeighbours[i] = make([]int, yAxisLength)
	}

	grid = make([][]byte, xAxisLength)
	for i := range grid {
		grid[i] = []byte(lines[i])
	}

	prevResult := -1
	result := 0
	for {
		result += shipOutReachable()
		if (prevResult == result) {
			break
		}
		prevResult = result
	}

	fmt.Println(result)
}

func shipOutReachable() int {

	for x := 0; x < xAxisLength; x++ {
		for y := 0; y < yAxisLength; y++ {
			if (grid[x][y] == '.') {
				numOfOccupiedNeighbours[x][y] = -10
				continue
			}
			updateNeighbours(x, y, numOfOccupiedNeighbours)
		}
	}

	result := 0
	for x, row := range numOfOccupiedNeighbours {
		for y, cell := range row {
			if cell >= 0 && cell < 4 {
				result++
				grid[x][y] = '.'
			}
			numOfOccupiedNeighbours[x][y] = 0
		}
	}
	return result
}

func updateNeighbours(x int, y int, s [][]int) {
	for dx := -1; dx <= 1; dx++ {
    for dy := -1; dy <= 1; dy++ {

			if dx == 0 && dy == 0 {
					continue
			}

			nx := x + dx
			ny := y + dy

			if nx >= 0 && nx < xAxisLength && ny >= 0 && ny < yAxisLength {
					s[nx][ny]++
			}
    }
}
}