package day5

import (
	"fmt"

	"github.com/samber/lo"
)

func findMinLocation(almanac *Almanac) int {
	locations := lo.Map(almanac.seedsToPlant, func(seed int, _ int) int { return almanac.GetLocationForSeed(seed) })
	return lo.Min(locations)
}

func Execute(filename string) {
	almanac := readAlmanac(filename)
	fmt.Println("Part 1:", findMinLocation(almanac)) // sampleinput = 35
}
