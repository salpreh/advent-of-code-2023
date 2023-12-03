package main

import (
	"com.github/salpreh/advent-of-code-2023/day2/game"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
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

	input = getPart2ExampleInput()
	fmt.Printf("Part 2 example result: %d\n", game.SumGameRecordsPowers(input))

	input = getPart2Input()
	fmt.Printf("Part 2 result: %d\n", game.SumGameRecordsPowers(input))
}

func getPart1ExampleInput() []game.Results {
	lines := utils.ReadInputFile("input/p1Example.txt")

	return game.ParseInputLines(lines)
}

func getPart1Input() []game.Results {
	lines := utils.ReadInputFile("input/p1.txt")

	return game.ParseInputLines(lines)
}

func getPart2ExampleInput() []game.Results {
	lines := utils.ReadInputFile("input/p2Example.txt")

	return game.ParseInputLines(lines)
}

func getPart2Input() []game.Results {
	lines := utils.ReadInputFile("input/p2.txt")

	return game.ParseInputLines(lines)
}
