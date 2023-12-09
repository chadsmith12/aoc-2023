package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type WinningCard struct {
	numbers map[int]struct{}
}

func newCard(line string) *WinningCard {
	fields := strings.Split(line, ":")
	winningNumbers := numberSlice(strings.Split(strings.TrimSpace(fields[1]), " "))
	card := &WinningCard{
		numbers: make(map[int]struct{}),
	}
	for _, winningNumber := range winningNumbers {
		card.numbers[winningNumber] = struct{}{}
	}

	return card
}

func (c *WinningCard) getWinningNumbers(numbers []int) []int {
	winningNumbers := make([]int, 0, 10)

	for _, number := range numbers {
		if _, ok := c.numbers[number]; ok {
			winningNumbers = append(winningNumbers, number)
		}
	}

	return winningNumbers
}

func (c *WinningCard) calculateScore(winningNumbers []int) int {
	points := 0

	for range winningNumbers {
		if points == 0 {
			points++
		} else {
			points *= 2
		}
	}

	return points
}

func numberSlice(chars []string) []int {
	numbers := make([]int, 0, len(chars))

	for _, value := range chars {
		if number, err := strconv.Atoi(value); err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
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
		parts := strings.Split(line, "|")
		myNumbers := numberSlice(strings.Split(strings.TrimSpace(parts[1]), " "))
		card := newCard(parts[0])
		sum += card.calculateScore(card.getWinningNumbers(myNumbers))
	}

	fmt.Printf("The solution for part 1 is %d\n", sum)
}

func solvePart2(file string) {
	lines := readLines(file)
	matches := make([]int, len(lines))
	for i := range lines {
		matches[i] = 1
	}
	for i, line := range lines {
		parts := strings.Split(line, "|")
		myNumbers := numberSlice(strings.Split(strings.TrimSpace(parts[1]), " "))
		card := newCard(parts[0])
		winningNumbers := card.getWinningNumbers(myNumbers)
		countTil := i + len(winningNumbers)
		if countTil >= len(lines) {
			countTil = len(lines) - 1
		}

		for n := i + 1; n <= countTil; n++ {
			matches[n] += matches[i]
		}
	}

	sum := 0
	for _, value := range matches {
		sum += value
	}

	fmt.Printf("The solution for part 2 is %d\n", sum)
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
