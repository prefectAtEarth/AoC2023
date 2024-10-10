package day03

import (
	"reflect"
	"testing"
)

var testdatapath = "../data/day03_test.txt"
var realdatapath = "../data/day03.txt"

var expectedAllParts = []PartNumber{
	{0, 0, 2, 467}, {0, 5, 7, 114},
	{2, 2, 3, 35}, {2, 6, 8, 633},
	{4, 0, 2, 617},
	{5, 7, 8, 58},
	{6, 2, 4, 592},
	{7, 6, 8, 755},
	{9, 1, 3, 664}, {9, 5, 7, 598},
}

var expectedMarker = []PartMarker{
	{1, 3, "*"},
	{3, 6, "#"},
	{4, 3, "*"},
	{5, 5, "+"},
	{8, 3, "$"},
	{8, 5, "*"},
}

func TestReadParts(t *testing.T) {
	lines := ReadLines(testdatapath)
	parts, _ := ReadParts(lines)
	for index, part := range parts {
		if !reflect.DeepEqual(part, expectedAllParts[index]) {
			t.Fatalf("Part content wrong.\nWanted: %v\nGot: %v", expectedAllParts, parts)
		}
	}
	if !reflect.DeepEqual(parts, expectedAllParts) {
		t.Fatalf("Part content wrong.\nWanted: %v\nGot: %v", expectedAllParts, parts)
	}
}

func TestMarker(t *testing.T) {
	lines := ReadLines(testdatapath)
	_, marker := ReadParts(lines)
	if !reflect.DeepEqual(marker, expectedMarker) {
		t.Fatalf("Marker content wrong.\nWanted: %v\nGot: %v", expectedMarker, marker)
	}
	for index, mark := range marker {
		println("Index:", index, "Marker:", mark.value, mark.lineNumber, mark.index)
		if !reflect.DeepEqual(mark, expectedMarker[index]) {
			t.Fatalf("Marker content wrong.\nWanted: %v\nGot: %v", expectedMarker[index], mark)
		}
	}
}

func TestSymbols(t *testing.T) {
	arr := []rune(".*#")
	dot := arr[0]
	asterisk := arr[1]
	hash := arr[2]
	if IsDigit(dot) {
		t.Fatalf(" . should not be a digit")
	}
	if IsDigit(asterisk) {
		t.Fatalf(" * should not be a digit")
	}
	if IsDigit(hash) {
		t.Fatalf(" # should not be a digit")
	}
	if IsPartMarker(dot) {
		t.Fatalf(" . should not be a part marker")
	}
	if !IsPartMarker(asterisk) {
		t.Fatalf(" * should be a part marker")
	}
	if !IsPartMarker(hash) {
		t.Fatalf(" # should be a part marker")
	}
}

func TestDay03MainTest(t *testing.T) {
	sum := Day03Main(testdatapath)
	if sum != 4361 {
		t.Fatalf("Expected Sum: 4361, got %d", sum)
	}
}

func TestDay03Main(t *testing.T) {
	lines := ReadLines(realdatapath)
	_, marker := ReadParts(lines)
	for index, mark := range marker {
		println("Index:", index, "Marker:", mark.value, mark.lineNumber, mark.index)
	}
	actualsum := Day03Main(realdatapath)
	if actualsum != 554003 {
		t.Fatalf("Expected Sum: 554003, got %d", actualsum)
	}
}
