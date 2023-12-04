package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []GameSet
}

type GameSet struct {
	red   int
	green int
	blue  int
}

func newGame(id int, gameSets []GameSet) Game {
	return Game{id: id, sets: gameSets}
}

func (g Game) isValidGame(maxRed, maxGreen, maxBlue int) bool {
	for _, set := range g.sets {
		if set.red > maxRed {
			return false
		}
		if set.green > maxGreen {
			return false
		}
		if set.blue > maxBlue {
			return false
		}
	}

	return true
}

func (g Game) minGameSet() GameSet {
	minRed, minGreen, minBlue := 0, 0, 0

	for _, set := range g.sets {
		if set.red > minRed {
			minRed = set.red
		}
		if set.green > minGreen {
			minGreen = set.green
		}
		if set.blue > minBlue {
			minBlue = set.blue
		}
	}

	return GameSet{red: minRed, green: minGreen, blue: minBlue}
}

func parseGames(line string) Game {
	fields := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Fields(fields[0])[1])
	sets := strings.Split(fields[1], ";")
	gameSets := make([]GameSet, 0, len(sets))
	for _, set := range sets {
		red, green, blue := parseGameSets(set)
		gameSets = append(gameSets, GameSet{red: red, green: green, blue: blue})
	}

	return newGame(gameId, gameSets)
}

func parseGameSets(set string) (red int, green int, blue int) {
	gameInfo := strings.Split(set, ",")
	for _, game := range gameInfo {
		data := strings.Fields(game)
		if data[1] == "red" {
			red, _ = strconv.Atoi(data[0])
		} else if data[1] == "blue" {
			blue, _ = strconv.Atoi(data[0])
		} else if data[1] == "green" {
			green, _ = strconv.Atoi(data[0])
		} else {
			log.Printf("Parsing an invalid token of %v", data[1])
		}
	}

	return red, green, blue
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please Input file to read\n")
	}

	solvePart1(os.Args[1])
	solvePart2(os.Args[1])
}

func solvePart1(file string) {
	lines := readLines(file)
	sum := 0
	for _, line := range lines {
		games := parseGames(line)
		if games.isValidGame(12, 13, 14) {
			sum += games.id
		}
	}

	fmt.Printf("The Solution for part 1 is %d\n", sum)
}

func solvePart2(file string) {
	lines := readLines(file)
	sum := 0
	for _, line := range lines {
		games := parseGames(line)
		minSet := games.minGameSet()
		power := minSet.red * minSet.green * minSet.blue
		sum += power
	}

	fmt.Printf("The Solution for part 2 is %d\n", sum)
}

func readLines(file string) []string {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatalf("Failed to read %s: error - %v\n", file, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
