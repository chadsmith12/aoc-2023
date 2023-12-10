package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

// Represents a range map to map a source value to a destination value.
type SourceRangeMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

// Holds all the ranges to map sources to destination for a section
// If any of our range maps don't fit for a value, then just assume it is 1-1
type SectionMap struct {
	ranges []SourceRangeMap
}

func parseRange(line string) SourceRangeMap {
	parts := strings.Fields(line)
	rangeMap := SourceRangeMap{}
	for i, value := range parts {
		number, _ := strconv.Atoi(value)
		if i == 0 {
			rangeMap.destinationRangeStart = number
		} else if i == 1 {
			rangeMap.sourceRangeStart = number
		} else {
			rangeMap.rangeLength = number
		}
	}

	return rangeMap
}

func createMap(ranges []SourceRangeMap) SectionMap {
	return SectionMap{
		ranges: ranges,
	}
}

// the offset for a range is the difference between the destination and start, when not 1-1
func (m SourceRangeMap) offset() int {
	return m.destinationRangeStart - m.sourceRangeStart
}

// Tells you if the value you pass in will fit in this range and needs to mapped differently
func (m SourceRangeMap) isInRange(value int) bool {
	end := m.sourceRangeStart + m.rangeLength

	return value >= m.sourceRangeStart && value < end
}

// if the value passed in fits in this range, finds the new destination based off this range.
// if the value passed in doesn't fit in this range, then just return the value back.
func (m SourceRangeMap) calculateDestination(value int) int {
	if m.isInRange(value) {
		return value + m.offset()
	}

	return value
}

// Given a value, finds the range, if any, that should be used to calculate the destination for this value
// Returns a pointer to the range, nil if not in the range
func (s SectionMap) findRange(value int) *SourceRangeMap {
	if len(s.ranges) == 0 {
		return nil
	}

	for _, rangeMap := range s.ranges {
		if rangeMap.isInRange(value) {
			return &rangeMap
		}
	}

	return nil
}

// Returns the destination for the source value by finding what range is belongs in, and calculating from that range
func (s SectionMap) destination(source int) int {
	sectionRange := s.findRange(source)

	if sectionRange == nil {
		return source
	}

	return sectionRange.calculateDestination(source)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please input file to read\n")
	}

	solvePart1(os.Args[1])
}

func solvePart1(file string) {
	lines := readLines(file)
	seeds := parseSeeds(lines[0])
	maps := lines[2:]
	sections := getSections(maps)
	locations := make([]int, 0, len(seeds))
	for _, seed := range seeds {
		locations = append(locations, findLocation(sections, seed))
	}

	fmt.Printf("The solution for part 1 is %d\n", slices.Min(locations))
}

func findLocation(maps []SectionMap, seed int) int {
	seed2Soil := maps[0]
	soil2Fertilizer := maps[1]
	fertilizer2Water := maps[2]
	water2Light := maps[3]
	light2Temperature := maps[4]
	temperature2Humidity := maps[5]
	humidity2Location := maps[6]

	soil := seed2Soil.destination(seed)
	fertizer := soil2Fertilizer.destination(soil)
	water := fertilizer2Water.destination(fertizer)
	light := water2Light.destination(water)
	temperature := light2Temperature.destination(light)
	humidity := temperature2Humidity.destination(temperature)

	return humidity2Location.destination(humidity)
}

func parseSeeds(line string) []int {
	parts := strings.Split(line, ":")
	seedStrings := strings.Fields(parts[1])
	seeds := make([]int, 0, len(seedStrings))
	for _, value := range seedStrings {
		if number, err := strconv.Atoi(value); err == nil {
			seeds = append(seeds, number)
		}
	}

	return seeds
}

func getSections(lines []string) []SectionMap {
	var sections []SectionMap
	section := []SourceRangeMap{}

	for _, line := range lines {
		if line == "" {
			sectionMap := createMap(section)
			sections = append(sections, sectionMap)
			section = []SourceRangeMap{}
		} else if !unicode.IsDigit(rune(line[0])) {
			continue
		} else {
			rangeMap := parseRange(line)
			section = append(section, rangeMap)
		}
	}

	if len(section) > 0 {
		sectionMap := createMap(section)
		sections = append(sections, sectionMap)
	}

	return sections
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
