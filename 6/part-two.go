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
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	operatorLine := lines[len(lines)-1]
	numberLines := lines[:len(lines)-1]
	currentOperator := rune(operatorLine[0])

	currentSum := 0
	totalSum := 0
	for i, operator := range operatorLine {
		// New operator == new column
		if rune(operator) != ' ' {
			currentOperator = rune(operatorLine[i])
			totalSum += currentSum
			currentSum = 0
		}

		currentNum := make([]string, 0, len(numberLines))
		for x := range numberLines {
			num := numberLines[x][i]
			if rune(num) != ' ' {
				currentNum = append(currentNum, string(rune(num)))
			}
		}

		if len(currentNum) == 0 {
			continue
		}

		numInt, err := strconv.Atoi(strings.Join(currentNum, ""))
		if err != nil {
			fmt.Println("failed to convert string to int", strings.Join(currentNum, ""), err)
			continue
		}

		if currentSum == 0 {
			currentSum = numInt
			continue
		}

		switch currentOperator {
		case '+':
			currentSum += numInt
		case '*':
			currentSum *= numInt
		}

	}

	// Add the sum for the last column
	totalSum += currentSum

	fmt.Println(totalSum)
}
