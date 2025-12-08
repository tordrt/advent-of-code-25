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

	pairIdx := 0
	for range 1000 {
		if pairIdx >= len(pairs) {
			break
		}
		p := pairs[pairIdx]
		pairIdx++

		boxA, boxB := boxes[p.i], boxes[p.j]
		aCircuit, aHasCircuit := circuits[boxA]
		bCircuit, bHasCircuit := circuits[boxB]

		// Already in same circuit - counts as a connection but nothing happens
		if aHasCircuit && bHasCircuit && aCircuit == bCircuit {
			continue
		}

		if aHasCircuit && bHasCircuit {
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
		continue
	}

	// Deduplicate circuits (multiple boxes point to the same circuit)
	seen := make(map[*[]JunctionBox]bool)
	var sizes []int
	for _, circuit := range circuits {
		if !seen[circuit] {
			seen[circuit] = true
			sizes = append(sizes, len(*circuit))
		}
	}

	// Find top 3
	largestA, largestB, largestC := 0, 0, 0
	for _, size := range sizes {
		if size > largestA {
			largestC = largestB
			largestB = largestA
			largestA = size
		} else if size > largestB {
			largestC = largestB
			largestB = size
		} else if size > largestC {
			largestC = size
		}
	}

	fmt.Println("Largest circuits:", largestA, largestB, largestC)
	fmt.Println("Result:", largestA*largestB*largestC)
}

func (b1 JunctionBox) DistanceToo(b2 JunctionBox) float64 {
	dx := b2.X - b1.X
	dy := b2.Y - b1.Y
	dz := b2.Z - b1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
