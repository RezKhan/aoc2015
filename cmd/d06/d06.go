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
	on         bool
	brightness int
	coords     Coords
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
			light := Light{false, 0, Coords{j, i}}
			lightLine = append(lightLine, light)
		}
		lightGrid = append(lightGrid, lightLine)
	}
	return lightGrid
}

func parseLineToCommand(line string) Command {
	tstr := strings.Split(strings.ReplaceAll(line, " ", ","), ",")
	if tstr[0] == "turn" {
		tstr = tstr[1:]
	}
	var cmd Command
	cmd.command = tstr[0]
	cmd.startRange.x, _ = strconv.Atoi(tstr[1])
	cmd.startRange.y, _ = strconv.Atoi(tstr[2])
	cmd.endRange.x, _ = strconv.Atoi(tstr[4])
	cmd.endRange.y, _ = strconv.Atoi(tstr[5])

	return cmd
}

func partOne(lines []string) {
	lightGrid := setupGrid()
	for _, line := range lines {
		// tstr := string(line[:strings.LastIndex(line[:strings.Index(line, ",")], " ")])
		// tstr := strings.Split(line[strings.LastIndex(line[:strings.Index(line, ",")], " ")+1:], " ")
		cmd := parseLineToCommand(line)
		for y := cmd.startRange.y; y <= cmd.endRange.y; y++ {
			for x := cmd.startRange.x; x <= cmd.endRange.x; x++ {
				if cmd.command == "on" {
					lightGrid[y][x].on = true
				}
				if cmd.command == "off" {
					lightGrid[y][x].on = false
				}
				if cmd.command == "toggle" {
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

func partTwo(lines []string) {
	lightGrid := setupGrid()
	for _, line := range lines {
		cmd := parseLineToCommand(line)
		for y := cmd.startRange.y; y <= cmd.endRange.y; y++ {
			for x := cmd.startRange.x; x <= cmd.endRange.x; x++ {
				if cmd.command == "on" {
					lightGrid[y][x].brightness++
				}
				if cmd.command == "off" && lightGrid[y][x].brightness > 0 {
					lightGrid[y][x].brightness--
				}
				if cmd.command == "toggle" {
					lightGrid[y][x].brightness += 2
				}
			}
		}
	}
	totalBrightness := 0
	for _, lline := range lightGrid {
		for _, light := range lline {
			totalBrightness += light.brightness
		}
	}
	fmt.Println(totalBrightness)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	// lines := rf.ReadFile("test2.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
