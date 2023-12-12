package day02

import (
	"aoc2023/stuff"
	"log"
	"strconv"
	"strings"
	"fmt"
)

type CubeCounts = struct {
	red   int
	green int
	blue  int
}

var maxCubes = CubeCounts{12, 13, 14}

type Game = struct {
	id    int
	Pulls []Pull
}

type Pull = struct {
	red   int
	green int
	blue  int
}

func ReadGamesFromFile(fileName string) []Game {
	err, gameStrings := stuff.ScanFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file %s", fileName)
	}
	games := []Game{}
	for _, gameString := range gameStrings {
		games = append(games, ParseGameFromString(gameString))
	}

	return games
}

func ParseGameFromString(gameString string) Game {
	idSplit := strings.Split(gameString, ": ")
	pullsSplit := strings.Split(idSplit[1], "; ")

	pulls := []Pull{}
	for _, pull := range pullsSplit {
		pullSplit := strings.Split(pull, ", ")
		var red, green, blue int = 0, 0, 0
		for _, colour := range pullSplit {
			colourSplit := strings.Split(colour, " ")
			switch colourSplit[1] {
			case "red":
				red, _ = strconv.Atoi(colourSplit[0])
			case "green":
				green, _ = strconv.Atoi(colourSplit[0])
			case "blue":
				blue, _ = strconv.Atoi(colourSplit[0])
			}
		}
		pulls = append(pulls, Pull{red, green, blue})
	}
	gameId, _ := strconv.Atoi(strings.Split(idSplit[0], " ")[1])
	game := Game{gameId, pulls}
	return game
}

func CheckGamePossibility(game Game) bool {
	for _, pull := range game.Pulls {
		if pull.red > maxCubes.red {
			return false
		}
		if pull.green > maxCubes.green {
			return false
		}
		if pull.blue > maxCubes.blue {
			return false
		}
	}
	return true
}

func GetMinCubesPerColourPerGame(game Game) CubeCounts {
	var red, green, blue int
	for _, pull := range game.Pulls {
		if red < pull.red {
			red = pull.red
		}
		if green < pull.green {
			green = pull.green
		}
		if blue < pull.blue {
			blue = pull.blue
		}
	}
	return CubeCounts{red, green, blue}
}

func SumPowerOfMinCube(games []Game) int {
	var cubeCounts []CubeCounts
	for _,game := range games {
		cubeCounts = append(cubeCounts, GetMinCubesPerColourPerGame(game))
	}
	sum := 0
	for _,minCube := range cubeCounts {
		minCubePower := minCube.red * minCube.green * minCube.blue
		sum += minCubePower
	}
	return sum
}

func SumPossibleGameIdsFromFile(fileName string) int {
	games := ReadGamesFromFile(fileName)
	var idSum int
	for _, game := range games {
		if CheckGamePossibility(game) {
			idSum += game.id
		}
	}
	return idSum
}

func Day02Main(filePath string) {
	idSum := SumPossibleGameIdsFromFile(filePath)
	fmt.Println("Day02: ID SUM: ", idSum)
	minCubePowerSum := SumPowerOfMinCube(ReadGamesFromFile(filePath))
	fmt.Println("Day02: Sum Cube Power: ", minCubePowerSum)
}