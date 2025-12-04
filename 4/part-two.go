package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func countNeighboursGrid(grid [][]byte, row, col int) int {
	count := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			r, c := row+dr, col+dc
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) && grid[r][c] == '@' {
				count++
			}
		}
	}
	return count
}

type Index struct {
	Row int // row
	Col int // column
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

	// Convert to [][]byte for removing rolls from the grid
	grid := make([][]byte, len(lines)+2)
	padding := bytes.Repeat([]byte{'.'}, len(lines[0]))
	grid[0] = padding
	for i, line := range lines {
		grid[i+1] = []byte(line)
	}
	grid[len(grid)-1] = padding

	totalRolls := 0

	for {
		var rollsToRemove []Index
		for row := 1; row < len(grid)-1; row++ {
			for col, ch := range grid[row] {
				if ch == '@' && countNeighboursGrid(grid, row, col) < 4 {
					rollsToRemove = append(rollsToRemove, Index{Row: row, Col: col})
				}
			}
		}

		if len(rollsToRemove) == 0 {
			break
		}

		for _, roll := range rollsToRemove {
			grid[roll.Row][roll.Col] = '.'
		}

		totalRolls += len(rollsToRemove)
	}

	fmt.Println(totalRolls)
}
