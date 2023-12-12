package day01

import "testing"

type data = struct {
	input  string
	output int
}

var testData = []data{
	data{"two1nine", 29},
	data{"eightwothree", 83},
	data{"abcone2threexyz", 13},
	data{"xtwone3four", 24},
	data{"4nineeightseven2", 42},
	data{"zoneight234", 14},
	data{"7pqrstsixteen", 76},
}

var expectedResult = 281

func TestExtractFirstAndLastIntegers(t *testing.T) {
	for _, test := range testData {
		if result := ExtractFirstAndLastIntegers(test.input); result != test.output {
			t.Errorf("Expected: %d, got: %d", test.output, result)
		}
	}
}

func TestCalcSumOfInts(t *testing.T) {
	var sum int
	var ints []int
	for _, test := range testData {
		ints = append(ints, test.output)
	}
	if sum = CalcSumOfInts(ints); sum != expectedResult {
		t.Errorf("Expected: %d, got: %d", expectedResult, sum)
	}
}
