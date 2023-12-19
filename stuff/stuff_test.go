package stuff

import (
	"testing"
)

func TestScanFile(t *testing.T) {

	err, result := ScanFile("stuff_testdata.txt")
	wantLines := []string{
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		"no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy",
	}

	if err != nil {
		t.Fatalf("Error opening file: %s", err)
	}

	if result != nil && result[0] != wantLines[0] {
		t.Fatalf("Error reading file")
	}
}

func TestFindAdjacentNumbersInString(t *testing.T) {
	_, result := ScanFile("../../AoC2023_data/day03/testdata.txt")
	row1 := FindAdjacentNumbersInString(result[0])
	if row1[0] != 467 {
		t.Fatalf("row1[0] error:\nWanted: %d\nGot: %d\n", 467, row1[0])
	}
	if row1[1] != 114 {
		t.Fatalf("row1[1] error:\nWanted: %d\nGot: %d\n", 114, row1[0])
	}
}
