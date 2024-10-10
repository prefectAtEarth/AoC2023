package day03

import (
	"aoc2023/stuff"
	"log"
	"strconv"
)

type PartNumber struct {
	lineNumber int
	startIndex int
	endIndex   int
	value      int
}

type PartMarker struct {
	lineNumber int
	index      int
	value      string
}

func ReadLines(filename string) []string {
	err, lines := stuff.ScanFile(filename)
	if err != nil {
		log.Fatal("WHAAA")
	}
	return lines
}

func ReadParts(lines []string) ([]PartNumber, []PartMarker) {
	var parts []PartNumber
	var marker []PartMarker

	for lineNumber, line := range lines {

		start := false
		var currentPartNumberStart int
		var currentPartNumberEnd int
		currentPartNumber := ""

		for index, symbol := range line {
			if IsPartMarker(symbol) {
				marker = append(marker, PartMarker{lineNumber, index, string(symbol)})
			}
			if IsDigit(symbol) {
				if !start {
					currentPartNumberStart = index
				}
				start = true
				currentPartNumber = currentPartNumber + string(symbol)
			} else {
				if start {
					partValue, _ := strconv.Atoi(currentPartNumber)
					currentPartNumberEnd = index - 1
					partNumber := PartNumber{lineNumber, currentPartNumberStart, currentPartNumberEnd, partValue}
					parts = append(parts, partNumber)
					start = false
					currentPartNumber = ""
					currentPartNumberEnd = 0
					currentPartNumberStart = 0
				}
			}
		}
	}
	return parts, marker
}

func FilterParts(parts []PartNumber, lines []string) []PartNumber {
	var filteredParts []PartNumber
	for _, part := range parts {
		if part.hasNeighbourSymbol(part.lineNumber, lines) {
			filteredParts = append(filteredParts, part)
		}
	}
	return filteredParts
}

func SumParts(parts []PartNumber) int {
	var sum int
	for _, part := range parts {
		sum += part.value
	}
	return sum
}

func IsPartMarker(symbol rune) bool {
	if IsDigit(symbol) {
		return false
	} else if string(symbol) == "." {
		return false
	} else {
		return true
	}
}

func (partNumber PartNumber) hasNeighbourSymbol(currentLine int, lines []string) bool {
	if len(lines) < 0 {
		log.Fatalf("Run ReadLines first")
	}
	previousLine := getPreviousLine(currentLine, lines)
	var startIndex int
	if partNumber.startIndex-1 < 0 {
		startIndex = 0
	} else {
		startIndex = partNumber.startIndex - 1
	}
	var endIndex int
	if partNumber.endIndex+1 > len(lines[currentLine]) {
		endIndex = len(lines[currentLine]) - 1
	} else {
		endIndex = partNumber.endIndex + 1
	}
	if len(previousLine) > 0 {
		for _, symbol := range previousLine[startIndex : endIndex+1] {
			if IsPartMarker(symbol) {
				return true
			}
		}
	}
	followingLine := getFollowingLine(currentLine, lines)
	if len(followingLine) > 0 {
		for _, symbol := range followingLine[startIndex : endIndex+1] {
			if IsPartMarker(symbol) {
				return true
			}
		}
	}
	for _, symbol := range lines[currentLine][startIndex : endIndex+1] {
		if IsPartMarker(symbol) {
			return true
		}
	}
	return false
}

func getPreviousLine(linenumber int, lines []string) string {
	if (linenumber - 1) < 0 {
		return ""
	} else {
		return lines[linenumber-1]
	}
}

func getFollowingLine(linenumber int, lines []string) string {
	if (linenumber + 1) > (len(lines) - 1) {
		return ""
	} else {
		return lines[linenumber+1]
	}
}

func IsDigit(c rune) bool {
	if _, err := strconv.Atoi(string(c)); err == nil {
		return true
	} else {
		return false
	}
}

func Day03Main(filepath string) int {
	lines := ReadLines(filepath)
	parts, _ := ReadParts(lines)
	filtered := FilterParts(parts, lines)
	sum := SumParts(filtered)
	println("PartSum:", sum)
	return sum
}
