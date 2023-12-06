package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	multiarray "github.com/chadsmith12/aoc-2023/multi_array"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please Input file to read\n")
	}
	solvePart1(os.Args[1])
}

func solvePart1(file string) {
	lines := readLines(file)
	grid := multiarray.New[rune](len(lines)+2, len(lines[0])+2)
	grid.FillRow(0, '.')
	grid.FillRow(grid.Rows()-1, '.')
	for i, line := range lines {
		for j, value := range line {
			if j == 0 {
				grid.Set(i+1, 0, '.')
			}
			grid.Set(i+1, j+1, value)
			if j == len(line)-1 {
				grid.Set(i+1, j+2, '.')
			}
		}
	}

	numbers := make([]int, 0)
	var numberBuilder strings.Builder
	for row := 0; row < grid.Rows(); row++ {
		for col := 0; col < grid.Cols(); col++ {
			// not a number anymore, need to see if we built one out
			if !unicode.IsDigit(grid.Get(row, col)) {
				// we weren't building a number out, we can't just continue
				if numberBuilder.Len() == 0 {
					continue
				}
				// we are done building this number
				partNumber, _ := strconv.Atoi(numberBuilder.String())
				previousRow := grid.Row(row - 1)
				nextRow := grid.Row(row + 1)
				starting := col - numberBuilder.Len()

				if isAdjacent(numberBuilder.Len(), starting, col, grid.Row(row), previousRow, nextRow, partNumber) {
					numbers = append(numbers, partNumber)
				}
				numberBuilder.Reset()
			} else {
				numberBuilder.WriteRune(grid.Get(row, col))
			}
		}
	}

	//fmt.Printf("%v\n", numbers)
	fmt.Printf("The solution to part 1 is %d\n", sum(numbers))
}

func sum(s []int) int {
	sum := 0
	for _, value := range s {
		sum += value
	}

	return sum
}

func lineHasSymbol(window []rune) bool {
	for _, value := range window {
		if !unicode.IsDigit(value) && value != '.' {
			return true
		}
	}

	return false
}

func isAdjacent(lenth int, starting, end int, currentRow, previousRow, nextRow []rune, partNumber int) bool {
	startingIndex := starting - 1
	endingIndex := end
	if currentRow[starting-1] != '.' {
		return true
	}
	if currentRow[end] != '.' {
		return true
	}

	previousWindow := previousRow[startingIndex : endingIndex+1]
	if lineHasSymbol(previousWindow) {
		return true
	}

	nextWindow := nextRow[startingIndex : endingIndex+1]
	if lineHasSymbol(nextWindow) {
		return true
	}

	return false
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
