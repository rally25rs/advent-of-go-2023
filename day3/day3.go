// https://adventofcode.com/2023/day/3

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type partnumber struct {
	number               int
	schematicRow         int
	schematicStartColumn int
	schematicEndColumn   int
	adjacentSymbol       bool
}

type gear struct {
	schematicRow    int
	schematicColumn int
	adjacentNumbers []int
}

func readSchematic() []string {
	inputFile, err := os.Open("./input.txt")
	schematic := make([]string, 0, 100)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		line := lineSplitter.Text()
		schematic = append(schematic, line)
	}

	return schematic
}

// Locate the position of all numbers in the schematic.
func locateNumbers(schematic []string) []partnumber {
	digits := regexp.MustCompile(`\d+`)
	partnumbers := make([]partnumber, 0, 100)
	for row, line := range schematic {
		allMatches := digits.FindAllString(line, -1)
		allIndexes := digits.FindAllStringIndex(line, -1)
		for i, match := range allMatches {
			num, err := strconv.Atoi(string(match))
			if err != nil {
				log.Fatal(err)
			}
			partnumber := partnumber{
				number:               num,
				schematicRow:         row,
				schematicStartColumn: allIndexes[i][0],
				schematicEndColumn:   allIndexes[i][1],
				adjacentSymbol:       hasAdjacentSymbol(schematic, row, allIndexes[i][0], allIndexes[i][1]),
			}
			partnumbers = append(partnumbers, partnumber)
		}
	}
	return partnumbers
}

func locateGears(schematic []string) []gear {
	gearSymbol := regexp.MustCompile(`\*`)
	gears := make([]gear, 0, 10)
	for row, line := range schematic {
		allIndexes := gearSymbol.FindAllStringIndex(line, -1)
		for _, indexes := range allIndexes {
			gear := gear{
				schematicRow:    row,
				schematicColumn: indexes[0],
				adjacentNumbers: make([]int, 0, 6),
			}
			gears = append(gears, gear)
		}
	}
	return gears
}

func hasAdjacentSymbol(schematic []string, schematicRow int, schematicStartColumn int, schematicEndColumn int) bool {
	symbolMatch := regexp.MustCompile(`[^\.\d]`)
	startRow := max(0, schematicRow-1)
	endRow := min(len(schematic)-1, schematicRow+2)

	startColumn := max(0, schematicStartColumn-1)
	endColumn := min(len(schematic[0])-1, schematicEndColumn+1)

	for _, row := range schematic[startRow:endRow] {
		if symbolMatch.MatchString(row[startColumn:endColumn]) {
			return true
		}
	}
	return false
}

func findAdjacentNumbers(gear gear, partNumbers []partnumber) []int {
	adjacentNumbers := make([]int, 0, 6)
	for _, partNumber := range partNumbers {
		if partNumber.schematicRow >= gear.schematicRow-1 && partNumber.schematicRow <= gear.schematicRow+1 {
			if partNumber.schematicStartColumn <= gear.schematicColumn+1 && partNumber.schematicEndColumn >= gear.schematicColumn {
				adjacentNumbers = append(adjacentNumbers, partNumber.number)
			}
		}
	}
	return adjacentNumbers
}

// Go really doesn't have filter/map/reduce? :(
// and people dislike JS because it has a small standard library...
// const sum = partNumbers.filter(part => !part.adjacentSymbol).reduce((filterAndSum, part) => filterAndSum + part.number, 0)
func filterAndSum(partNumbers []partnumber) int {
	sum := 0
	for _, partNumber := range partNumbers {
		if partNumber.adjacentSymbol {
			sum += partNumber.number
		}
	}
	return sum
}

func sumGearRatios(gears []gear) int {
	sum := 0
	for _, gear := range gears {
		ratio := 0
		if len(gear.adjacentNumbers) == 2 {
			for _, n := range gear.adjacentNumbers {
				if ratio == 0 {
					ratio = n
				} else {
					ratio = ratio * n
				}
			}
		}
		sum += ratio
	}
	return sum
}

func main() {
	schematic := readSchematic()
	partNumbers := locateNumbers(schematic)
	fmt.Println("Part 1:", filterAndSum(partNumbers)) // 498559

	gears := locateGears(schematic)
	for i := range gears {
		gear := &gears[i]
		gear.adjacentNumbers = findAdjacentNumbers(*gear, partNumbers)
	}
	fmt.Println("Part 2:", sumGearRatios(gears)) // 72246648
}
