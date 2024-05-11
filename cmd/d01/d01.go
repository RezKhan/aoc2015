package main

import (
	"fmt"
	rf "readfile"
)

func partOne(lines []string) {
	for _, line := range lines {
		lineval := 0
		for _, char := range line {
			if string(char) == "(" {
				lineval++
			}
			if string(char) == ")" {
				lineval--
			}
		}
		fmt.Println(line, "=", lineval)
	}
}

func partTwo(lines []string) {
	for _, line := range lines {
		lineval := 0
		for i, char := range line {
			if string(char) == "(" {
				lineval++
			}
			if string(char) == ")" {
				lineval--
			}
			if lineval == -1 {
				fmt.Println(line[:i], "at", i+1)
				break
			}
		}
	}
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
