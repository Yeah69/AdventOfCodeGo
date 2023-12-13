package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day12SpringsRow struct {
	Description   string
	CheckSequence []int
}

func Day12ParseInput(input string) []Day12SpringsRow {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	report := make([]Day12SpringsRow, 0)
	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, " ")
		checks := make([]int, 0)
		sequence := strings.Split(parts[1], ",")
		for i2 := range sequence {
			num, _ := strconv.Atoi(sequence[i2])
			checks = append(checks, num)
		}
		report = append(report, Day12SpringsRow{
			Description:   parts[0],
			CheckSequence: checks})
	}
	return report
}

func Day12GetArrangementCount(description []rune, checkSequence []int, cache map[string]int) int {
	if len(checkSequence) == 0 {
		if slices.Contains(description, '#') {
			return 0
		} else {
			return 1
		}
	}
	checkSequenceSum := len(checkSequence) - 1
	for i := range checkSequence {
		checkSequenceSum = checkSequenceSum + checkSequence[i]
	}
	if checkSequenceSum > len(description) {
		return 0
	}

	start := 0
	end := len(description) - 1
	for start < len(description) && description[start] == '.' {
		start++
	}
	if start == len(description) {
		return 0
	}
	for description[end] == '.' {
		end--
	}

	description = description[start : end+1]
	key := fmt.Sprintf("%s;%d", string(description), checkSequenceSum)

	if v, ok := cache[key]; ok {
		return v
	}

	possiblePositions := make([]int, 0)
	for i := range description {
		if len(description)-i < checkSequence[0] || i > 0 && description[i-1] == '#' {
			break
		}
		possible := true
		e := i + checkSequence[0]
		for i2 := i; i2 < e; i2++ {
			if description[i2] == '.' {
				possible = false
				break
			}
		}
		if possible && i > 0 && description[i-1] == '#' ||
			len(description) > checkSequence[0]+i && description[checkSequence[0]+i] == '#' {
			possible = false
		}
		if possible {
			possiblePositions = append(possiblePositions, i)
		}
	}

	sum := 0

	for i := len(possiblePositions) - 1; i >= 0; i-- {
		newDescription := description[possiblePositions[i]+checkSequence[0]:]
		if len(newDescription) > 0 {
			newDescription = newDescription[1:]
		}
		summand := Day12GetArrangementCount(newDescription, checkSequence[1:], cache)
		sum = sum + summand
	}

	cache[key] = sum
	return sum
}

func Day12Task0(input string) string {
	springsReports := Day12ParseInput(input)

	sum := 0

	for i := range springsReports {
		summand := Day12GetArrangementCount([]rune(springsReports[i].Description), springsReports[i].CheckSequence, make(map[string]int))
		sum = sum + summand
	}

	return strconv.Itoa(sum)
}

func Day12Task1(input string) string {
	springsReports := Day12ParseInput(input)

	sum := 0

	for i := range springsReports {
		reportDescription := springsReports[i].Description
		reportCheckSequence := springsReports[i].CheckSequence
		description := ""
		checkSequence := make([]int, len(reportCheckSequence)*5)
		for i := 0; i < 5; i++ {
			if i == 0 {
				description = description + reportDescription
			} else {
				description = description + "?" + reportDescription
			}
			for i2 := range reportCheckSequence {
				checkSequence[i*len(reportCheckSequence)+i2] = reportCheckSequence[i2]
			}
		}

		sum = sum + Day12GetArrangementCount([]rune(description), checkSequence, make(map[string]int))
	}

	return strconv.Itoa(sum)
}
