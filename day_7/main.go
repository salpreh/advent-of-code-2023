package main

import (
	"com.github/salpreh/advent-of-code-2023/day_7/cardgame"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", cardgame.CalculateGameTotalWinnings(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", cardgame.CalculateGameTotalWinnings(input))
}

func getPart1ExampleInput() []cardgame.CardHand {
	input := utils.ReadInputFile("input/p1Example.txt")
	return parseCardGame(input)
}

func getPart1Input() []cardgame.CardHand {
	input := utils.ReadInputFile("input/p1.txt")
	return parseCardGame(input)
}

func parseCardGame(input []string) []cardgame.CardHand {
	hands := make([]cardgame.CardHand, 0)
	for _, handData := range input {
		data := strings.Split(handData, " ")
		bid, _ := strconv.Atoi(data[1])
		hands = append(hands, *cardgame.NewHand(data[0], bid))
	}

	return hands
}
