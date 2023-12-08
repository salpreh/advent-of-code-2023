package calibrator

import (
	"com.github/salpreh/advent-of-code-2023/day1/calibrator"
	"com.github/salpreh/advent-of-code-2023/utils"
	"testing"
)

func TestBasicTrebuchetCalibrationExample(t *testing.T) {
	// given
	expected := 142
	input := utils.ReadInputFile("../../input/p1Example.txt")

	// when
	result := calibrator.CalculateBasicTrebuchetCalibration(input)

	// then
	assertEquals(t, expected, result)
}

func TestBasicTrebuchetCalibration(t *testing.T) {
	// given
	expected := 54605
	input := utils.ReadInputFile("../../input/p1.txt")

	// when
	result := calibrator.CalculateBasicTrebuchetCalibration(input)

	// then
	assertEquals(t, expected, result)
}

func TestTrebuchetCalibrationExample(t *testing.T) {
	// given
	expected := 281
	input := utils.ReadInputFile("../../input/p2Example.txt")

	// when
	result := calibrator.CalculateTrebuchetCalibration(input)

	// then
	assertEquals(t, expected, result)
}

func TestTrebuchetCalibration(t *testing.T) {
	// given
	expected := 55429
	input := utils.ReadInputFile("../../input/p2.txt")

	// when
	result := calibrator.CalculateTrebuchetCalibration(input)

	// then
	assertEquals(t, expected, result)
}

func assertEquals(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("Unexpected result, expected %d got %d", expected, result)
	}
}
