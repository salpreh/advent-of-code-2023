package game

import (
	"com.github/salpreh/advent-of-code-2023/day2/game"
	"com.github/salpreh/advent-of-code-2023/utils"
	"testing"
)

func TestSumPossibleGamesExamples(t *testing.T) {
	// given
	expected := 8
	input := getParsedInput("../../input/p1Example.txt")
	gamePieces := getGamePieces()

	// when
	result := game.SumPossibleGames(input, gamePieces)

	// then
	assertEquals(t, expected, result)
}

func TestSumPossibleGames(t *testing.T) {
	// given
	expected := 2439
	input := getParsedInput("../../input/p1.txt")
	gamePieces := getGamePieces()

	// when
	result := game.SumPossibleGames(input, gamePieces)

	// then
	assertEquals(t, expected, result)
}

func TestSumGameRecordsPowersExample(t *testing.T) {
	// given
	expected := 2286
	input := getParsedInput("../../input/p2Example.txt")

	// when
	result := game.SumGameRecordsPowers(input)

	// then
	assertEquals(t, expected, result)
}

func TestSumGameRecordsPowers(t *testing.T) {
	// given
	expected := 63711
	input := getParsedInput("../../input/p2.txt")

	// when
	result := game.SumGameRecordsPowers(input)

	// then
	assertEquals(t, expected, result)
}

func assertEquals(t *testing.T, expected int, result int) {
	if result != expected {
		t.Errorf("Unexpected result, expected %d got %d", expected, result)
	}
}

func getParsedInput(filePath string) []game.Results {
	lines := utils.ReadInputFile(filePath)

	return game.ParseInputLines(lines)
}

func getGamePieces() game.Record {
	return game.Record{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
}
