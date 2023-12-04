package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var word2Digit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please input file to read")
		os.Exit(1)
	}

	file := os.Args[1]
	solvePart1(file)
	solvePart2(file)
}

func solvePart1(file string) {
	sum := 0
	for _, line := range readLines(file) {
		number := parseLinePart1(line)
		sum += number
	}

	fmt.Printf("Solution For %v part 1 is %d\n", file, sum)
}

func solvePart2(file string) {
	sum := 0
	for _, line := range readLines(file) {
		number := parseLinePart2(line)
		sum += number
	}

	fmt.Printf("Solution for %v part 2 is %d\n", file, sum)
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

func parseLinePart1(line string) int {
	var digits []int
	for _, c := range line {
		if !unicode.IsDigit(c) {
			continue
		}

		digit, _ := strconv.Atoi(string(c))
		digits = append(digits, digit)
	}

	if len(digits) == 0 {
		return 0
	}
	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	number, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))

	return number
}

func parseLinePart2(line string) int {
	var digits []int

	for i := 0; i < len(line); i++ {
		for word, digit := range word2Digit {
			if len(word) > len(line)-i {
				continue
			}
			if line[i:i+len(word)] == word {
				digits = append(digits, digit)
				// we are subtracting 2 here so we start from the end of this word
				// otherwise we could miss like: "2oneight"
				i += len(word) - 2
				break
			}
		}
		if unicode.IsDigit(rune(line[i])) {
			number, _ := strconv.Atoi(string(line[i]))
			digits = append(digits, number)
		}
	}

	firstDigit := digits[0]
	lastDigit := digits[len(digits)-1]
	number, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))

	return number
}
