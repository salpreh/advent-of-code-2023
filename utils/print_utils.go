package utils

import (
	"fmt"
)

func PrintP12Solutions(part1Example int, part1 int, part2Example int, part2 int) {
	PrintP1Solutions(part1Example, part1)
	fmt.Printf("Part 2 example result: %d\n", part2Example)
	fmt.Printf("Part 2 result: %d\n", part2)
}

func PrintP1Solutions(part1Example int, part1 int) {
	fmt.Printf("Part 1 example result: %d\n", part1Example)
	fmt.Printf("Part 1 result: %d\n", part1)
}
