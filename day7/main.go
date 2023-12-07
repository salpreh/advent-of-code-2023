package main

import (
	"com.github/salpreh/advent-of-code-2023/day7/cardgame"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
)

func main() {
	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", cardgame.CalculateGameTotalWinnings(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", cardgame.CalculateGameTotalWinnings(input))

	input = getPart2ExampleInput()
	fmt.Printf("Part 2 example result: %d\n", cardgame.CalculateGameTotalWinnings(input))

	input = getPart2Input()
	fmt.Printf("Part 2 result: %d\n", cardgame.CalculateGameTotalWinnings(input))
}

func getPart1ExampleInput() []cardgame.CardHand {
	input := utils.ReadInputFile("input/p1Example.txt")
	return cardgame.ParseRegularCardGame(input)
}

func getPart1Input() []cardgame.CardHand {
	input := utils.ReadInputFile("input/p1.txt")
	return cardgame.ParseRegularCardGame(input)
}

func getPart2ExampleInput() []cardgame.CardHand {
	input := utils.ReadInputFile("input/p1Example.txt")
	return cardgame.ParseJokerCardGame(input)
}

func getPart2Input() []cardgame.CardHand {
	input := utils.ReadInputFile("input/p1.txt")
	return cardgame.ParseJokerCardGame(input)
}
