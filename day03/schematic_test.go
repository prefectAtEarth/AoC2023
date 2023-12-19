package day03

import (
	"reflect"
	"testing"
)

var testdata = "../../AoC2023_data/day03/testdata.txt"
var schematicSum = 4361
var partNums = []int{467, 35, 633, 617, 592, 755, 664, 598}
var notPartNums = []int{114, 58}
var expectedRows = []Number{
	{467, 0, 2, 0},
	{114, 5, 7, 0},
	{35, 2, 3, 2},
	{633, 6, 8, 2},
	{617, 0, 2, 4},
	{58, 7, 8, 5},
	{592, 2, 4, 6},
	{755, 6, 8, 7},
	{664, 1, 3, 9},
	{598, 5, 7, 9},
}
var expectedSymbols = []Symbol{
	{"*", 3, 1},
	{"#", 6, 3},
	{"*", 3, 4},
	{"+", 5, 5},
	{"$", 3, 8},
	{"*", 5, 8},
}

func TestParseFileToSchematic(t *testing.T) {
	schematic := ParseFileToSchematic(testdata)
	if !reflect.DeepEqual(schematic.Numbers, expectedRows) {
		t.Fatalf("Error:\nWanted:\t%v\nGot:\t%v\n", expectedRows, schematic.Numbers)
	}
	if !reflect.DeepEqual(schematic.Symbols, expectedSymbols) {
		t.Fatalf("Error:\nWanted:\t%v\nGot:\t%v\n", expectedSymbols, schematic.Symbols)
	}
}

func TestHasSymbolicNeighbour(t *testing.T) {
	schematic := ParseFileToSchematic(testdata)
	got := HasSymbolicNeighbour(schematic)
	if !reflect.DeepEqual(got, partNums) {
		t.Fatalf("Error:\nWanted:\t%v\nGot:\t%v\n", partNums, got)
	}
}

func TestCalculatePartNumSum(t *testing.T) {
	got := CalculatePartNumSum(partNums)
	if got != schematicSum {
		t.Fatalf("Error:\nWanted:\t%v\nGot:\t%v\n", schematicSum, got)
	}
}
