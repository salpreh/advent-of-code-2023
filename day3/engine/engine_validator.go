package engine

import (
	"strconv"
	"unicode"
)

func SumSchematicEngineNums(schematic []string) int {
	engineSum := 0
	res := ParseEngineSchema(schematic)
	for _, num := range res {
		engineSum += num
	}

	return engineSum
}

func ParseEngineSchema(lines []string) []int {
	engineNums := make([]int, 0)
	for i, line := range lines {
		currentNum := make([]rune, 0)
		isValid := false
		for j, char := range line {
			if unicode.IsDigit(char) { // If digit keep building the number
				currentNum = append(currentNum, char)
				isValid = isValid || validateEngineNumber(lines, j, i)
			} else if len(currentNum) > 0 { // If not store number if valid
				if isValid {
					engineNum, _ := strconv.Atoi(string(currentNum))
					engineNums = append(engineNums, engineNum)
				}
				currentNum = make([]rune, 0)
				isValid = false
			}
		}

		if len(currentNum) > 0 && isValid { // If number remained in line last iteration
			if isValid {
				engineNum, _ := strconv.Atoi(string(currentNum))
				engineNums = append(engineNums, engineNum)
			}
		}
	}

	return engineNums
}

func validateEngineNumber(lines []string, x int, y int) bool {
	for i := y - 1; i <= y+1; i += 1 {
		if i < 0 || i > len(lines)-1 {
			continue
		}

		for j := x - 1; j <= x+1; j += 1 {
			if j < 0 || j > len(lines[i]) || (i == y && j == x) {
				continue
			}

			if isEngineMarker(rune(lines[i][j])) {
				return true
			}
		}
	}

	return false
}

func isEngineMarker(mark rune) bool {
	return !unicode.IsDigit(mark) && mark != '.'
}
