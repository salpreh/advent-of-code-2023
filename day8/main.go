package main

import (
	"com.github/salpreh/advent-of-code-2023/day8/mapreader"
	"com.github/salpreh/advent-of-code-2023/utils"
)

func main() {
	utils.PrintP1Solutions(
		mapreader.GetStepsTillDestination(getPart1ExampleInput()),
		mapreader.GetStepsTillDestination(getPart1Input()),
	)
}

func getPart1ExampleInput() mapreader.Map {
	input := utils.ReadInputFile("input/p1Example.txt")
	return mapreader.ParseMap(input)
}

func getPart1Input() mapreader.Map {
	input := utils.ReadInputFile("input/p1.txt")
	return mapreader.ParseMap(input)
}
