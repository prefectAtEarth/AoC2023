package day01

import (
	"fmt"
	"strconv"
	"strings"
)

var numWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ExtractFirstAndLastIntegers(row string) int {
	var all []string
	for index, entry := range row {
		if isDigit(entry) {
			all = append(all, string(entry))
		}
		if isDigit, number := isDigitWord(row[index:]); isDigit {
			all = append(all, number)
		}
	}

	var out []string
	if len(all) > 0 {
		out = append(out, all[0])
		out = append(out, all[len(all)-1])
	}
	
	result, err := strconv.Atoi(strings.Join(out, ""))
	if err != nil {
		fmt.Println("error converting out to joined string:", out)
		return 0
	}
	
	return result
}

func CalcSumOfInts(ints []int) int {
	var sum int
	for _, entry := range ints {
		sum += entry
	}
	return sum
}

func isDigit(c rune) bool {
	if _, err := strconv.Atoi(string(c)); err == nil {
		return true
	} else {
		return false
	}
}

func isDigitWord(row string) (bool, string) {
	for index, numWord := range numWords {
		if strings.Index(row, numWord) == 0 {
			return true, fmt.Sprint(index)
		}
	}
	return false, ""
}

func input(s []string, err error) []string {
	if err != nil {
		return s
	}
	var i string
	count, err := fmt.Scanln(&i)
	if count == 1 {
		s = append(s, i)
	}
	return input(s, err)
}

func main() {
	userInput := input([]string{}, nil)
	fmt.Println("Input: ", userInput)

	var ints []int
	for _, row := range userInput {
		ints = append(ints, ExtractFirstAndLastIntegers(row))
	}
	fmt.Println("RESULT: ", CalcSumOfInts(ints))
}
