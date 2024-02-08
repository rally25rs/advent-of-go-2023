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

func processCard(line string) (int, int) {
	cardValue := 0
	numWinningNumbers := 0
	regexp := regexp.MustCompile(`([\d\ ]+)\|([\d\ ]+)`)
	matches := regexp.FindStringSubmatch(line)
	winningNumbers := strings.Split(matches[1], " ")
	ourNumbers := strings.Split(matches[2], " ")
	for _, ourNumber := range ourNumbers {
		if len(strings.Trim(ourNumber, " ")) == 0 {
			continue
		}

		if slices.Contains(winningNumbers, ourNumber) {
			numWinningNumbers += 1
			if cardValue == 0 {
				cardValue = 1
			} else {
				cardValue = cardValue * 2
			}
		}
	}
	return cardValue, numWinningNumbers
}

func sumCardValues(cardInfo []CardInfo) int {
	sum := 0
	for _, card := range cardInfo {
		sum += card.value
	}
	return sum
}

func countMoreCards(cardInfo []CardInfo, cardIndex int) int {
	card := cardInfo[cardIndex]
	count := 1
	for i := 1; i <= card.numWinningNumbers; i++ {
		count += countMoreCards(cardInfo, cardIndex+i)
	}
	return count
}

func countAllCards(cardInfo []CardInfo) int {
	count := 0
	for i := 0; i < len(cardInfo); i++ {
		count += countMoreCards(cardInfo, i)
	}
	return count
}

type CardInfo struct {
	value             int
	numWinningNumbers int
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cardInfo := make([]CardInfo, 0)

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		line := lineSplitter.Text()
		cardValue, numWinningNumbers := processCard(line)
		cardInfo = append(cardInfo, CardInfo{cardValue, numWinningNumbers})
	}

	fmt.Println("Part 1:", sumCardValues(cardInfo)) // sampleinput = 13 // input = 22897
	fmt.Println("Part 2:", countAllCards(cardInfo)) // sampleinput = 30 // input = 5095824
}
