package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	position := 50
	password := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())
		if instruction == "" {
			continue
		}

		direction := instruction[0]
		ticks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			fmt.Println("Error parsing ticks:", err)
			continue
		}

		if direction == 'L' {
			if position > 0 && ticks >= position {
				password += (ticks-position)/100 + 1
			} else if position == 0 {
				password += ticks / 100
			}

			position = (position - ticks) % 100

			if position < 0 {
				position += 100
			}
		} else if direction == 'R' {
			if position > 0 {
				password += (ticks + position) / 100
			} else {
				password += ticks / 100
			}

			position = (position + ticks) % 100
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(password)
}
