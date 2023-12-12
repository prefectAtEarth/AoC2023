package main

import (
	"day02"
	"fmt"
	"os"
)

func main() {
	idSum := day02.SumPossibleGameIdsFromFile(os.Args[1])
	fmt.Println("ID SUM: ", idSum)
}