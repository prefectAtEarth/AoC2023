package day03

import (
	"aoc2023/stuff"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type PartRow []PartNumber

type SymbolRow []PartSymbol

type PartNumber struct {
	value      int
	startIndex int
	endIndex   int
}

type PartSymbol struct {
	value      string
	startIndex int
}

type Schematic struct {
	PartRows   []PartRow
	SymbolRows []SymbolRow
}

func (schematic *Schematic) addPartRow(partRow PartRow) {
	schematic.PartRows = append(schematic.PartRows, partRow)
}

func (schematic *Schematic) addSymbolRow(symbolRow SymbolRow) {
	schematic.SymbolRows = append(schematic.SymbolRows, symbolRow)
}

func ParseFileToSchematic(filepath string) Schematic {
	err, rows := stuff.ScanFile(filepath)
	schematic := Schematic{}

	if err != nil {
		log.Fatalf("Error reading file %s\n", filepath)
	}

	for _, row := range rows {
		partNums := []PartNumber{}
		partSyms := []PartSymbol{}
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
					partSyms = append(partSyms, PartSymbol{string(char), index})
				}
				if len(numstrings) > 0 {
					stringNum, _ := strconv.Atoi(strings.Join(numstrings[:], ""))
					partNums = append(partNums, PartNumber{stringNum, numStart, numStart + len(numstrings) - 1})
					numstrings = nil
					numStart = -1
				}
			}
		}

		schematic.addPartRow(partNums)
		schematic.addSymbolRow(partSyms)
	}
	return schematic
}
