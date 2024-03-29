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

func readPart1Input(filename string) ([]race, error) {
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

func readPart2Input(filename string) (race, error) {
	digitsRegex := regexp.MustCompile(`[^\d]+`)
	inputFile, err := os.Open(filename)
	if err != nil {
		return race{}, err
	}

	lineSplitter := bufio.NewScanner(inputFile)
	lineSplitter.Scan()
	if lineSplitter.Err() != nil {
		return race{}, err
	}
	timesLine := lineSplitter.Text()
	lineSplitter.Scan()
	if lineSplitter.Err() != nil {
		return race{}, err
	}
	distancesLine := lineSplitter.Text()

	raceTime := digitsRegex.ReplaceAllString(timesLine, "")
	raceDistance := digitsRegex.ReplaceAllString(distancesLine, "")
	time, err := strconv.Atoi(raceTime)
	if err != nil {
		return race{}, err
	}
	distance, err := strconv.Atoi(raceDistance)
	if err != nil {
		return race{}, err
	}

	return race{
		time:     time,
		distance: distance,
	}, nil
}

func getRaceOutcomes(race race) (int, error) {
	for speed := 0; speed < race.time; speed++ {
		timeToMove := race.time - speed
		distanceMoved := speed * timeToMove
		if distanceMoved > race.distance {
			return race.time + 1 - speed - speed, nil
		}
	}
	return 0, fmt.Errorf("Unwinnable")
}

func part1(filename string) int {
	races, err := readPart1Input(filename)
	if err != nil {
		log.Fatal("Error reading input:", err)
		os.Exit(1)
	}

	winValue := 1
	for _, race := range races {
		numWins, err := getRaceOutcomes(race)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		winValue *= numWins
	}
	return winValue
}

func part2(filename string) int {
	race, err := readPart2Input(filename)
	if err != nil {
		log.Fatal("Error reading input:", err)
		os.Exit(1)
	}
	outcomes, err := getRaceOutcomes(race)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return outcomes
}

func Execute(filename string) {
	fmt.Println("Part 1:", part1(filename)) // sampleinput = 288 // input = 2269432
	fmt.Println("Part 2:", part2(filename)) // sampleinput = 71503 // input = 35865985
}
