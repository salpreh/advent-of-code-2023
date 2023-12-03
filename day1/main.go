package main

import (
	"com.github/salpreh/advent-of-code-2023/day1/calibrator"
	utils "com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
)

func main() {
	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", calibrator.CalculateBasicTrebuchetCalibration(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", calibrator.CalculateBasicTrebuchetCalibration(input))

	input = getPart2ExampleInput()
	fmt.Printf("Part 2 example result: %d\n", calibrator.CalculateTrebuchetCalibration(input))

	input = getPart2Input()
	fmt.Printf("Part 2 result: %d\n", calibrator.CalculateTrebuchetCalibration(input))
}

func getPart1ExampleInput() []string {
	return utils.ReadInputFile("input/p1Example.txt")
}

func getPart1Input() []string {
	return utils.ReadInputFile("input/p1.txt")
}

func getPart2ExampleInput() []string {
	return utils.ReadInputFile("input/p2Example.txt")
}

func getPart2Input() []string {
	return utils.ReadInputFile("input/p2.txt")
}
