package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isInvalid checks if a number is made of a sequence repeated twice
// e.g., 55 (5+5), 6464 (64+64), 123123 (123+123)
func isInvalid(n int64) bool {
	s := strconv.FormatInt(n, 10)
	// Must have even length to be a doubled sequence
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
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
