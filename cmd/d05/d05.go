package main

import (
	"fmt"
	rf "readfile"
	"regexp"
	"strings"
)

func partOne(lines []string) {
	// nice: regex for vowels (a|e|i|o|u) || regex for 2 or more ([a-z])\1{1,}
	reVowels := regexp.MustCompile(`(aeiou)`)
	// CURRENTLY UNSUPPORTED IN GO: reDoubles := regexp.MustCompile(`([a-z])\1{1,}`)
	// naughty: just do as array and strings.index ["ab", "cd", "pq", "xy"]
	naughtystrings := []string{"ab", "cd", "pq", "xy"}
	counter := 0

	for _, line := range lines {
		nice := true
		for i := 1; i < len(line); i++ {
			if line[i] != line[i-1] {
				nice = false
			} else {
				nice = true
				break
			}
		}
		if len(reVowels.FindAllString(line, -1)) < 3 {
			nice = false
		}
		for _, naughtystr := range naughtystrings {
			if strings.Contains(line, naughtystr) {
				nice = false
			}
		}
		if nice {
			counter++
		}
	}
	fmt.Println("P1 Nice strings: ", counter)
}

func dupCheck(line string) bool {
	for i := 0; i < len(line)-1; i++ { // len-1 because the chunk is +1
		chunk := line[i : i+2]
		reChunk := regexp.MustCompile(chunk)
		if len(reChunk.FindAllString(line, 2)) > 1 {
			return true
		} else {
			// fmt.Println(chunk)
			continue
		}
	}
	return false
}

func tripCheck(line string) bool {
	for i := 3; i < len(line); i++ {
		chunk := line[i-3 : i]
		if chunk[0] == chunk[2] {
			return true
		} else {
			continue
		}
	}
	return false
}

func partTwo(lines []string) {
	// this one seems to want to do sliding windows
	counter := 0
	nice := true
	for _, line := range lines {
		// nice = dupCheck(line)
		// nice = tripCheck(line)
		if dupCheck(line) && tripCheck(line) { // real check
			// if nice {
			counter++
			fmt.Println(line, nice)
		}
	}
	fmt.Println("P2 Nice strings: ", counter)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	// lines := rf.ReadFile("test2.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
