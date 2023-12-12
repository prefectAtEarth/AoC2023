package stuff

import (
	"os"
	"log"
	"bufio"
)

func ScanFile(fileName string) (error, []string) {
    f, err := os.OpenFile(fileName , os.O_RDONLY, os.ModePerm)
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