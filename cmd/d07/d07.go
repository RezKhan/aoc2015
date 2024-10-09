package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func valueLeftRight(operations map[string]string, operation string, signals map[string]uint16) uint16 {
	t, err := strconv.Atoi(operation)
	if err != nil {
		_, ok := signals[operation]
		if !ok {
			signals = parseOperationsIntoSignals(operations, operation, signals)
		}
		return signals[operation]
	} else {
		return uint16(t)
	}
}

func parseOperationsIntoSignals(operations map[string]string, target string, signals map[string]uint16) map[string]uint16 {
	operation := strings.Split(operations[target], " ")
	fmt.Println(operations[target], " to ", target)
	// BASE CASE
	if len(operation) == 1 {
		tmp, err := strconv.Atoi(operation[0])
		if err != nil {
			// fmt.Println(err)
			signals = parseOperationsIntoSignals(operations, operation[0], signals)
			return signals
		}
		signals[target] = uint16(tmp)
	}
	// RECURSIVE CASES
	if len(operation) == 2 {
		if operation[0] == "NOT" {
			_, ok := signals[operation[1]]
			if !ok {
				signals = parseOperationsIntoSignals(operations, operation[1], signals)
			}
			x := uint16(signals[operation[1]])
			signals[target] = ^x
		}
	}

	if len(operation) == 3 {
		// LEFT or RIGHT side can be a number, check for both and assign
		left := valueLeftRight(operations, operation[0], signals)
		right := valueLeftRight(operations, operation[2], signals)
		switch operation[1] {
		case "AND":
			signals[target] = left & right
			break
		case "OR":
			signals[target] = left | right
			break
		case "LSHIFT":
			signals[target] = left << right
			break
		case "RSHIFT":
			signals[target] = left >> right
		}
	}
	return signals
}

func partOne(lines []string) {
	signals := make(map[string]uint16)
	operations := make(map[string]string)
	for _, line := range lines {
		tmpline := strings.Split(line, "->")
		tsource := strings.TrimRight(tmpline[0], " ")
		destination := strings.TrimSpace(tmpline[1])
		operations[destination] = tsource
	}
	for k := range operations {
		signals = parseOperationsIntoSignals(operations, k, signals)
	}
	target := "a"
	for k, v := range signals {
		fmt.Println(k, ": ", v)
	}
	fmt.Println("Target value for:", target, " = ", signals[target])
}

func main() {
	// fmt.Println(123 & 456)
	// lines := rf.ReadFile("./test.txt")
	lines := rf.ReadFile("./input.txt")
	partOne(lines)
}
