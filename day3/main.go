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

type NumberElement struct {
	start int
	end   int
	row   int
	value int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please Input file to read\n")
	}
	solvePart1(os.Args[1])
}

func solvePart1(file string) {
	lines := readLines(file)
	grid := createGrid(lines)
	var numbers []NumberElement

	for row := 0; row < grid.Rows(); row++ {
		rowNumbers := readNumbers(grid.Row(row), row)
		numbers = append(numbers, rowNumbers...)
	}

	var partNumbers []int
	for _, number := range numbers {
		if hasSymbolAround(number, grid) {
			partNumbers = append(partNumbers, number.value)
		}
	}

	fmt.Printf("The solution to part 1 is %d\n", sum(partNumbers))
}

func createGrid(lines []string) *multiarray.TwoDArray[rune] {
	grid := multiarray.New[rune](len(lines), len(lines[0]))

	for i, line := range lines {
		for j, value := range line {
			grid.Set(i, j, value)
		}
	}
	return grid
}

func readNumbers(row []rune, rowNumber int) []NumberElement {
	var numberBuilder strings.Builder
	start := 0
	numbers := make([]NumberElement, 0)

	for i, char := range row {
		if !unicode.IsDigit(char) {
			if numberBuilder.Len() == 0 {
				continue
			}
			partNumber, _ := strconv.Atoi(numberBuilder.String())
			numberElem := NumberElement{
				value: partNumber,
				row:   rowNumber,
				start: start,
				end:   i,
			}
			numbers = append(numbers, numberElem)
			numberBuilder.Reset()
		} else {
			if numberBuilder.Len() == 0 {
				start = i
			}
			numberBuilder.WriteRune(char)
		}
	}
	if numberBuilder.Len() != 0 {
		number, _ := strconv.Atoi(numberBuilder.String())
		numberElem := NumberElement{
			value: number,
			row:   rowNumber,
			start: start,
			end:   len(row),
		}
		numbers = append(numbers, numberElem)
	}

	return numbers
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

func hasSymbolAround(number NumberElement, grid *multiarray.TwoDArray[rune]) bool {
	// check around current around to see if it's beside it
	row := grid.Row(number.row)
	if number.start > 0 && row[number.start-1] != '.' {
		return true
	}
	if number.end < len(row) && row[number.end] != '.' {
		return true
	}

	// check to see if we can check the previous row
	if number.row != 0 {
		prevRow := grid.Row(number.row - 1)
		start := clampAbove(number.start-1, 0)
		end := clampBelow(number.end+1, len(prevRow))
		window := prevRow[start:end]
		if lineHasSymbol(window) {
			return true
		}
	}

	if number.row < grid.Rows()-1 {
		nextRow := grid.Row(number.row + 1)
		start := clampAbove(number.start-1, 0)
		end := clampBelow(number.end+1, len(nextRow))
		window := nextRow[start:end]
		if lineHasSymbol(window) {
			return true
		}
	}

	return false
}

func clampAbove(value, clampTo int) int {
	if value < clampTo {
		return clampTo
	}

	return value
}

func clampBelow(value, clampTo int) int {
	if value > clampTo {
		return clampTo
	}

	return value
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
