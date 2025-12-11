package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("11/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Build the graph
	graph := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		device := parts[0]
		outputs := strings.Split(parts[1], " ")
		graph[device] = outputs
	}

	// Count all paths from "you" to "out" using DFS
	count := countPaths(graph, "you", "out")
	fmt.Println(count)
}

func countPaths(graph map[string][]string, current, target string) int {
	if current == target {
		return 1
	}

	outputs, exists := graph[current]
	if !exists {
		return 0
	}

	total := 0
	for _, next := range outputs {
		total += countPaths(graph, next, target)
	}
	return total
}
