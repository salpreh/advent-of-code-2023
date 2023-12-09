package garden

import (
	"com.github/salpreh/advent-of-code-2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

const numSep = " "

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

func ParseGardenAlmanac(input []string) Almanac {
	almanacMaps := [7]map[int]int{}
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

type Almanac struct {
	Seeds                 []int
	seedToSoil            map[int]int
	soilToFertilizer      map[int]int
	fertilizerToWater     map[int]int
	waterToLight          map[int]int
	lightToTemperature    map[int]int
	temperatureToHumidity map[int]int
	humidityToLocation    map[int]int
}

func NewAlmanac(seeds []int, almanacMaps [7]map[int]int) *Almanac {
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

func (a Almanac) GetSeedSoil(seed int) int {
	return utils.GetOrDefault(a.seedToSoil, seed, seed)
}

func (a Almanac) GetSoilFertilizer(soil int) int {
	return utils.GetOrDefault(a.soilToFertilizer, soil, soil)
}

func (a Almanac) GetFertilizerWater(fertilizer int) int {
	return utils.GetOrDefault(a.fertilizerToWater, fertilizer, fertilizer)
}

func (a Almanac) GetWaterLight(water int) int {
	return utils.GetOrDefault(a.waterToLight, water, water)
}

func (a Almanac) GetLightTemperature(light int) int {
	return utils.GetOrDefault(a.lightToTemperature, light, light)
}

func (a Almanac) GetTemperatureHumidity(temp int) int {
	return utils.GetOrDefault(a.temperatureToHumidity, temp, temp)
}

func (a Almanac) GetHumidityLocation(humidity int) int {
	return utils.GetOrDefault(a.humidityToLocation, humidity, humidity)
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

func parseAlmanacMap(input []string, lineIdx int) (map[int]int, int) {
	if !isTitleLine(input[lineIdx]) {
		panic(fmt.Sprintf("expected title line to start map parsing: %s", input[lineIdx]))
	}

	idx := lineIdx + 1
	almanacMap := make(map[int]int)
	for ; idx < len(input); idx++ {
		if !isNumberLine(input[idx]) {
			break
		}
		nums := parseNumsLine(input[idx])

		srcNum := nums[1]
		dstNum := nums[0]
		numRange := nums[2]
		for i := 0; i < numRange; i++ {
			almanacMap[srcNum+i] = dstNum + i
		}
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

func isEmptyLine(line string) bool {
	return len(line) == 0
}
