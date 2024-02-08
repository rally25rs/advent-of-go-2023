package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func getLineValue(line string) int {
	sum := 0
	regexp := regexp.MustCompile(`([\d\ ]+)\|([\d\ ]+)`)
	matches := regexp.FindStringSubmatch(line)
	winningNumbers := strings.Split(matches[1], " ")
	ourNumbers := strings.Split(matches[2], " ")
	for _, ourNumber := range ourNumbers {
		if len(strings.Trim(ourNumber, " ")) == 0 {
			continue
		}

		if slices.Contains(winningNumbers, ourNumber) {
			if sum == 0 {
				sum = 1
			} else {
				sum = sum * 2
			}
		}
	}
	return sum
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lineSplitter := bufio.NewScanner(inputFile)
	sum := 0
	for lineSplitter.Scan() {
		line := lineSplitter.Text()
		sum += getLineValue(line)
	}

	fmt.Println("Part 1:", sum) // sampleinput = 13 // input = 22897
}
