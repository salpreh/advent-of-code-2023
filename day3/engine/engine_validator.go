package engine

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	emptyMark = '.'
	gearMark  = '*'
)

func SumSchematicEngineNums(schematic []string) int {
	engineSum := 0
	res := ParseEnginePartNumbers(schematic)
	for _, num := range res {
		engineSum += num
	}

	return engineSum
}

func SumSchematicEngineGearRatios(schematic []string) int {
	gearRatioSum := 0
	res := ParseEngineGearRatios(schematic)
	for _, num := range res {
		gearRatioSum += num
	}

	return gearRatioSum
}

func ParseEnginePartNumbers(schematic []string) []int {
	engineNums := make([]int, 0)
	for i, line := range schematic {
		currentNum := make([]rune, 0)
		isValid := false
		for j, char := range line {
			if unicode.IsDigit(char) { // If digit keep building the number
				currentNum = append(currentNum, char)
				isValid = isValid || validateEngineNumber(schematic, j, i)
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

func ParseEngineGearRatios(schematic []string) []int {
	gearRatios := make([]int, 0)
	for i, line := range schematic {
		for j, _ := range line {
			if ratio, isValid := validateGearNumber(schematic, j, i); isValid {
				gearRatios = append(gearRatios, ratio)
			}
		}
	}

	return gearRatios
}

func validateGearNumber(schematic []string, x int, y int) (int, bool) {
	if !isGearMarker(rune(schematic[y][x])) {
		return -1, false
	}

	gearNums := make([]int, 0)
	// Upper numbers
	if y > 0 && unicode.IsDigit(rune(schematic[y-1][x])) {
		num, _ := getSchematicNumber(schematic, x, y-1)
		gearNums = append(gearNums, num)
	} else if y > 0 {
		if num, err := getSchematicNumber(schematic, x-1, y-1); err == nil {
			gearNums = append(gearNums, num)
		}
		if num, err := getSchematicNumber(schematic, x+1, y-1); err == nil {
			gearNums = append(gearNums, num)
		}
	}

	// Left number
	if x > 0 {
		if num, err := getSchematicNumber(schematic, x-1, y); err == nil {
			gearNums = append(gearNums, num)
		}
	}
	// Right number
	if x < len(schematic[y]) {
		if num, err := getSchematicNumber(schematic, x+1, y); err == nil {
			gearNums = append(gearNums, num)
		}
	}

	// Lower numbers
	if y < len(schematic)-1 && unicode.IsDigit(rune(schematic[y+1][x])) {
		num, _ := getSchematicNumber(schematic, x, y+1)
		gearNums = append(gearNums, num)
	} else if y > 0 {
		if num, err := getSchematicNumber(schematic, x-1, y+1); err == nil {
			gearNums = append(gearNums, num)
		}
		if num, err := getSchematicNumber(schematic, x+1, y+1); err == nil {
			gearNums = append(gearNums, num)
		}
	}

	if len(gearNums) != 2 {
		return -1, false
	}

	return gearNums[0] * gearNums[1], true
}

func getSchematicNumber(schematic []string, x int, y int) (int, error) {
	if y < 0 || x < 0 || y >= len(schematic) || x >= len(schematic[y]) {
		return -1, fmt.Errorf("schema position out of bounds: %d, %d", x, y)
	}
	if !unicode.IsDigit(rune(schematic[y][x])) {
		return -1, fmt.Errorf("schema char in %d, %d is not a number: %d", x, y, schematic[y][x])
	}

	numChars := make([]rune, 0)
	currentX := x
	// First we go top left to start building the number
	for ; currentX > 0 && unicode.IsDigit(rune(schematic[y][currentX-1])); currentX-- {
	}

	// From top left build the number
	for ; currentX < len(schematic[y]) && unicode.IsDigit(rune(schematic[y][currentX])); currentX++ {
		numChars = append(numChars, rune(schematic[y][currentX]))
	}

	return strconv.Atoi(string(numChars))
}

func validateEngineNumber(schematic []string, x int, y int) bool {
	for i := y - 1; i <= y+1; i += 1 {
		if i < 0 || i > len(schematic)-1 {
			continue
		}

		for j := x - 1; j <= x+1; j += 1 {
			if j < 0 || j > len(schematic[i]) || (i == y && j == x) {
				continue
			}

			if isEngineMarker(rune(schematic[i][j])) {
				return true
			}
		}
	}

	return false
}

func isEngineMarker(mark rune) bool {
	return !unicode.IsDigit(mark) && mark != emptyMark
}

func isGearMarker(mark rune) bool {
	return mark == gearMark
}
