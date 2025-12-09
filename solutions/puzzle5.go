package solutions

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
	"github.com/eeropu/adventOfCode_2025/assets"
)

type idRange struct {
	min int
	max int
}

var ranges []idRange

func SolvePuzzle5() {
	lines := strings.Split(strings.TrimSpace(assets.Puzzle5Input), "\n")

	for _, line := range lines {
		if line == "" {
			break
		}
		
		parts := strings.Split(line, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		r := idRange{min, max}
		appendRange(r)
	}

	idCount := 0
	for _, r := range ranges {
		idCount += r.max - r.min + 1
	}

	fmt.Println(idCount)
}


func appendRange(r idRange) {
    ranges = append(ranges, r)
    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].min < ranges[j].min
    })

    merged := []idRange{}
    current := ranges[0]

    for _, next := range ranges[1:] {
        if next.min <= current.max {
            if next.max > current.max {
                current.max = next.max
            }
        } else {
            merged = append(merged, current)
            current = next
        }
    }

    merged = append(merged, current)
    ranges = merged
}
