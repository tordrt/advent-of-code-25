package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type freshRange struct {
	start, end int
}

func main() {
	file, err := os.Open("5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ranges []freshRange
	var ingredients []int
	onRanges := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			onRanges = false
			continue
		}

		if onRanges {
			values := strings.Split(line, "-")
			if len(values) != 2 {
				log.Println("Invalid range format: " + line)
				continue
			}

			start, err := strconv.Atoi(values[0])
			if err != nil {
				log.Println("Failed to convert start: " + err.Error())
				continue
			}
			end, err := strconv.Atoi(values[1])
			if err != nil {
				log.Println("Failed to convert end: " + err.Error())
				continue
			}

			ranges = append(ranges, freshRange{
				start: start,
				end:   end,
			})
		} else {
			ingredient, err := strconv.Atoi(line)
			if err != nil {
				log.Println("Failed to convert ingredient integer string: " + err.Error())
				continue
			}
			ingredients = append(ingredients, ingredient)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	freshIngredientsCount := 0
	for _, ingredient := range ingredients {
		for _, rng := range ranges {
			if ingredient >= rng.start && ingredient <= rng.end {
				freshIngredientsCount++
				break
			}
		}
	}

	fmt.Println(freshIngredientsCount)
}
