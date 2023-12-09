package garden

import (
	"com.github/salpreh/advent-of-code-2023/day5/garden"
	"com.github/salpreh/advent-of-code-2023/utils"
	"testing"
)

func TestLowestSeedLocationExample(t *testing.T) {
	// given
	expected := 35
	input := getInput("../../input/p1Example.txt")

	// when
	result := garden.GetLowestSeedLocation(input)

	// then
	assertEquals(t, expected, result)
}

func TestLowestSeedLocation(t *testing.T) {
	// given
	expected := 289863851
	input := getInput("../../input/p1.txt")

	// when
	result := garden.GetLowestSeedLocation(input)

	// then
	assertEquals(t, expected, result)
}

func TestLowestSeedRangedLocationExample(t *testing.T) {
	// given
	expected := 46
	input := getInput("../../input/p1Example.txt")

	// when
	result := garden.GetLowestSeedRangedLocation(input)

	// then
	assertEquals(t, expected, result)
}

func TestLowestSeedRangedLocation(t *testing.T) {
	// given
	expected := 60568880
	input := getInput("../../input/p1.txt")

	// when
	result := garden.GetLowestSeedRangedLocation(input)

	// then
	assertEquals(t, expected, result)
}

func assertEquals(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("Unexpected result, expected %d got %d", expected, result)
	}
}

func getInput(filePath string) garden.Almanac {
	input := utils.ReadInputFile("../../input/p1Example.txt")

	return garden.ParseGardenAlmanac(input)
}
