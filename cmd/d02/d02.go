package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func partOne(lines []string) {
	// We have a string per line with numbers representing the dimensions of a box separated by x e.g. 2x3x4
	// formula to get the are of each side is 2*l*w + 2*w*h 2*h*l i.e. 2x3x4 should be 52
	// additional slack is needed which is the area of the smallest side i.e. 2x3x4 has smallest side 2x3 = 6
	totalArea := 0
	for _, line := range lines {
		var nums []int
		strnums := strings.Split(line, "x")
		for _, snum := range strnums {
			n, err := strconv.Atoi(snum)
			if err != nil {
				fmt.Println(err)
				break
			}
			nums = append(nums, n)
		}
		minside := 9 << 31
		area := 0
		for i, numi := range nums {
			for j, numj := range nums {
				if i == j {
					continue
				}
				side := numi * numj
				if side < minside {
					minside = side
				}
				area += side
			}
		}
		area += minside
		totalArea += area
		fmt.Println("box: ", line, "area: ", area, "total: ", totalArea)
	}
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
}
