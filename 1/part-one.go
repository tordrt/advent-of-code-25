package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	password := 0
	dialPosition := 50

	file, err := os.Open("1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		direction := instruction[0]
		ticks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			dialPosition = (dialPosition - ticks) % 100
		} else if direction == 'R' {
			dialPosition = (dialPosition + ticks) % 100
		} else {
			continue
		}

		if dialPosition == 0 {
			password += 1
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Result
	println(password)
}
