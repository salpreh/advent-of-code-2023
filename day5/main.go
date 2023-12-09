package main

import (
	"com.github/salpreh/advent-of-code-2023/day5/garden"
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
)

func main() {
	input := getPart1ExampleInput()
	fmt.Printf("Part 1 example result: %d\n", garden.GetLowestSeedLocation(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", garden.GetLowestSeedLocation(input))

	input = getPart2ExampleInput()
	fmt.Printf("Part 2 example result: %d\n", garden.GetLowestSeedRangedLocation(input))

	input = getPart2Input()
	fmt.Printf("Part 2 result: %d\n", garden.GetLowestSeedRangedLocation(input))
}

func getPart1ExampleInput() garden.Almanac {
	input := utils.ReadInputFile("input/p1Example.txt")

	return garden.ParseGardenAlmanac(input)
}

func getPart1Input() garden.Almanac {
	input := utils.ReadInputFile("input/p1.txt")

	return garden.ParseGardenAlmanac(input)
}

func getPart2ExampleInput() garden.Almanac {
	input := utils.ReadInputFile("input/p1Example.txt")

	return garden.ParseGardenAlmanac(input)
}

func getPart2Input() garden.Almanac {
	input := utils.ReadInputFile("input/p1.txt")

	return garden.ParseGardenAlmanac(input)
}
