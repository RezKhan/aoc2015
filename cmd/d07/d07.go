package main

import (
	"fmt"
	rf "readfile"
	"strings"
)

func partOne(lines []string) {
	for _, line := range lines {
		tline := strings.Split(line, "->")
		tsource := strings.Split(strings.TrimRight(tline[0], " "), " ")
		destination := strings.TrimSpace(tline[1])
		fmt.Println(len(tsource), tsource, destination)
	}
}

func main() {
	fmt.Println(123 & 456)
	lines := rf.ReadFile("./test.txt")
	partOne(lines)
}
