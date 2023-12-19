package day03

import (
	"reflect"
	"testing"
)

var schematicSum = 4361
var partNums = []int{467, 35, 633, 617, 592, 755, 664, 598}
var notPartNums = []int{114, 58}
var expectedRows = []PartRow{
	{{467, 0, 2}, {114, 5, 7}},
	{},
	{{35, 2, 3}, {633, 6, 8}},
	{},
	{{617, 0, 2}},
	{{58, 7, 8}},
	{{592, 2, 4}},
	{{755, 6, 8}},
	{},
	{{664, 1, 3}, {598, 5, 7}},
}
var expectedSymbols = []SymbolRow{
	{},
	{{"*", 3}},
	{},
	{{"#", 6}},
	{{"*", 3}},
	{{"+", 5}},
	{},
	{},
	{{"$", 3}, {"*", 5}},
	{},
}

func TestParseFileToSchematic(t *testing.T) {
	schematic := ParseFileToSchematic("../../AoC2023_data/day03/testdata.txt")
	if !reflect.DeepEqual(schematic.PartRows, expectedRows) {
		t.Fatalf("Error:\nWanted:\t%v\nGot:\t%v\n", expectedRows, schematic.PartRows)
	}
	if !reflect.DeepEqual(schematic.SymbolRows, expectedSymbols) {
		t.Fatalf("Error:\nWanted:\t%v\nGot:\t%v\n", expectedSymbols, schematic.SymbolRows)
	}
}
