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
	input := getExampleInput()
	fmt.Printf("Example result: %d\n", calibrator.CalculateBasicTrebuchetCalibration(input))

	input = getPart1Input()
	fmt.Printf("Part 1 result: %d\n", calibrator.CalculateBasicTrebuchetCalibration(input))

}

func getExampleInput() []string {
	return readInputFile("input/example.txt")
}

func getPart1Input() []string {
	return readInputFile("input/p1.txt")
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
