package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func parseCommandsIntoSignals(cmd map[string]string, signals map[string]int) {
	// fmt.Println("elements in map: ", len(cmd))
	for k, str := range cmd {
		// fmt.Println(k, str, len(strings.Split(str, " ")))
		if len(strings.Split(str, " ")) == 1 {
			n, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			signals[k] = n
		}
	}
	fmt.Println(signals)
}

func partOne(lines []string) {
	signals := make(map[string]int)
	commands := make(map[string]string)
	for _, line := range lines {
		tline := strings.Split(line, "->")
		tsource := strings.TrimRight(tline[0], " ")
		destination := strings.TrimSpace(tline[1])
		commands[destination] = tsource
		// fmt.Println(len(tsource), tsource, destination)
	}
	// target := "e"
	parseCommandsIntoSignals(commands, signals)
}

func main() {
	// fmt.Println(123 & 456)
	lines := rf.ReadFile("./test.txt")
	partOne(lines)
}
