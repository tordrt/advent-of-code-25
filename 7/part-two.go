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
	// Track position -> number of timelines with a particle at that position
	beamPositions := map[int]int{startIdx: 1}

	for _, line := range lines[:len(lines)-1] {
		nextPositions := make(map[int]int)
		for pos, count := range beamPositions {
			if line[pos] == '^' {
				// Each timeline splits into 2: one going left, one going right
				nextPositions[pos-1] += count
				nextPositions[pos+1] += count
			} else {
				nextPositions[pos] += count
			}
		}
		beamPositions = nextPositions
	}

	// Sum all timelines
	total := 0
	for _, count := range beamPositions {
		total += count
	}
	fmt.Println(total)
}
