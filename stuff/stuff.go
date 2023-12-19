package stuff

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ScanFile(fileName string) (error, []string) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return err, nil
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	out := []string{}
	for sc.Scan() {
		out = append(out, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return err, nil
	}
	return nil, out
}

func FindAdjacentNumbersInString(input string) []int {
	var fullNums []int
	nums := []string{}
	for _, char := range input {
		if unicode.IsNumber(char) {
			nums = append(nums, string(char))
		} else {
			if len(nums) > 0 {
				stringNum, _ := strconv.Atoi(strings.Join(nums[:], ""))
				nums = nil
				fullNums = append(fullNums, stringNum)
			}
		}
	}
	return fullNums
}
