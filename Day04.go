package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day04ParseInput(input string) []int {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	matches := make([]int, len(lines))
	for i := range lines {
		line := lines[i]
		scratchCardSplit := strings.Split(line, ":")
		partsSplit := strings.Split(scratchCardSplit[1], "|")
		winningNumbers := make([]int, 0)
		gotNumbers := make([]int, 0)
		if len(partsSplit) == 2 {
			getNumbers := func(part string) []int {
				numbers := make([]int, 0)
				split := strings.Split(part, " ")
				for i2 := range split {
					item := split[i2]
					if item != "" {
						number, _ := strconv.Atoi(item)
						numbers = append(numbers, number)
					}
				}
				return numbers
			}
			winningNumbers = getNumbers(partsSplit[0])
			gotNumbers = getNumbers(partsSplit[1])
		}
		matches[i] = Day04CountMatchedNumbers(winningNumbers, gotNumbers)
	}
	return matches
}

func Day04CountMatchedNumbers(winningNumbers []int, gotNumbers []int) int {
	matchCount := 0
	for i2 := range gotNumbers {
		gotNumber := gotNumbers[i2]
		if slices.ContainsFunc(winningNumbers, func(n int) bool { return n == gotNumber }) {
			matchCount++
		}
	}
	return matchCount
}

func Day04CountCards(matches []int, start int) int {
	sum := 0

	matchedNumbersCount := matches[start]
	end := int(math.Min(float64(len(matches)), float64(start+matchedNumbersCount+1)))
	for i := start + 1; i < end; i++ {
		sum = sum + Day04CountCards(matches, i)
	}

	return 1 + sum
}

func Day04CommonTask(input string, getPoints func([]int, int) int) string {
	scratchCards := Day04ParseInput(input)

	sum := 0
	for i := range scratchCards {
		sum = sum + getPoints(scratchCards, i)
	}

	return strconv.Itoa(sum)
}

func Day04Task0(input string) string {
	return Day04CommonTask(input, func(matches []int, index int) int { return int(math.Pow(2.0, float64(matches[index])-1.0)) })
}

func Day04Task1(input string) string {
	return Day04CommonTask(input, func(matches []int, index int) int { return Day04CountCards(matches, index) })
}
