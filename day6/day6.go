package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type race struct {
	time     int
	distance int
}

func readInput(filename string) ([]race, error) {
	digitsRegex := regexp.MustCompile(`\d+`)
	inputFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	lineSplitter := bufio.NewScanner(inputFile)
	lineSplitter.Scan()
	if lineSplitter.Err() != nil {
		return nil, err
	}
	timesLine := lineSplitter.Text()
	lineSplitter.Scan()
	if lineSplitter.Err() != nil {
		return nil, err
	}
	distancesLine := lineSplitter.Text()

	raceTimes := digitsRegex.FindAllString(timesLine, -1)
	raceDistances := digitsRegex.FindAllString(distancesLine, -1)
	races := make([]race, len(raceTimes))

	for i := 0; i < len(raceTimes); i++ {
		races[i].time, err = strconv.Atoi(raceTimes[i])
		if err != nil {
			return nil, err
		}
		races[i].distance, err = strconv.Atoi(raceDistances[i])
		if err != nil {
			return nil, err
		}
	}

	return races, nil
}

func getRaceOutcomes(race race) int {
	wins := 0
	for speed := 0; speed < race.time; speed++ {
		timeToMove := race.time - speed
		distanceMoved := speed * timeToMove
		if distanceMoved > race.distance {
			wins++
		}
	}
	return wins
}

func part1(races []race) int {
	winValue := 0
	for _, race := range races {
		numWins := getRaceOutcomes(race)
		if winValue == 0 {
			winValue = numWins
		} else {
			winValue = winValue * numWins
		}
	}
	return winValue
}

func Execute(filename string) {
	races, err := readInput(filename)
	if err != nil {
		log.Fatal("Error reading input:", err)
		os.Exit(1)
	}
	fmt.Println("Part 1:", part1(races)) // sampleinput = 288 // input = 2269432
}
