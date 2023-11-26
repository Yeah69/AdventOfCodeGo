package main

import (
	"slices"
	"strconv"
	"strings"
)

type Day05MapItem struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

type Day05Almanac struct {
	Seeds []int
	Maps  [][]Day05MapItem
}

func Day05ParseInput(input string) Day05Almanac {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	parts := strings.Split(input, "\n\n")

	seedsTexts := strings.Split(strings.ReplaceAll(parts[0], "seeds: ", ""), " ")
	seeds := make([]int, len(seedsTexts))
	for i := range seedsTexts {
		seeds[i], _ = strconv.Atoi(seedsTexts[i])
	}

	maps := make([][]Day05MapItem, len(parts)-1)
	for i := 1; i < len(parts); i++ {
		mapPart := parts[i]
		itemParts := strings.Split(mapPart, "\n")
		items := make([]Day05MapItem, len(itemParts)-1)
		for i2 := 1; i2 < len(itemParts); i2++ {
			data := strings.Split(itemParts[i2], " ")
			destinationStart, _ := strconv.Atoi(data[0])
			sourceStart, _ := strconv.Atoi(data[1])
			rangeData, _ := strconv.Atoi(data[2])
			items[i2-1] = Day05MapItem{DestinationStart: destinationStart, SourceStart: sourceStart, Range: rangeData}
		}
		maps[i-1] = items
	}

	return Day05Almanac{Seeds: seeds, Maps: maps}
}

func Day05Task0(input string) string {
	almanac := Day05ParseInput(input)

	locations := make([]int, len(almanac.Seeds))
	for i := range almanac.Seeds {
		current := almanac.Seeds[i]
		for i2 := range almanac.Maps {
			mapItems := almanac.Maps[i2]
			for i3 := range mapItems {
				mapItem := mapItems[i3]
				if current >= mapItem.SourceStart && current < mapItem.SourceStart+mapItem.Range {
					current = mapItem.DestinationStart + current - mapItem.SourceStart
					break
				}
			}
		}
		locations[i] = current
	}

	return strconv.Itoa(slices.Min(locations))
}

func Day05Task1(input string) string {
	almanac := Day05ParseInput(input)

	initialSeeds := make([]Pair[int, int], len(almanac.Seeds)/2)
	for i := 0; i < len(almanac.Seeds); i = i + 2 {
		initialSeeds[i/2] = Pair[int, int]{First: almanac.Seeds[i], Second: almanac.Seeds[i+1]}
	}

	potentialMin := 0
	for {
		current := potentialMin

		for i := len(almanac.Maps) - 1; i >= 0; i-- {
			mapItems := almanac.Maps[i]
			for i2 := range mapItems {
				mapItem := mapItems[i2]
				if current >= mapItem.DestinationStart && current < mapItem.DestinationStart+mapItem.Range {
					current = mapItem.SourceStart + current - mapItem.DestinationStart
					break
				}
			}
		}

		for i := range initialSeeds {
			initialSeed := initialSeeds[i]
			if current >= initialSeed.First && current < initialSeed.First+initialSeed.Second {
				return strconv.Itoa(potentialMin)
			}
		}

		potentialMin++
	}
}
