package stuff

import "testing"

func TestScanFile(t* testing.T) {

	err, result := ScanFile("stuff_testdata.txt")
	wantLines := []string {
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