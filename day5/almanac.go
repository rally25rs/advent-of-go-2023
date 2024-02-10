package day5

type AlmanacMapItem struct {
	destStart   int
	sourceStart int
	rangeLength int
}

type AlmanacMap struct {
	items []AlmanacMapItem
}

type Almanac struct {
	seedsToPlant             []int
	seedToSoilMap            AlmanacMap
	soilToFertilizerMap      AlmanacMap
	fertilizerToWaterMap     AlmanacMap
	waterToLightMap          AlmanacMap
	lightToTemperatureMap    AlmanacMap
	temperatureToHumidityMap AlmanacMap
	humidityToLocationMap    AlmanacMap
}

func (ami *AlmanacMap) GetDestination(source int) int {
	return source
}

func (a *Almanac) GetLocationForSeed(seed int) int {
	return a.humidityToLocationMap.GetDestination(
		a.temperatureToHumidityMap.GetDestination(
			a.lightToTemperatureMap.GetDestination(
				a.waterToLightMap.GetDestination(
					a.fertilizerToWaterMap.GetDestination(
						a.soilToFertilizerMap.GetDestination(
							a.seedToSoilMap.GetDestination(seed)))))))
}
