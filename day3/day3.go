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

func readSchematic() [][]byte {
	inputFile, err := os.Open("./input.txt")
	schematic := make([][]byte, 0)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		line := lineSplitter.Text()
		schematic = append(schematic, []byte(line))
	}

	return schematic
}

// Locate the position of all numbers in the schematic.
func locateNumbers(schematic [][]byte) []partnumber {
	digits := regexp.MustCompile(`\d+`)
	partnumbers := make([]partnumber, 0, 100)
	for row, line := range schematic {
		allMatches := digits.FindAll(line, -1)
		allIndexes := digits.FindAllIndex(line, -1)
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

func hasAdjacentSymbol(schematic [][]byte, schematicRow int, schematicStartColumn int, schematicEndColumn int) bool {
	symbolMatch := regexp.MustCompile(`[^\.\d]`)
	startRow := max(0, schematicRow-1)
	endRow := min(len(schematic)-1, schematicRow+2)

	startColumn := max(0, schematicStartColumn-1)
	endColumn := min(len(schematic[0])-1, schematicEndColumn+1)

	for _, row := range schematic[startRow:endRow] {
		if symbolMatch.Match(row[startColumn:endColumn]) {
			return true
		}
	}
	return false
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

func main() {
	schematic := readSchematic()
	partNumbers := locateNumbers(schematic)
	fmt.Println("Part 1:", filterAndSum(partNumbers))
}
