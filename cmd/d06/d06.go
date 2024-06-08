package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Light struct {
	on     bool
	coords Coords
}

type Command struct {
	command    string
	startRange Coords
	endRange   Coords
}

func setupGrid() [][]Light {
	var lightGrid [][]Light
	for i := 0; i < 1000; i++ {
		var lightLine []Light
		for j := 0; j < 1000; j++ {
			light := Light{false, Coords{j, i}}
			lightLine = append(lightLine, light)
		}
		lightGrid = append(lightGrid, lightLine)
	}
	return lightGrid
}

func partOne(lines []string) {
	lightGrid := setupGrid()
	for _, line := range lines {
		// tstr := string(line[:strings.LastIndex(line[:strings.Index(line, ",")], " ")])
		// tstr := strings.Split(line[strings.LastIndex(line[:strings.Index(line, ",")], " ")+1:], " ")
		tstr := strings.Split(line, " ")
		if tstr[0] == "turn" {
			// _, tstr = tstr[0], tstr[1:]
			tstr = tstr[1:]
		}
		fmt.Println(len(tstr), tstr)
		var tcmd Command
		tcmd.command = tstr[0]
		start := strings.Split(tstr[1], ",")
		end := strings.Split(tstr[3], ",")
		tcmd.startRange.x, _ = strconv.Atoi(start[0])
		tcmd.startRange.y, _ = strconv.Atoi(start[1])
		tcmd.endRange.x, _ = strconv.Atoi(end[0])
		tcmd.endRange.y, _ = strconv.Atoi(end[1])
		for y := tcmd.startRange.y; y <= tcmd.endRange.y; y++ {
			for x := tcmd.startRange.x; x <= tcmd.endRange.x; x++ {
				if tcmd.command == "on" {
					lightGrid[y][x].on = true
				}
				if tcmd.command == "off" {
					lightGrid[y][x].on = false
				}
				if tcmd.command == "toggle" {
					lightGrid[y][x].on = !lightGrid[y][x].on
				}
			}
		}
	}
	counter := 0
	for _, lline := range lightGrid {
		for _, light := range lline {
			if light.on {
				counter++
			}
		}
	}
	fmt.Println(counter)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
}
