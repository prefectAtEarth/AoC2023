package day02

import (
	"aoc2023/stuff"
	"reflect"
	"testing"
)

var testdatapath = "../../AoC2023_data/day02/testdata.txt"

var expectedGames = []Game{
	{1, []Pull{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}},
	{2, []Pull{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}},
	{3, []Pull{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}}},
	{4, []Pull{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}}},
	{5, []Pull{{6, 3, 1}, {1, 2, 2}}},
}

var expectedMinCubes = []CubeCounts{
	{4, 2, 6},
	{1, 3, 4},
	{20, 13, 6},
	{14, 3, 15},
	{6, 3, 2},
}

var expectedPowerOfMinCubes = []int {48, 12, 1560, 630, 36}
var expectedSumOfPowerOfMinCubes = 2286

var expectedIdSum = 8

func TestReadGamesFromFile(t *testing.T) {
	games := ReadGamesFromFile(testdatapath)
	wantedGames := expectedGames

	if len(games) != len(wantedGames) {
		t.Fatalf("Games len wrong.\nWanted: %d\nGot: %d", len(wantedGames), len(games))
	}

	for index, entry := range wantedGames {
		game := games[index]
		if !reflect.DeepEqual(game, entry) {
			t.Fatalf("Games content wrong.\nWanted: %v\nGot: %v", entry, games[index])
		}
	}
}

func TestParseGameFromString(t *testing.T) {
	_, games := stuff.ScanFile(testdatapath)

	game := ParseGameFromString(games[0])

	if !reflect.DeepEqual(game, expectedGames[0]) {
		t.Fatalf("Error parsing gamestring to Game:\nWanted: %v\nGot: %v\n", expectedGames[0], game)
	}
}

func TestCheckGamePossibility(t *testing.T) {
	_, games := stuff.ScanFile(testdatapath)
	game1 := ParseGameFromString(games[0])
	game3 := ParseGameFromString(games[2])

	possibleGame := CheckGamePossibility(game1)
	impossibleGame := CheckGamePossibility(game3)
	if possibleGame == false {
		t.Fatalf("Game 1 should be possible")
	}
	if impossibleGame == true {
		t.Fatalf("Game 3 should be impossible")
	}
}

func TestGetMinCubesPerColourPerGame(t *testing.T) {
	_, games := stuff.ScanFile(testdatapath)
	for index,gameString := range games {
		game := ParseGameFromString(gameString)
		minCubes := GetMinCubesPerColourPerGame(game)
		if !reflect.DeepEqual(minCubes, expectedMinCubes[index]) {
			t.Fatalf("MinCubes wrong.\nWanted: %d\nGot: %d\n", expectedMinCubes[index], minCubes)
		}
	}	
}

func TestSumPowerOfMinCubesPerGame(t *testing.T) {
	games := ReadGamesFromFile(testdatapath)
	sum := SumPowerOfMinCube(games)

	if sum != expectedSumOfPowerOfMinCubes {
		t.Fatalf("Sum Power of MinCubes wrong.\nWanted: %d\nGot: %d\n", sum, expectedSumOfPowerOfMinCubes)
	}
}

func TestSumPossibleGameIds(t *testing.T) {
	idSum := SumPossibleGameIdsFromFile(testdatapath)

	if idSum != expectedIdSum {
		t.Fatalf("IdSum wrong: %d", idSum)
	}
}
