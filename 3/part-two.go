package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalOutput := 0
	maxOutputStrBuilder := strings.Builder{}
	maxOutputStrBuilder.Grow(12)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()

		searchStart := 0
		for digitsRemaining := 12; digitsRemaining > 0; digitsRemaining-- {
			maxIdx := searchStart
			maxByte := bank[searchStart]

			for i := searchStart + 1; i <= len(bank)-digitsRemaining; i++ {
				if bank[i] > maxByte {
					maxByte = bank[i]
					maxIdx = i
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
