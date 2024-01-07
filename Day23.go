package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day23ParseInput0(input string) Pair[map[Pair[int, int]]rune, Pair[int, int]] {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	hikeMap := make(map[Pair[int, int]]rune)

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '#' {
				hikeMap[Pair[int, int]{x, y}] = rune(lines[y][x])
			}
		}
	}

	return Pair[map[Pair[int, int]]rune, Pair[int, int]]{hikeMap, Pair[int, int]{len(lines[0]) - 2, len(lines) - 1}}
}

func Day23ParseInput1(input string) Pair[map[Pair[int, int]]rune, Pair[int, int]] {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	hikeMap := make(map[Pair[int, int]]rune)

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '#' {
				hikeMap[Pair[int, int]{x, y}] = '.'
			}
		}
	}

	return Pair[map[Pair[int, int]]rune, Pair[int, int]]{hikeMap, Pair[int, int]{len(lines[0]) - 2, len(lines) - 1}}
}

func Day23GetMaxSteps(origin, current, target Pair[int, int], hikeMap map[Pair[int, int]]rune, intersections []Pair[int, int]) int {
	steps := 0
	for {
		steps++
		// Global aborting condition
		if current == target {
			return steps
		}
		// Local aborting condition
		if slices.Contains(intersections, current) {
			return math.MinInt
		}

		nextChoices := make([]Pair[Pair[int, int], bool], 0)

		left := Pair[int, int]{current.First - 1, current.Second}
		if char, ok := hikeMap[left]; ok && left != origin {
			nextChoices = append(nextChoices, Pair[Pair[int, int], bool]{left, char != '>'})
		}

		right := Pair[int, int]{current.First + 1, current.Second}
		if char, ok := hikeMap[right]; ok && right != origin {
			nextChoices = append(nextChoices, Pair[Pair[int, int], bool]{right, char != '<'})
		}

		up := Pair[int, int]{current.First, current.Second - 1}
		if char, ok := hikeMap[up]; ok && up != origin {
			nextChoices = append(nextChoices, Pair[Pair[int, int], bool]{up, char != 'v'})
		}

		down := Pair[int, int]{current.First, current.Second + 1}
		if char, ok := hikeMap[down]; ok && down != origin && char != '^' {
			nextChoices = append(nextChoices, Pair[Pair[int, int], bool]{down, char != '^'})
		}

		origin = current
		if len(nextChoices) > 1 {
			maxSteps := math.MinInt
			intersections = append(intersections, current)
			for i := range nextChoices {
				currentChoice := nextChoices[i]
				if currentChoice.Second {
					currentSteps := Day23GetMaxSteps(origin, currentChoice.First, target, hikeMap, intersections) + steps
					if currentSteps > maxSteps {
						maxSteps = currentSteps
					}
				}
			}
			return maxSteps
		} else if len(nextChoices) == 1 {
			if nextChoices[0].Second {
				current = nextChoices[0].First
			} else {
				return math.MinInt
			}
		} else {
			return math.MinInt
		}
	}
}

func Day23Task0(input string) string {
	pair := Day23ParseInput0(input)

	origin := Pair[int, int]{1, 0}
	current := Pair[int, int]{1, 1}
	target := pair.Second
	hikeMap := pair.First

	return strconv.Itoa(Day23GetMaxSteps(origin, current, target, hikeMap, make([]Pair[int, int], 0)))
}

func Day23Task1(input string) string {
	pair := Day23ParseInput1(input)

	origin := Pair[int, int]{1, 0}
	current := Pair[int, int]{1, 1}
	target := pair.Second
	hikeMap := pair.First

	return strconv.Itoa(Day23GetMaxSteps(origin, current, target, hikeMap, make([]Pair[int, int], 0)))
}
