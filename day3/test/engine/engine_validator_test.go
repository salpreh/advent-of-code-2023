package engine

import (
	"com.github/salpreh/advent-of-code-2023/day3/engine"
	"com.github/salpreh/advent-of-code-2023/utils"
	"testing"
)

func TestSumSchematicEngineNumsExample(t *testing.T) {
	// given
	expected := 4361
	input := utils.ReadInputFile("../../input/p1Example.txt")

	// when
	result := engine.SumSchematicEngineNums(input)

	// then
	assertEquals(t, expected, result)
}

func TestSumSchematicEngineNums(t *testing.T) {
	// given
	expected := 536202
	input := utils.ReadInputFile("../../input/p1.txt")

	// when
	result := engine.SumSchematicEngineNums(input)

	// then
	assertEquals(t, expected, result)
}

func TestSumSchematicEngineGearRatiosExample(t *testing.T) {
	// given
	expected := 467835
	input := utils.ReadInputFile("../../input/p1Example.txt")

	// when
	result := engine.SumSchematicEngineGearRatios(input)

	// then
	assertEquals(t, expected, result)
}

func TestSumSchematicEngineGearRatios(t *testing.T) {
	// given
	expected := 78272573
	input := utils.ReadInputFile("../../input/p1.txt")

	// when
	result := engine.SumSchematicEngineGearRatios(input)

	// then
	assertEquals(t, expected, result)
}

func assertEquals(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("Unexpected result, expected %d got %d", expected, result)
	}
}
