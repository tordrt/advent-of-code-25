package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	X, Y, Z float64
}

func main() {
	file, err := os.Open("8/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var boxes []JunctionBox
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		if len(line) != 3 {
			fmt.Println("Invalid line:", line)
			continue
		}
		x, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			panic(err)
		}
		z, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			panic(err)
		}

		boxes = append(boxes, JunctionBox{x, y, z})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	circuits := make(map[JunctionBox]*[]JunctionBox, 0)

	// Pre-compute all pairs and sort by distance
	type IndexPair struct {
		i, j     int
		distance float64
	}
	var pairs []IndexPair
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, IndexPair{i, j, boxes[i].DistanceToo(boxes[j])})
		}
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].distance < pairs[b].distance
	})

	var lastBoxA, lastBoxB JunctionBox

	for _, p := range pairs {
		boxA, boxB := boxes[p.i], boxes[p.j]
		aCircuit, aHasCircuit := circuits[boxA]
		bCircuit, bHasCircuit := circuits[boxB]

		// Already in same circuit - skip
		if aHasCircuit && bHasCircuit && aCircuit == bCircuit {
			continue
		}

		// Track this as a meaningful connection
		lastBoxA, lastBoxB = boxA, boxB

		if aHasCircuit && bHasCircuit {
			// Merge two circuits
			circuitJBs := make(map[JunctionBox]struct{}, 0)

			for _, box := range *aCircuit {
				circuitJBs[box] = struct{}{}
			}
			for _, box := range *bCircuit {
				circuitJBs[box] = struct{}{}
			}

			newMergedCircuit := make([]JunctionBox, 0)

			for box := range circuitJBs {
				newMergedCircuit = append(newMergedCircuit, box)
			}

			for _, box := range *aCircuit {
				circuits[box] = &newMergedCircuit
			}
			for _, box := range *bCircuit {
				circuits[box] = &newMergedCircuit
			}

			// Check if all boxes are now in one circuit
			if len(newMergedCircuit) == len(boxes) {
				break
			}
			continue
		}

		if bHasCircuit {
			*bCircuit = append(*bCircuit, boxA)
			circuits[boxA] = bCircuit
			continue
		}
		if aHasCircuit {
			*aCircuit = append(*aCircuit, boxB)
			circuits[boxB] = aCircuit
			continue
		}

		// None of them are a part of a circuit yet
		newCircuit := make([]JunctionBox, 0)
		newCircuit = append(newCircuit, boxA)
		newCircuit = append(newCircuit, boxB)
		circuits[boxA] = &newCircuit
		circuits[boxB] = &newCircuit
	}

	fmt.Printf("Last connection: (%.0f,%.0f,%.0f) and (%.0f,%.0f,%.0f)\n",
		lastBoxA.X, lastBoxA.Y, lastBoxA.Z, lastBoxB.X, lastBoxB.Y, lastBoxB.Z)
	fmt.Println("Result:", int(lastBoxA.X)*int(lastBoxB.X))
}

func (b1 JunctionBox) DistanceToo(b2 JunctionBox) float64 {
	dx := b2.X - b1.X
	dy := b2.Y - b1.Y
	dz := b2.Z - b1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
