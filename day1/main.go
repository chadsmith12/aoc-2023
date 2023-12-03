package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please input file to read")
		os.Exit(1)
	}
	file := os.Args[1]
	sum := 0
	for _, line := range readLines(file) {
		number := parseLine(line)
		sum += number
	}

	fmt.Printf("Solution For %v is %d\n", file, sum)
}

func readLines(file string) []string {
	f, err := os.Open(file)
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

func parseLine(line string) int {
	leftMost, rightMost := "", ""
	for _, c := range line {
		if !unicode.IsDigit(c) {
			continue
		}

		if leftMost == "" {
			leftMost = string(c)
		}
		rightMost = string(c)
	}

	number, _ := strconv.Atoi(leftMost + rightMost)

	return number
}
