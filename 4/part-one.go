package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func countNeighbours(window []string, i int) int {
	count := 0
	for row := 0; row < 3; row++ {
		for col := i - 1; col <= i+1; col++ {
			if col < 0 || col >= len(window[row]) {
				continue
			}
			if row == 1 && col == i {
				continue // skip self
			}
			if window[row][col] == '@' {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("4/input.txt")
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

	if len(lines) == 0 {
		panic("no input")
	}

	padding := strings.Repeat(".", len(lines[0]))
	lines = append([]string{padding}, append(lines, padding)...)

	totalRolls := 0
	for row := 1; row < len(lines)-1; row++ {
		window := lines[row-1 : row+2]
		for i, ch := range lines[row] {
			if ch == '@' && countNeighbours(window, i) < 4 {
				totalRolls++
			}
		}
	}

	log.Println(totalRolls)
}
