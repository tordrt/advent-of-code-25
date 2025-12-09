package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RedTile struct {
	X, Y int
}

func main() {
	file, err := os.Open("9/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var redTiles []RedTile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		if len(line) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		tile := RedTile{}
		tile.X, err = strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		tile.Y, err = strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		redTiles = append(redTiles, tile)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	maxTiles := 0
	for x, tile1 := range redTiles[:len(redTiles)-1] {
		for _, tile2 := range redTiles[x+1:] {
			width := abs(tile1.X-tile2.X) + 1
			height := abs(tile1.Y-tile2.Y) + 1
			tileCount := width * height
			if tileCount > maxTiles {
				maxTiles = tileCount
			}
		}
	}

	fmt.Println(maxTiles)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
