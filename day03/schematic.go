package day03

import (
	"aoc2023/stuff"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Row = struct {
	PartNumbers []PartNumber
}

type PartNumber = struct {
	value      int
	startIndex int
	endIndex   int
}

type PartSymbol = struct {
	value      string
	startIndex int
}

func ParseFileToPartRows(filepath string) []PartNumber {
	err, rows := stuff.ScanFile(filepath)

	if err != nil {
		log.Fatalf("Error reading file %s\n", filepath)
	}

	var partNums []PartNumber
	for _, row := range rows {
		start := -1
		numstrings := []string{}
		for index, char := range row {
			if unicode.IsNumber(char) {
				start = index
				numstrings = append(numstrings, string(char))
			} else {
				if len(numstrings) > 0 {
					stringNum, _ := strconv.Atoi(strings.Join(numstrings[:], ""))
					partNums = append(partNums, PartNumber{stringNum, start, start + len(numstrings) - 1})
					numstrings = nil
				}
			}
		}
	}
	return partNums
}
