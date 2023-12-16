package day03

import (
	"reflect"
	"testing"
)

var schematicSum = 4361
var partNums = []int{467, 35, 633, 617, 592, 755, 664, 598}
var notPartNums = []int{114, 58}
var expectedPartNums = []PartNumber{
	{467, 0, 2},
	{114, 5, 7},
	{35, 2, 3},
	{633, 6, 8},
	{617, 0, 2},
	{58, 7, 8},
	{592, 2, 4},
	{755, 6, 8},
	{664, 1, 3},
	{598, 5, 7},
}

func TestParseFileToPartRows(t *testing.T) {
	partNums := ParseFileToPartRows("../../AoC2023_data/day03/testdata.txt")
	if !reflect.DeepEqual(partNums, expectedPartNums) {
		t.Fatalf("drror:\nWanted: %v\nGot: %v\n", expectedPartNums, partNums)
	}
}
