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
		line := scanner.Text()
		direction := line[0]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			dialPosition = (dialPosition - number) % 100
		} else if direction == 'R' {
			dialPosition = (dialPosition + number) % 100
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
