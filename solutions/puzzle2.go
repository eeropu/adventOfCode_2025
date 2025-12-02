package solutions

import (
	"math"
	"strings"
	"strconv"
	"fmt"
	"github.com/eeropu/adventOfCode_2025/assets"
)

func SolvePuzzle2() {
	idRanges := strings.Split(strings.TrimSpace(assets.Puzzle2Input), ",")

	invalidIds := []int{}
	for _, idRange := range idRanges {
		i := strings.Index(idRange, "-")

		start, _ := strconv.Atoi(idRange[:i])
		endStr := idRange[i+1:]
		end, _ := strconv.Atoi(endStr)

		maxLength := len(endStr)
		for j := maxLength; j > 0; j-- {
			invalidIds = append(invalidIds, checkRepetitionsOfLength(start, end, j, maxLength)...)
		}
	}

	sum := 0
	for _, value := range uniqueInts(invalidIds) {
		sum += value
	}
	fmt.Println(sum)
}

func padding(n int, repetitions int) int {
    s := strconv.Itoa(n)
    repeated := strings.Repeat(s, repetitions)
    result, _ := strconv.Atoi(repeated)
    return result
}

func checkRepetitionsOfLength(start int, end int, length int, maxLength int) []int {
	min := math.Pow(10, float64(length - 1))
	max := math.Pow(10, float64(length))

	result := []int{}
	for j := int(min); j < int(max); j++ {
		for repetitions := 2; repetitions <= maxLength / length; repetitions++ {

			valueToCheck := padding(j, repetitions)
			if (valueToCheck >= start && valueToCheck <= end) {
				result = append(result, valueToCheck)
			}

			if (valueToCheck > end) {
				break
			}
		}

	}
	return result
}

func uniqueInts(xs []int) []int {
    seen := make(map[int]bool)
    result := []int{}

    for _, v := range xs {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }
    return result
}