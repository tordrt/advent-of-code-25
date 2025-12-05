package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type freshRange struct {
	start, end int
}

func main() {
	file, err := os.Open("5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ranges []freshRange

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		values := strings.Split(line, "-")
		start, err := strconv.Atoi(values[0])
		end, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Failed to convert integers in range: " + err.Error())
			continue
		}

		ranges = append(ranges, freshRange{
			start: start,
			end:   end,
		})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sort by start position
	slices.SortFunc(ranges, func(a, b freshRange) int {
		return a.start - b.start
	})

	// Single-pass merge
	merged := []freshRange{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.start <= last.end+1 { // overlaps or adjacent
			last.end = max(last.end, r.end)
		} else {
			merged = append(merged, r)
		}
	}

	// Count total fresh IDs
	totalFreshIds := 0
	for _, r := range merged {
		totalFreshIds += r.end + 1 - r.start
	}

	fmt.Println(totalFreshIds)
}
