package day5

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readNumbersFromLine(line string) []int {
	digitsRegex := regexp.MustCompile(`\d+`)
	matches := digitsRegex.FindAllString(line, -1)
	var result []int
	for _, number := range matches {
		n, err := strconv.Atoi(number)
		if err == nil {
			result = append(result, n)
		}
	}
	return result
}

func readAlmanac(filename string) (*Almanac, []int) {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	almanac := NewAlmanac()
	var readIntoMap *AlmanacMap
	var seedsToPlant []int

	lineSplitter := bufio.NewScanner(inputFile)
	for lineSplitter.Scan() {
		line := lineSplitter.Text()

		if line == "" {
			continue
		}

		if strings.Index(line, "seeds:") == 0 {
			seedsToPlant = readNumbersFromLine(line)
		} else if line == "seed-to-soil map:" {
			readIntoMap = &almanac.seedToSoilMap
		} else if line == "soil-to-fertilizer map:" {
			readIntoMap = &almanac.soilToFertilizerMap
		} else if line == "fertilizer-to-water map:" {
			readIntoMap = &almanac.fertilizerToWaterMap
		} else if line == "water-to-light map:" {
			readIntoMap = &almanac.waterToLightMap
		} else if line == "light-to-temperature map:" {
			readIntoMap = &almanac.lightToTemperatureMap
		} else if line == "temperature-to-humidity map:" {
			readIntoMap = &almanac.temperatureToHumidityMap
		} else if line == "humidity-to-location map:" {
			readIntoMap = &almanac.humidityToLocationMap
		} else {
			numbers := readNumbersFromLine(line)
			if len(numbers) != 3 {
				log.Fatalf("Expected 3 numbers in %v, got %d", line, len(numbers))
			}
			readIntoMap.items = append(readIntoMap.items, AlmanacMapItem{
				destStart:   numbers[0],
				sourceStart: numbers[1],
				rangeLength: numbers[2],
			})
		}
	}
	return almanac, seedsToPlant
}
