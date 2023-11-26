package main

import (
	"strconv"
	"strings"
)

func Day08ParseInput(input string) Pair[string, map[string]Pair[string, string]] {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	input = strings.ReplaceAll(input, ",", "=") // Makes parsing a bit more pragmatic later on
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	navigationMap := make(map[string]Pair[string, string])
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		lineParts := strings.Split(line, "=")
		navigationMap[lineParts[0]] = Pair[string, string]{First: lineParts[1], Second: lineParts[2]}
	}

	return Pair[string, map[string]Pair[string, string]]{First: lines[0], Second: navigationMap}
}

func Day08Task0(input string) string {
	pair := Day08ParseInput(input)

	current := "AAA"
	steps := 0
	for current != "ZZZ" {
		direction := pair.First[steps%len(pair.First)]
		if direction == 'L' {
			current = pair.Second[current].First
		}
		if direction == 'R' {
			current = pair.Second[current].Second
		}
		steps++
	}

	return strconv.Itoa(steps)
}

func Day08Task1(input string) string {
	pair := Day08ParseInput(input)

	currents := make([]string, 0)

	for key := range pair.Second {
		if key[len(key)-1] == 'A' {
			currents = append(currents, key)
		}
	}

	loopSteps := make([]int, len(currents))
	for i := range currents {
		current := currents[i]
		steps := 0
		for current[len(current)-1] != 'Z' {
			direction := pair.First[steps%len(pair.First)]
			if direction == 'L' {
				current = pair.Second[current].First
			}
			if direction == 'R' {
				current = pair.Second[current].Second
			}
			steps++
		}
		loopSteps[i] = steps
	}

	getCommonLoopInterval := func(a, b int) int {
		greater := a
		lesser := b
		if b > a {
			greater = b
			lesser = a
		}
		ret := greater
		for ret%lesser != 0 {
			ret = ret + greater
		}
		return ret
	}

	ret := loopSteps[0]
	for i := 1; i < len(loopSteps); i++ {
		ret = getCommonLoopInterval(ret, loopSteps[i])
	}
	return strconv.Itoa(ret)
}
