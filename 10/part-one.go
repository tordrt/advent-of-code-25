package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("10/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Parse the line
		target, buttons := parseLine(line)
		minPresses := findMinPresses(target, buttons)
		total += minPresses
	}

	fmt.Println(total)
}

func parseLine(line string) (uint64, []uint64) {
	// Find indicator light diagram in [...]
	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	diagram := line[start+1 : end]

	// Parse target state
	var target uint64
	for i, c := range diagram {
		if c == '#' {
			target |= 1 << i
		}
	}

	// Parse button wiring schematics in (...)
	var buttons []uint64
	rest := line[end+1:]

	for {
		pStart := strings.Index(rest, "(")
		if pStart == -1 {
			break
		}
		pEnd := strings.Index(rest, ")")
		buttonStr := rest[pStart+1 : pEnd]

		var button uint64
		if buttonStr != "" {
			parts := strings.Split(buttonStr, ",")
			for _, p := range parts {
				idx, _ := strconv.Atoi(p)
				button |= 1 << idx
			}
		}
		buttons = append(buttons, button)

		rest = rest[pEnd+1:]
		// Stop when we hit the joltage requirements
		if strings.Contains(rest[:min(len(rest), 2)], "{") {
			break
		}
	}

	return target, buttons
}

func findMinPresses(target uint64, buttons []uint64) int {
	n := len(buttons)

	// Try all 2^n combinations
	minPresses := n + 1 // impossible value
	for mask := 0; mask < (1 << n); mask++ {
		var state uint64
		presses := bits.OnesCount(uint(mask))

		// Early termination if we already found better
		if presses >= minPresses {
			continue
		}

		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				state ^= buttons[i]
			}
		}

		if state == target {
			minPresses = presses
		}
	}

	if minPresses == n+1 {
		// No solution found (shouldn't happen for valid input)
		return 0
	}
	return minPresses
}
