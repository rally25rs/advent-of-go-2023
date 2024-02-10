// https://adventofcode.com/2023/day/1

package day1

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var tokens = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

// custom bufio.SplitFunc
// Moves 1 byte at a time through the line and checks if any of the keys in the tokens map are at the start.
// If so, return the map value as the token. If not, return an empty byte slice to indicate no match.
func splitOnDigits(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	for key, value := range tokens {
		if strings.Index(string(data), key) == 0 {
			return 1, []byte(value), nil
		}
	}
	return 1, []byte{}, nil
}

func getLineValue(line []byte) int {
	var first, last int

	scanner := bufio.NewScanner(bytes.NewReader(line))
	scanner.Split(splitOnDigits)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue // ignore empty tokens since they are not a digit.
		}
		asInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			os.Exit(3)
		}
		if first == 0 {
			first = asInt
		}
		last = asInt
	}

	value := first*10 + last
	return value
}

func Execute(filename string) {
	sum := 0

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		line := lineSplitter.Bytes()
		sum += getLineValue(line)
	}

	fmt.Println("Part 2: ", sum)
}
