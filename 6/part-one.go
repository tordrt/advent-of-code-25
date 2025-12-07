package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	operators := strings.Fields(lines[len(lines)-1])

	numberLines := lines[:len(lines)-1]

	numbers := make([][]int, 0, len(numberLines))
	for i := range numberLines {
		row := make([]int, 0, len(operators))
		for _, s := range strings.Fields(lines[i]) {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		numbers = append(numbers, row)
	}

	total := 0
	for col, op := range operators {
		result := numbers[0][col]
		for row := 1; row < len(numbers); row++ {
			switch op {
			case "+":
				result += numbers[row][col]
			case "*":
				result *= numbers[row][col]
			}
		}
		total += result
	}

	fmt.Println(total)
}
