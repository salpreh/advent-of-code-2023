package main

import (
	"com.github/salpreh/advent-of-code-2023/day3/engine"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
)

func main() {
	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", engine.SumSchematicEngineNums(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", engine.SumSchematicEngineNums(input))

	input = getPart2ExampleInput()
	fmt.Printf("Part 2 example result: %d\n", engine.SumSchematicEngineGearRatios(input))

	input = getPart2Input()
	fmt.Printf("Part 2 result: %d\n", engine.SumSchematicEngineGearRatios(input))
}

func getPart1ExampleInput() []string {
	return utils.ReadInputFile("input/p1Example.txt")
}

func getPart1Input() []string {
	return utils.ReadInputFile("input/p1.txt")
}

func getPart2ExampleInput() []string {
	return utils.ReadInputFile("input/p1Example.txt")
}

func getPart2Input() []string {
	return utils.ReadInputFile("input/p1.txt")
}
