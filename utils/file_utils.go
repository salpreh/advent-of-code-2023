package utils

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

func ReadInputFile(filePath string) []string {
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
