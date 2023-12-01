package main

import (
	"bufio"
	"com.github/salpreh/advent-of-code-2023/day1/calibrator"
	"fmt"
	"os"
	"path"
	"runtime"
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
	return readInputFile("input/p1Example.txt")
}

func getPart1Input() []string {
	return readInputFile("input/p1.txt")
}

func getPart2ExampleInput() []string {
	return readInputFile("input/p2Example.txt")
}

func getPart2Input() []string {
	return readInputFile("input/p2.txt")
}

func readInputFile(filePath string) []string {
	_, currentFile, _, _ := runtime.Caller(1)
	fullPath := path.Join(path.Dir(currentFile), filePath)
	file, err := os.Open(fullPath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	input := make([]string, 0)
	for reader.Scan() {
		input = append(input, reader.Text())
	}

	return input
}
