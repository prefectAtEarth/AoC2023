package day03

import (
	"aoc2023/stuff"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	value      int
	startIndex int
	endIndex   int
	rowIndex   int
}

type Symbol struct {
	value      string
	startIndex int
	rowIndex   int
}

type Schematic struct {
	Numbers []Number
	Symbols []Symbol
}

type IndexRange struct {
	start int
	end   int
}

func (schematic *Schematic) addPart(number Number) {
	schematic.Numbers = append(schematic.Numbers, number)
}

func (schematic *Schematic) addSymbol(symbol Symbol) {
	schematic.Symbols = append(schematic.Symbols, symbol)
}

func ParseFileToSchematic(filepath string) Schematic {
	err, rows := stuff.ScanFile(filepath)
	schematic := Schematic{}

	if err != nil {
		log.Fatalf("Error reading file %s\n", filepath)
	}

	for rowIndex, row := range rows {
		numStart := -1
		numstrings := []string{}
		for index, char := range row {
			if unicode.IsNumber(char) {
				if numStart < 0 {
					numStart = index
				}
				numstrings = append(numstrings, string(char))
			} else {
				if string(char) != "." {
					schematic.addSymbol(Symbol{string(char), index, rowIndex})
				}
				if len(numstrings) > 0 {
					stringNum, _ := strconv.Atoi(strings.Join(numstrings[:], ""))
					schematic.addPart(Number{stringNum, numStart, numStart + len(numstrings) - 1, rowIndex})
					numstrings = nil
					numStart = -1
				}
			}
		}
	}
	return schematic
}

func HasSymbolicNeighbour(schematic Schematic) []int {
	var partNumsWithNeighbours = []int{}
	for _, part := range schematic.Numbers {
		for _, sym := range schematic.Symbols {
			partFence := IndexRange{part.startIndex - 1, part.endIndex + 1}
			if sym.rowIndex-1 == part.rowIndex || sym.rowIndex == part.rowIndex || sym.rowIndex+1 == part.rowIndex {
				hasNeighbour := sym.startIndex >= partFence.start && sym.startIndex <= partFence.end
				if hasNeighbour {
					partNumsWithNeighbours = append(partNumsWithNeighbours, part.value)
					break
				}
			}
		}
	}
	return partNumsWithNeighbours
}

func CalculatePartNumSum(partNumbers []int) int {
	var sum int
	for _, partNum := range partNumbers {
		sum += partNum
	}
	return sum
}
