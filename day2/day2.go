// https://adventofcode.com/2023/day/2

package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

type Game struct {
	Number int
	Red    int
	Green  int
	Blue   int
}

func processLine(line []byte) Game {
	gameRegexp := regexp.MustCompile(`Game (\d+)`)
	gameMatch := gameRegexp.FindSubmatch(line)
	gameId, err := strconv.Atoi(string(gameMatch[1]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	red := max(`(\d+) red`, line)
	green := max(`(\d+) green`, line)
	blue := max(`(\d+) blue`, line)

	return Game{Number: gameId, Red: red, Green: green, Blue: blue}
}

func max(regexString string, line []byte) int {
	max := 0
	rx := regexp.MustCompile(regexString)
	matches := rx.FindAllSubmatch(line, -1)
	for _, match := range matches {
		asInt, err := strconv.Atoi(string(match[1]))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if max < asInt {
			max = asInt
		}
	}
	return max
}

func calculatePower(game Game) int {
	return game.Red * game.Green * game.Blue
}

func Execute(filename string) {
	part1Sum := 0
	part2Sum := 0

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		line := lineSplitter.Bytes()
		game := processLine(line)
		gameValid := game.Red <= MAX_RED && game.Green <= MAX_GREEN && game.Blue <= MAX_BLUE
		if gameValid {
			part1Sum += game.Number
		}
		part2Sum += calculatePower(game)
	}

	fmt.Println("Part 1: ", part1Sum)
	fmt.Println("Part 2: ", part2Sum)
}
