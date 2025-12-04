package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	totalOutput := 0
	maxOutputStrBuilder := strings.Builder{}
	maxOutputStrBuilder.Grow(2)

	file, err := os.Open("3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()

		searchStart := 0
		for digitsRemaining := 2; digitsRemaining > 0; digitsRemaining-- {
			maxIdx := searchStart
			maxByte := bank[searchStart]

			for y := searchStart + 1; y <= len(bank)-digitsRemaining; y++ {
				if bank[y] > maxByte {
					maxByte = bank[y]
					maxIdx = y
				}
			}

			maxOutputStrBuilder.WriteByte(maxByte)
			searchStart = maxIdx + 1
		}

		maxOutput, err := strconv.Atoi(maxOutputStrBuilder.String())
		if err != nil {
			log.Println("Failed to convert to int error:", err)
			continue
		}

		maxOutputStrBuilder.Reset()
		totalOutput += maxOutput
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Result
	println(totalOutput)
}
