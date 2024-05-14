package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	rf "readfile"
	"strconv"
)

func partOne(lines []string) {
	for _, line := range lines {
		found := false
		i := 0
		for found == false {
			tohashStr := line + strconv.Itoa(i)
			hash := md5.Sum([]byte(tohashStr))
			if hex.EncodeToString(hash[:])[:5] == "00000" {
				fmt.Println(i, line, hex.EncodeToString(hash[:]))
				found = true
			}
			i++
		}
	}
}

func partTwo(lines []string) {
	for _, line := range lines {
		found := false
		i := 0
		for found == false {
			tohashStr := line + strconv.Itoa(i)
			hash := md5.Sum([]byte(tohashStr))
			if hex.EncodeToString(hash[:])[:6] == "000000" {
				fmt.Println(i, line, hex.EncodeToString(hash[:]))
				found = true
			}
			i++
		}
	}
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
