package solutions

import (
	"fmt"
	"math"
	"strings"
	"github.com/eeropu/adventOfCode_2025/assets"
)

func SolvePuzzle3() {
	lines := strings.Split(strings.TrimSpace(assets.Puzzle3Input), "\n")

	sum := 0
	for _, line := range lines {

		lineLength := len(line)
		outer := make([][]int, 10)

		for i, _ := range outer {
			outer[i] = []int{}
		}
		
		for i, r := range line {
				current := int(r - '0')
				outer[current] = append(outer[current], i)
		}

		result := 0
		index := 0
		for i := 12; i > 0; i-- {
			for j := 9; j >= 0; j-- {
				removeablePassedIndices := 0
				for z := 0; z < len(outer[j]); z++ {
					if (outer[j][z] < index) {
						removeablePassedIndices = z + 1
					} else {
						break
					}
				}
				outer[j] = outer[j][removeablePassedIndices:]
				if (len(outer[j]) > 0 && outer[j][0] <= lineLength - (i)) {
					index = outer[j][0]
					result += j * int(math.Pow(10, float64(i-1)))
					outer[j] = outer[j][1:]
					break
				}
			}
		}

		//fmt.Printf("Result for line %s is: %d\n", line, result)
		sum += result
	}
	fmt.Printf("Sum is: %d\n", sum)
}