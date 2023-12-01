package calibrator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var textNumbers = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func CalculateBasicTrebuchetCalibration(input []string) int {
	var result int
	for _, line := range input {
		result += calculateLineDigits(line)
	}

	return result
}

func CalculateTrebuchetCalibration(input []string) int {
	var result int
	for _, line := range input {
		result += calculateLineDigitsExtended(line)
	}

	return result
}

func calculateLineDigits(line string) int {
	var first *rune
	var second = new(rune)

	for _, char := range line {
		if unicode.IsDigit(char) {
			first, second = calculateCalibrationNumbers(first, second, char)
		}
	}

	result, _ := strconv.Atoi(fmt.Sprintf("%d%d", *first, *second))

	return result
}

func calculateLineDigitsExtended(line string) int {
	var first *rune
	var second = new(rune)

	for i := 0; i < len(line); {
		char := rune(line[i])
		if unicode.IsDigit(char) { // Is numeric char?
			first, second = calculateCalibrationNumbers(first, second, char)

		} else if unicode.IsLetter(char) { // Is text number?
			numChar := calculateNumberString(line[i:])
			if numChar != nil {
				first, second = calculateCalibrationNumbers(first, second, *numChar)
			}
		}

		i += 1
	}

	result, _ := strconv.Atoi(fmt.Sprintf("%d%d", *first, *second))

	return result
}

func calculateNumberString(input string) *rune {
	var num *rune
	for pos, textNum := range textNumbers {
		if strings.HasPrefix(input, textNum) {
			num = new(rune)
			*num = rune(pos + 1 + '0')
			break
		}
	}

	return num
}

func calculateCalibrationNumbers(first *rune, second *rune, input rune) (*rune, *rune) {
	newFirst := first
	newSecond := second
	if first == nil {
		newFirst = new(rune)
		*newFirst = input - '0'
	}
	*newSecond = input - '0'

	return newFirst, newSecond
}
