package main

import (
	"com.github/salpreh/advent-of-code-2023/day2/game"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const matchSep = ";"

var (
	gameRegex, _  = regexp.Compile(`Game (\d+):`)
	blueRegex, _  = regexp.Compile(`(\d+) blue`)
	redRegex, _   = regexp.Compile(`(\d+) red`)
	greenRegex, _ = regexp.Compile(`(\d+) green`)
)

func main() {
	gamePieces := game.Record{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", game.SumPossibleGames(input, gamePieces))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", game.SumPossibleGames(input, gamePieces))
}

func getPart1ExampleInput() []game.Results {
	lines := utils.ReadInputFile("input/p1Example.txt")

	return parseInputLines(lines)
}

func getPart1Input() []game.Results {
	lines := utils.ReadInputFile("input/p1.txt")

	return parseInputLines(lines)
}

func parseInputLines(lines []string) []game.Results {
	gamesResults := make([]game.Results, 0)
	for _, line := range lines {
		gameNum, _ := strconv.Atoi(gameRegex.FindStringSubmatch(line)[1])
		gameRecords := make([]game.Record, 0)
		for _, match := range strings.Split(line, matchSep) {
			gameRecord := game.Record{}

			blueNum := blueRegex.FindStringSubmatch(match)
			if blueNum != nil && len(blueNum) > 1 {
				gameRecord.Blue, _ = strconv.Atoi(blueNum[1])
			}

			redNum := redRegex.FindStringSubmatch(match)
			if redNum != nil && len(redNum) > 1 {
				gameRecord.Red, _ = strconv.Atoi(redNum[1])
			}

			greenNum := greenRegex.FindStringSubmatch(match)
			if greenNum != nil && len(greenNum) > 1 {
				gameRecord.Green, _ = strconv.Atoi(greenNum[1])
			}

			gameRecords = append(gameRecords, gameRecord)
		}

		gamesResults = append(gamesResults, game.Results{gameNum, gameRecords})
	}

	return gamesResults
}
