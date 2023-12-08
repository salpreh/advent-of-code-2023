package main

import (
	"com.github/salpreh/advent-of-code-2023/day4/lotterycard"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
)

func main() {
	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", lotterycard.SumLotteryPoints(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", lotterycard.SumLotteryPoints(input))

	input = getPart2ExampleInput()
	fmt.Printf("Part 2 example result: %d\n", lotterycard.SumLotteryCards(input))

	input = getPart2Input()
	fmt.Printf("Part 2 result: %d\n", lotterycard.SumLotteryCards(input))
}

func getPart1ExampleInput() []lotterycard.Card {
	input := utils.ReadInputFile("input/p1Example.txt")

	return lotterycard.ParseLotteryCards(input)
}

func getPart1Input() []lotterycard.Card {
	input := utils.ReadInputFile("input/p1.txt")

	return lotterycard.ParseLotteryCards(input)
}

func getPart2ExampleInput() []lotterycard.Card {
	input := utils.ReadInputFile("input/p1Example.txt")

	return lotterycard.ParseLotteryCards(input)
}

func getPart2Input() []lotterycard.Card {
	input := utils.ReadInputFile("input/p1.txt")

	return lotterycard.ParseLotteryCards(input)
}
