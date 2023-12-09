package garden

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

const numSep = " "
const finishSignal = -1

const (
	seedsTitle                 = "seeds:"
	seedsToSoilTitle           = "seed-to-soil map:"
	soilToFertilizerTitle      = "soil-to-fertilizer map:"
	fertilizerToWaterTitle     = "fertilizer-to-water map:"
	waterToLightTitle          = "water-to-light map:"
	lightToTemperatureTitle    = "light-to-temperature map:"
	temperatureToHumidityTitle = "temperature-to-humidity map:"
	humidityToLocationTitle    = "humidity-to-location map:"
)

const (
	seedsToSoilIdx = iota
	soilToFertilizerIdx
	fertilizerToWaterIdx
	waterToLightIdx
	lightToTemperatureIdx
	temperatureToHumidityIdx
	humidityToLocationIdx
)

var titleToIndex = map[string]int{
	seedsToSoilTitle:           seedsToSoilIdx,
	soilToFertilizerTitle:      soilToFertilizerIdx,
	fertilizerToWaterTitle:     fertilizerToWaterIdx,
	waterToLightTitle:          waterToLightIdx,
	lightToTemperatureTitle:    lightToTemperatureIdx,
	temperatureToHumidityTitle: temperatureToHumidityIdx,
	humidityToLocationTitle:    humidityToLocationIdx,
}

