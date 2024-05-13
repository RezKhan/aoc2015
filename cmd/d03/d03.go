package main

import (
	"fmt"
	rf "readfile"
)

type House struct {
	x      int
	y      int
	visits int
}

func partOne(lines []string) {
	var x, y int
	var houses []House
	for _, line := range lines {
		x = 0
		y = 0
		house := House{x, y, 1}
		houses = append(houses, house)
		for _, char := range line {
			switch string(char) {
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
				house.x = x
				house.y = y
				house.visits = 1
				houses = append(houses, house)
			}
		}
	}
	fmt.Println(len(houses) + 1)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
}
