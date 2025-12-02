package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isInvalid checks if a number is made of a sequence repeated at least twice
// e.g., 123123 (123 x2), 123123123 (123 x3), 1111111 (1 x7)
func isInvalid(n int64) bool {
	s := strconv.FormatInt(n, 10)
	length := len(s)

	// Try all possible pattern lengths that could repeat at least twice
	// Pattern length must divide the total length evenly
	// and be at most half the length (to repeat at least twice)
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		isRepeated := true

		// Check if the entire string is this pattern repeated
		for i := patternLen; i < length; i += patternLen {
			if s[i:i+patternLen] != pattern {
				isRepeated = false
				break
			}
		}

		if isRepeated {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	// Remove trailing comma if present and split by comma
	line = strings.TrimSuffix(line, ",")
	ranges := strings.Split(line, ",")

	var sum int64
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		for n := start; n <= end; n++ {
			if isInvalid(n) {
				sum += n
			}
		}
	}

	fmt.Println(sum)
}
