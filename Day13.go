package main

import (
	"math"
	"strconv"
	"strings"
)

func Day13ParseInput(input string) [][]string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	maps := strings.Split(input, "\n\n")

	mapsByLines := make([][]string, len(maps))
	for i := range maps {
		mapsByLines[i] = strings.Split(maps[i], "\n")
	}
	return mapsByLines
}

func Day13GetHorizontalMirror(image []string, ignore int) int {
	for i := 0; i < len(image)-1; i++ {
		if i == ignore {
			continue
		}
		if image[i] == image[i+1] {
			distanceLeft := int(math.Abs(float64(len(image) - (i + 1))))
			distance := int(math.Min(float64(distanceLeft), float64(i+1)))
			checks := true
			for i2 := 1; i2 < distance; i2++ {
				if image[i-i2] != image[i+1+i2] {
					checks = false
					break
				}
			}
			if checks {
				return i + 1
			}
		}
	}
	return 0
}

func Day13GetVerticalMirror(image []string, ignore int) int {
	for i := 0; i < len(image[0])-1; i++ {
		if i == ignore {
			continue
		}
		areColumnsEqual := func(leftColumn, rightColumn int) bool {
			for row := range image {
				if image[row][leftColumn] != image[row][rightColumn] {
					return false
				}
			}
			return true
		}

		if areColumnsEqual(i, i+1) {
			distanceLeft := int(math.Abs(float64(len(image[0]) - (i + 1))))
			distance := int(math.Min(float64(distanceLeft), float64(i+1)))
			checks := true
			for i2 := 0; i2 < distance; i2++ {
				if !areColumnsEqual(i-i2, i+1+i2) {
					checks = false
					break
				}
			}
			if checks {
				return i + 1
			}
		}
	}
	return 0
}

func Day13CommonPart(image []string, ignoreRowIndex, ignoreColumnIndex int) Pair[int, int] {
	row := Day13GetHorizontalMirror(image, ignoreRowIndex)
	column := Day13GetVerticalMirror(image, ignoreColumnIndex)
	return Pair[int, int]{First: row * 100, Second: column}
}

func Day13Task0(input string) string {
	maps := Day13ParseInput(input)

	sum := 0

	for i := range maps {
		pair := Day13CommonPart(maps[i], -1, -1)
		sum = sum + pair.First + pair.Second
	}

	return strconv.Itoa(sum)
}

func Day13Task1(input string) string {
	maps := Day13ParseInput(input)

	sum := 0

	for i := range maps {
		originalPair := Day13CommonPart(maps[i], -1, -1)
		ignoreRowIndex := -1
		if originalPair.First > 0 {
			ignoreRowIndex = originalPair.First/100 - 1
		}
		ignoreColumnIndex := -1
		if originalPair.Second > 0 {
			ignoreColumnIndex = originalPair.Second - 1
		}
		image := maps[i]
		for y := range image {
			shouldBreak := false
			row := image[y]
			for x := range row {
				char := '#'
				if row[x] == '#' {
					char = '.'
				}
				changeChar := func(sub string, char rune, position int) string {
					ret := []rune(sub)
					ret[position] = char
					return string(ret)
				}
				image[y] = changeChar(row, char, x)
				currentPair := Day13CommonPart(image, ignoreRowIndex, ignoreColumnIndex)
				if currentPair.First != 0 && currentPair.First != originalPair.First {
					sum = sum + currentPair.First
					shouldBreak = true
					break
				}
				if currentPair.Second != 0 && currentPair.Second != originalPair.Second {
					sum = sum + currentPair.Second
					shouldBreak = true
					break
				}
				image[y] = row
			}
			if shouldBreak {
				break
			}
		}
	}

	return strconv.Itoa(sum)
}
