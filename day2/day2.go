// https://adventofcode.com/2023/day/2

package main

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

func main() {
	sum := 0

	inputFile, err := os.Open("./input.txt")
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
			sum += game.Number
		}
		// fmt.Println("Line", string(line))
		// fmt.Println("Game", game.Number, "Red:", game.Red, "Green:", game.Green, "Blue:", game.Blue, "Valid?", gameValid, "/n")
	}

	fmt.Println("Part 1: ", sum)
}
