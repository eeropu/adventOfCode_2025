package solutions

import (
	"fmt"
	"strings"
	"github.com/eeropu/adventOfCode_2025/assets"
)

type set map[[2]int]int

func (s set) add(x int, y int, v int) {
	s[[2]int{x, y}] = v
}

func (s set) remove(x int, y int) {
	delete(s, [2]int{x, y})
}

func (s set) has(x int, y int) bool {
	_, ok := s[[2]int{x, y}]
	return ok
}

func (s set) get(x int, y int) int {
	value, _ := s[[2]int{x, y}]
	return value
}

var s set = set{}
var lines []string
func SolvePuzzle7() {
	lines = strings.Split(strings.TrimSpace(assets.Puzzle7Input), "\n")

	start := 0
	for i, value := range lines[0] {
		if rune(value) == 'S' {
			start = i
			break
		}
	}

	result := findPaths(start, 0)
	fmt.Println(result)
}

func findPaths(x int, y int) int {
	if (s.has(x, y)) {
		return s.get(x, y)
	}
	for i, line := range lines[y + 1:] {
		if rune(line[x]) == '^' {
			result := 0
			if x > 0 {
				result += findPaths(x - 1, y + i)
			}
			if x < len(line) - 1 {
				result += findPaths(x + 1, y + i)
			}
			s.add(x, y, result)
			return result
		}
	}
	s.add(x, y, 1)
	return 1
}