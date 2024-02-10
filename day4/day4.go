package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func filterEmptyStrings(s string, _ int) bool {
	return len(strings.Trim(s, " ")) > 0
}

func processCard(line string) (int, int) {
	regexp := regexp.MustCompile(`([\d\ ]+)\|([\d\ ]+)`)
	matches := regexp.FindStringSubmatch(line)
	winningNumbers := lo.Filter(strings.Split(matches[1], " "), filterEmptyStrings)
	ourNumbers := lo.Filter(strings.Split(matches[2], " "), filterEmptyStrings)
	matchingNumbers := lo.Intersect(winningNumbers, ourNumbers)
	cardValue := 2 ^ len(matchingNumbers)
	return cardValue, len(matchingNumbers)
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

func Execute(filename string) {
	inputFile, err := os.Open(filename)
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
