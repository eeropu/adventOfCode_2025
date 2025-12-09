package solutions

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/eeropu/adventOfCode_2025/assets"
)

func SolvePuzzle6() {
	lines := strings.Split(assets.Puzzle6Input, "\n")

	var sb strings.Builder
	arguments := []int{}
	result := 0
	for j, _ := range lines[0] {
		operation := 0
		col := len(lines[0]) - j - 1

		for i, _ := range lines {
			ascii := lines[i][col]
			if ascii == 42 || ascii == 43 {
				operation = int(ascii)
			} else if ascii >= 48 && ascii <= 57 {
				sb.WriteByte(ascii)
			}
		}

		value := sb.String()
		sb.Reset()
		if value != "" {
			addable, _ := strconv.Atoi(value)
			arguments = append(arguments, addable)
		}

		if (operation != 0) {
			if (operation == 42) {
				result += getProduct(arguments)
			} else {
				result += getSum(arguments)
			}
			arguments = []int{}
		}
	}
	fmt.Println(result)
}

func getProduct(arguments []int) int {
	product := 1
	for _, v := range arguments {
		product *= v
	}
	return product
}

func getSum(arguments []int) int {
	sum := 0
	for _, v := range arguments {
		sum += v
	}
	return sum
}