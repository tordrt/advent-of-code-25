package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	startIdx := strings.Index(lines[0], "S")
	beamPositions := map[int]struct{}{startIdx: {}}
	splitCount := 0

	for _, line := range lines[:len(lines)-1] {
		nextPositions := make(map[int]struct{})
		for pos := range beamPositions {
			if line[pos] == '^' {
				splitCount++
				nextPositions[pos-1] = struct{}{}
				nextPositions[pos+1] = struct{}{}
			} else {
				nextPositions[pos] = struct{}{}
			}
		}
		beamPositions = nextPositions
	}

	fmt.Println(splitCount)
}
