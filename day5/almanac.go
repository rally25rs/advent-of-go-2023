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

func NewAlmanac() *Almanac {
	return &Almanac{
		seedsToPlant: []int{},
		seedToSoilMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
		soilToFertilizerMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
		fertilizerToWaterMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
		waterToLightMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
		lightToTemperatureMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
		temperatureToHumidityMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
		humidityToLocationMap: AlmanacMap{
			items: make([]AlmanacMapItem, 0),
		},
	}
}

func (am *AlmanacMap) GetDestination(source int) int {
	for _, item := range am.items {
		if source >= item.sourceStart && source < item.sourceStart+item.rangeLength {
			dest := item.destStart + (source - item.sourceStart)
			// fmt.Println(source, ">", dest)
			return dest
		}
	}
	// fmt.Println(source, ">", source)
	return source
}

func (a *Almanac) GetLocationForSeed(seed int) int {
	// fmt.Println("Seed", seed)
	return a.humidityToLocationMap.GetDestination(
		a.temperatureToHumidityMap.GetDestination(
			a.lightToTemperatureMap.GetDestination(
				a.waterToLightMap.GetDestination(
					a.fertilizerToWaterMap.GetDestination(
						a.soilToFertilizerMap.GetDestination(
							a.seedToSoilMap.GetDestination(seed)))))))
}
