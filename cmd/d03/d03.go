package main

import (
	"fmt"
	rf "readfile"
)

type House struct {
	x         int
	y         int
	visits    int
	visitedby string
}

func parseChar(char string, x int, y int) (int, int) {
	switch char {
	case "<":
		x--
		break
	case ">":
		x++
		break
	case "^":
		y--
		break
	case "v":
		y++
		break
	}
	return x, y
}

func visitHouse(x int, y int, houses []House, visitedby string) []House {
	// check if there is a house at the current x / y, if there is not create a new house
	// if there is then increment visited at the index
	visited := false
	for i := range houses {
		if houses[i].x == x && houses[i].y == y {
			houses[i].visits++
			visited = true
			continue
		}
	}
	if !visited {
		house := House{x, y, 1, visitedby}
		houses = append(houses, house)
	}

	return houses
}

func partOne(lines []string) {
	var x, y int
	for _, line := range lines {
		var houses []House
		x = 0
		y = 0
		house := House{x, y, 1, "Santa"}
		houses = append(houses, house)
		for _, char := range line {
			x, y = parseChar(string(char), x, y)
			houses = visitHouse(x, y, houses, "Santa")
		}
		fmt.Println(len(houses))
	}
}

func partTwo(lines []string) {
	var x1, y1, x2, y2 int
	for _, line := range lines {
		var houses []House
		x1 = 0
		y1 = 0
		x2 = 0
		y2 = 0
		house := House{x1, y1, 1, "Santa"}
		houses = append(houses, house)
		house = House{x1, y1, 2, "RoboSanta"}
		houses = append(houses, house)
		for i := 0; i < len(line); i += 2 {
			x1, y1 = parseChar(string(line[i]), x1, y1)
			x2, y2 = parseChar(string(line[i+1]), x2, y2)
			houses = visitHouse(x1, y1, houses, "Santa")
			houses = visitHouse(x2, y2, houses, "RoboSanta")
		}
		fmt.Println(len(houses) - 1)
	}
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
