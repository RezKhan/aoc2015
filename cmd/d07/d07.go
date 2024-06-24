package main

import (
	"fmt"
	"log"
	rf "readfile"
	"strconv"
	"strings"
)

func parseCommandsIntoSignals(cmd map[string]string, target string, signals map[string]int) map[string]int {
	fmt.Println(target, cmd[target])
	strcmd := strings.Split(cmd[target], " ")
	// BASE CASE
	if len(strcmd) == 1 {
		n, err := strconv.Atoi(strcmd[0])
		if err != nil {
			return nil
		}
		signals[target] = n
		return signals
	}
	// SHIFT OPERATION:
	if len(strcmd) == 3 && strings.Contains(strcmd[1], "SHIFT") {
		// Handle shift operation
		// Check if the value to be shifted is already populated
		if v, ok := signals[strcmd[0]]; ok {
			log.Println("v, ok:", v, ok)
			shiftval, err := strconv.Atoi(strcmd[2])
			log.Println(shiftval)
			if err != nil {
				log.Fatalln(err, "Shift value could not be determined")
			}
			if strcmd[1] == "LSHIFT" {
				fmt.Println("Shifting", signals[strcmd[0]], "left")
				signals[target] = signals[strcmd[0]] >> shiftval
			}
			if strcmd[1] == "RSHIFT" {
				fmt.Println("Shifting", signals[strcmd[0]], "right")
				signals[target] = signals[strcmd[0]] << shiftval
			}
			fmt.Println(v)
		} else {
			// RECURSIVE CASE
			fmt.Println("v not found")
			signals = parseCommandsIntoSignals(cmd, strcmd[0], signals)
		}
		return signals
	}
	// AND | OR OPERATION
	if len(strcmd) == 3 && !strings.Contains(strcmd[1], "SHIFT") {
	}
	// NOT OPERATION
	if len(strcmd) == 2 {

	}
	return signals
}

func partOne(lines []string) {
	signals := make(map[string]int)
	commands := make(map[string]string)
	for _, line := range lines {
		tline := strings.Split(line, "->")
		tsource := strings.TrimRight(tline[0], " ")
		destination := strings.TrimSpace(tline[1])
		commands[destination] = tsource
	}
	target := "f"
	signals = parseCommandsIntoSignals(commands, target, signals)
	fmt.Println(signals)
}

func main() {
	// fmt.Println(123 & 456)
	lines := rf.ReadFile("./test.txt")
	partOne(lines)
}
