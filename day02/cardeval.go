package day02

import (
	"aoc2023/stuff"
	"log"
	"strconv"
	"strings"
)

type cubeCounts = struct {
	red   int
	green int
	blue  int
}

var maxCubes = cubeCounts{12, 13, 14}

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
