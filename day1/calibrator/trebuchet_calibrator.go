package calibrator

import (
	"fmt"
	"strconv"
	"unicode"
)

func CalculateBasicTrebuchetCalibration(input []string) int {
	var result int
	for _, line := range input {
		result += calculateLineDigits(line)
	}

	return result
}

func calculateLineDigits(line string) int {
	var first *rune
	var second *rune = new(rune)

	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == nil {
				first = new(rune)
				*first = char - '0'
			}
			*second = char - '0'
		}
	}

	result, _ := strconv.Atoi(fmt.Sprintf("%d%d", *first, *second))

	return result
}
