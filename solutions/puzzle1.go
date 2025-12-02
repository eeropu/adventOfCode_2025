package solutions

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/eeropu/adventOfCode_2025/assets"
)


func SolvePuzzle1() {
	lines := strings.Split(strings.TrimSpace(assets.Puzzle1Input), "\n")

	position := 50
	timesAtZero := 0

	for _, line := range lines {
		prevPosition := position

		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		fullSpins := steps / 100
		timesAtZero += fullSpins

		steps %= 100

		if dir == 'R' {
			position += steps
		} else {
			position -= steps
		}

		if position < 0 {
			position += 100
			if (prevPosition != 0) {
				timesAtZero++
			}
		} else if (position == 0) {
			timesAtZero++
		} else if (position >= 100) {
			position -= 100
			timesAtZero++
		}
	}
	fmt.Println(timesAtZero)
}
