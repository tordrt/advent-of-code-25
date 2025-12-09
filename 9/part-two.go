package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RedTileData struct {
	X, Y int
}

type Edge struct {
	x1, y1, x2, y2 int
	vertical       bool
}

func main() {
	file, err := os.Open("9/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var redTiles []RedTileData
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		if len(line) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		tile := RedTileData{}
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

	// Build polygon edges from consecutive red tiles
	var edges []Edge
	for i := range redTiles {
		t1 := redTiles[i]
		t2 := redTiles[(i+1)%len(redTiles)]
		edges = append(edges, Edge{
			x1:       t1.X,
			y1:       t1.Y,
			x2:       t2.X,
			y2:       t2.Y,
			vertical: t1.X == t2.X,
		})
	}

	// Collect unique X and Y coordinates for compression
	xSet := make(map[int]struct{})
	ySet := make(map[int]struct{})
	for _, t := range redTiles {
		xSet[t.X] = struct{}{}
		ySet[t.Y] = struct{}{}
	}

	var xCoords, yCoords []int
	for x := range xSet {
		xCoords = append(xCoords, x)
	}
	for y := range ySet {
		yCoords = append(yCoords, y)
	}
	sort.Ints(xCoords)
	sort.Ints(yCoords)

	// Create index maps for fast lookup
	xIndex := make(map[int]int)
	yIndex := make(map[int]int)
	for i, x := range xCoords {
		xIndex[x] = i
	}
	for i, y := range yCoords {
		yIndex[y] = i
	}

	// For each cell in compressed grid, determine if inside polygon
	// A cell [xCoords[i], xCoords[i+1]] x [yCoords[j], yCoords[j+1]]
	// We check using ray casting on the center point

	numCellsX := len(xCoords) - 1
	numCellsY := len(yCoords) - 1

	// insideCell[i][j] = true if cell (i,j) is inside the polygon
	insideCell := make([][]bool, numCellsX)
	for i := range insideCell {
		insideCell[i] = make([]bool, numCellsY)
	}

	for i := 0; i < numCellsX; i++ {
		for j := 0; j < numCellsY; j++ {
			// Center of this cell
			cx := (xCoords[i] + xCoords[i+1]) / 2.0
			cy := (yCoords[j] + yCoords[j+1]) / 2.0
			insideCell[i][j] = isInsidePolygon(cx, cy, edges)
		}
	}

	// Now check all pairs of red tiles
	maxArea := 0

	for i, t1 := range redTiles[:len(redTiles)-1] {
		for _, t2 := range redTiles[i+1:] {
			width := abs(t1.X-t2.X) + 1
			height := abs(t1.Y-t2.Y) + 1
			area := width * height

			if area <= maxArea {
				continue
			}

			// Get rectangle bounds
			rx1, rx2 := min(t1.X, t2.X), max(t1.X, t2.X)
			ry1, ry2 := min(t1.Y, t2.Y), max(t1.Y, t2.Y)

			// Find cell index range for this rectangle
			ix1, ix2 := xIndex[rx1], xIndex[rx2]
			iy1, iy2 := yIndex[ry1], yIndex[ry2]

			// Check all cells that this rectangle covers
			valid := true
			for ci := ix1; ci < ix2 && valid; ci++ {
				for cj := iy1; cj < iy2; cj++ {
					if !insideCell[ci][cj] {
						valid = false
						break
					}
				}
			}

			if valid {
				maxArea = area
			}
		}
	}

	fmt.Println(maxArea)
}

// Ray casting algorithm: count intersections with vertical edges
// going rightward from point (px, py)
func isInsidePolygon(px, py int, edges []Edge) bool {
	count := 0
	for _, e := range edges {
		if !e.vertical {
			continue
		}
		// Vertical edge at x = e.x1
		ex := e.x1
		ey1, ey2 := min(e.y1, e.y2), max(e.y1, e.y2)

		// Edge must be to the right of point
		if ex <= px {
			continue
		}

		// Point's y must be within edge's y range (exclusive of endpoints to handle corners)
		if py > ey1 && py < ey2 {
			count++
		}
	}
	return count%2 == 1
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