func GetLowestSeedLocation(almanac Almanac) int {
	minLocation := math.MaxInt
	for _, seed := range almanac.Seeds {
		location := almanac.GetSeedLocation(seed)
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func GetLowestSeedRangedLocation(almanac Almanac) int {
	minLocation := math.MaxInt
	rangesCount := 0
	locationChn := make(chan int, 15)
	ranges := almanac.GetSeedRanges()
	for i := 0; i < len(ranges); i++ {
		seedRange := ranges[i]
		rangesCount += 1
		go processSeedRangeLocationMinAsync(&seedRange, &almanac, locationChn)
	}

	for rangesCount > 0 {
		location := <-locationChn
		if location < minLocation {
			minLocation = location
		}
		rangesCount -= 1
	}

	return minLocation
}

func ParseGardenAlmanac(input []string) Almanac {
	almanacMaps := [7][]Range{}
	seeds := parseAlmanacSeeds(input[0])
	for i := 1; i < len(input); {
		line := input[i]
		if isTitleLine(line) {
			mapIdx := titleToIndex[line]
			almanacMap, nextIdx := parseAlmanacMap(input, i)
			almanacMaps[mapIdx] = almanacMap

			i = nextIdx

		} else if isNumberLine(line) {
			panic("unexpected line number")

		} else {
			i += 1
		}
	}

	return *NewAlmanac(seeds, almanacMaps)
}

type Range struct {
	From      int
	To        int
	Increment int
}

func (r *Range) IsInRange(num int) bool {
	return num >= r.From && num <= r.To
}

func (r *Range) GenerateRange() <-chan int {
	chn := make(chan int, 10)
	go func() {
		for i := r.From; i <= r.To; i++ {
			chn <- i
		}

		close(chn)
	}()

	return chn
}

type Almanac struct {
	Seeds                 []int
	seedToSoil            []Range
	soilToFertilizer      []Range
	fertilizerToWater     []Range
	waterToLight          []Range
	lightToTemperature    []Range
	temperatureToHumidity []Range
	humidityToLocation    []Range
}

func NewAlmanac(seeds []int, almanacMaps [7][]Range) *Almanac {
	return &Almanac{
		Seeds:                 seeds,
		seedToSoil:            almanacMaps[seedsToSoilIdx],
		soilToFertilizer:      almanacMaps[soilToFertilizerIdx],
		fertilizerToWater:     almanacMaps[fertilizerToWaterIdx],
		waterToLight:          almanacMaps[waterToLightIdx],
		lightToTemperature:    almanacMaps[lightToTemperatureIdx],
		temperatureToHumidity: almanacMaps[temperatureToHumidityIdx],
		humidityToLocation:    almanacMaps[humidityToLocationIdx],
	}
}

func (a Almanac) GetSeedRanges() []Range {
	seedsRanges := make([]Range, 0)
	for i := 0; i < len(a.Seeds); i += 2 {
		seedFrom := a.Seeds[i]
		seedTo := seedFrom + a.Seeds[i+1] - 1
		seedsRanges = append(seedsRanges, Range{seedFrom, seedTo, 0})
	}

	return seedsRanges
}

func (a Almanac) GetSeedSoil(seed int) int {
	return getValueFromRanges(a.seedToSoil, seed)
}

func (a Almanac) GetSoilFertilizer(soil int) int {
	return getValueFromRanges(a.soilToFertilizer, soil)
}

func (a Almanac) GetFertilizerWater(fertilizer int) int {
	return getValueFromRanges(a.fertilizerToWater, fertilizer)
}

func (a Almanac) GetWaterLight(water int) int {
	return getValueFromRanges(a.waterToLight, water)
}

func (a Almanac) GetLightTemperature(light int) int {
	return getValueFromRanges(a.lightToTemperature, light)
}

func (a Almanac) GetTemperatureHumidity(temp int) int {
	return getValueFromRanges(a.temperatureToHumidity, temp)
}

func (a Almanac) GetHumidityLocation(humidity int) int {
	return getValueFromRanges(a.humidityToLocation, humidity)
}

func (a Almanac) GetSeedLocation(seed int) int {
	nextItem := a.GetSeedSoil(seed)
	nextItem = a.GetSoilFertilizer(nextItem)
	nextItem = a.GetFertilizerWater(nextItem)
	nextItem = a.GetWaterLight(nextItem)
	nextItem = a.GetLightTemperature(nextItem)
	nextItem = a.GetTemperatureHumidity(nextItem)

	return a.GetHumidityLocation(nextItem)
}

func getValueFromRanges(ranges []Range, num int) int {
	for _, rng := range ranges {
		if rng.IsInRange(num) {
			return num + rng.Increment
		}
	}

	return num
}

func parseAlmanacMap(input []string, lineIdx int) ([]Range, int) {
	if !isTitleLine(input[lineIdx]) {
		panic(fmt.Sprintf("expected title line to start map parsing: %s", input[lineIdx]))
	}

	idx := lineIdx + 1
	almanacMap := make([]Range, 0)
	for ; idx < len(input); idx++ {
		if !isNumberLine(input[idx]) {
			break
		}
		nums := parseNumsLine(input[idx])

		srcNum := nums[1]
		dstNum := nums[0]
		numRange := nums[2]
		almanacMap = append(almanacMap, Range{srcNum, srcNum + numRange - 1, dstNum - srcNum})
	}

	return almanacMap, idx + 1
}

func parseAlmanacSeeds(line string) []int {
	if !strings.HasPrefix(line, seedsTitle) {
		panic("missing seeds line")
	}

	numsStartIdx := strings.Index(line, ":")

	return parseNumsLine(line[numsStartIdx+2:])
}

func parseNumsLine(line string) []int {
	numsData := strings.Split(line, numSep)
	nums := make([]int, 0)
	for _, numIt := range numsData {
		if num, err := strconv.Atoi(numIt); err == nil {
			nums = append(nums, num)
		}
	}

	return nums
}

func processSeedRangeLocationMinAsync(seedRange *Range, almanac *Almanac, chn chan<- int) {
	locationMin := math.MaxInt
	for seed := range seedRange.GenerateRange() {
		location := almanac.GetSeedLocation(seed)
		if location < locationMin {
			locationMin = location
		}
	}

	chn <- locationMin
}

func isTitleLine(line string) bool {
	if len(line) == 0 {
		return false
	}
	firstChar := rune(line[0])

	return unicode.IsLetter(firstChar)
}

func isNumberLine(line string) bool {
	if len(line) == 0 {
		return false
	}
	firstChar := rune(line[0])

	return unicode.IsDigit(firstChar)
}
