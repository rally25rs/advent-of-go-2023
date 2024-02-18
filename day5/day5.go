package day5

import (
	"fmt"

	"github.com/samber/lo"
)

func findMinLocationPart1(almanac *Almanac, seedsToPlant []int) int {
	locations := lo.Map(seedsToPlant, func(seed int, _ int) int { return almanac.GetLocationForSeed(seed) })
	return lo.Min(locations)
}

func findMinLocationPart2(almanac *Almanac, seedsToPlant []int) int {
	locations := make([]int, 0, 1000)
	for i := 0; i < len(seedsToPlant); i += 2 {
		start := seedsToPlant[i]
		end := start + seedsToPlant[i+1]
		fmt.Println("Checking Seeds", start, end)

		for j := start; j < end; j++ {
			locations = append(locations, almanac.GetLocationForSeed(j))
		}
	}
	return lo.Min(locations)
}

func Execute(filename string) {
	almanac, seedsToPlant := readAlmanac(filename)
	fmt.Println("Part 1:", findMinLocationPart1(almanac, seedsToPlant)) // sampleinput = 35 // input = 107430936
	fmt.Println("Part 2:", findMinLocationPart2(almanac, seedsToPlant)) // sampleinput = 46 // input =
}
