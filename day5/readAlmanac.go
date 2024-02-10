package day5

import (
	"bufio"
	"log"
	"os"
)

func readAlmanac(filename string) *Almanac {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		// line := lineSplitter.Text()
	}
	return &Almanac{}
}
