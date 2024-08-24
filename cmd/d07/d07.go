package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func parseCommandsIntoSignals(cmd map[string]string, target string, signals map[string]uint16) map[string]uint16 {
	fmt.Println("Entering target: ", target, "=", cmd[target])
	if target == "" {
		fmt.Println("You broke it")
		return nil
	}
	strcmd := strings.Split(cmd[target], " ")
	// fmt.Println(len(strcmd), strcmd[0])

	cmd0, cerr := strconv.Atoi(strcmd[0])
	// BASE CASE
	if len(strcmd) == 1 && cerr == nil {
		signals[target] = uint16(cmd0)
		return signals
	}
	// RECURSIVE CASES
	// The LEFT or RIGHT side may contain a number, and this needs to be checked for
	// Check if key exists in signals if not, go recursive
	if strcmd[0] != "NOT" && cerr != nil { // adding this nil check lets it get passed a couple of cases
		// needs more work
		_, ok := signals[target]
		if !ok {
			signals = parseCommandsIntoSignals(cmd, strcmd[0], signals)
		}
	}

	if strcmd[0] == "NOT" {
		_, ok := signals[strcmd[1]]
		if !ok {
			signals = parseCommandsIntoSignals(cmd, strcmd[1], signals)
		}
	}

	if strcmd[1] == "AND" || strcmd[1] == "OR" {
		_, ok := signals[strcmd[2]]
		if !ok {
			signals = parseCommandsIntoSignals(cmd, strcmd[2], signals)
		}
	}

	// ASSIGN SIGNALS TO WIRES
	num := 0
	if len(strcmd) == 3 && strcmd[1] == "SHIFT" {
		n, err := strconv.Atoi(strcmd[2])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		num = n
	}
	if strcmd[1] == "LSHIFT" {
		signals[target] = signals[strcmd[0]] << num
		return signals
	}
	if strcmd[1] == "RSHIFT" {
		signals[target] = signals[strcmd[0]] >> num
		return signals
	}
	if strcmd[1] == "AND" {
		signals[target] = signals[strcmd[0]] & signals[strcmd[2]]
		return signals
	}
	if strcmd[1] == "OR" {
		signals[target] = signals[strcmd[0]] | signals[strcmd[2]]
	}
	if strcmd[0] == "NOT" {
		x := uint16(signals[strcmd[1]])
		signals[target] = ^x
		return signals
	}
	return signals
}

func partOne(lines []string) {
	signals := make(map[string]uint16)
	commands := make(map[string]string)
	for _, line := range lines {
		tline := strings.Split(line, "->")
		tsource := strings.TrimRight(tline[0], " ")
		destination := strings.TrimSpace(tline[1])
		commands[destination] = tsource
	}
	target := "a"
	// for k, v := range commands {
	// 	fmt.Println("command:", k, " - has value: ", v)
	// }
	signals = parseCommandsIntoSignals(commands, target, signals)
	fmt.Println("Target value for:", target, " = ", signals[target])
}

func main() {
	// fmt.Println(123 & 456)
	// lines := rf.ReadFile("./test.txt")
	lines := rf.ReadFile("./input.txt")
	partOne(lines)
}
