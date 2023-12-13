package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day12SpringsRow struct {
	Description          string
	CheckSequence        []int
	ShouldCheckMarkCount int
	SetCheckMarkCount    int
	QuestionMarkCount    int
	CheckMarkDiff        int
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
		shouldCheckMarkCount := 0
		for i2 := range sequence {
			num, _ := strconv.Atoi(sequence[i2])
			checks = append(checks, num)
			shouldCheckMarkCount = shouldCheckMarkCount + num
		}
		setCheckMarkCount := 0
		questionMarkCount := 0
		for i2 := range parts[0] {
			if parts[0][i2] == '#' {
				setCheckMarkCount++
			}
			if parts[0][i2] == '?' {
				questionMarkCount++
			}
		}
		report = append(report, Day12SpringsRow{
			Description:          parts[0],
			CheckSequence:        checks,
			ShouldCheckMarkCount: shouldCheckMarkCount,
			SetCheckMarkCount:    setCheckMarkCount,
			QuestionMarkCount:    questionMarkCount,
			CheckMarkDiff:        shouldCheckMarkCount - setCheckMarkCount})
	}
	return report
}

func Day12GetArrangementCount(description []rune, checkSequence []int, questionMarkCount, checkMarkDiff int) int {
	/*if checkMarkDiff == 0 {
		text := string(description)
		text = strings.ReplaceAll(text, "?", ".")
		lengths := make([]int, 0)
		split := strings.Split(text, ".")
		for i := range split {
			if len(split[i]) > 0 {
				lengths = append(lengths, len(split[i]))
			}
		}
		if !slices.Equal(lengths, checkSequence) {
			return 0
		}
		return 1
	}

	if questionMarkCount < checkMarkDiff {
		return 0
	}

	nextQuestionMarkIndex := slices.Index(description, '?')

	description[nextQuestionMarkIndex] = '.'
	sum := Day12GetArrangementCount(description, checkSequence, questionMarkCount-1, checkMarkDiff)

	defer func() { description[nextQuestionMarkIndex] = '?' }()
	description[nextQuestionMarkIndex] = '#'

	part := string(description[:nextQuestionMarkIndex+1])
	lengths := make([]int, 0)
	split := strings.Split(part, ".")
	for i := range split {
		if len(split[i]) > 0 {
			lengths = append(lengths, len(split[i]))
		}
	}
	if len(lengths) > len(checkSequence) {
		return sum
	}
	for i := 0; i < len(lengths)-1; i++ {
		if checkSequence[i] != lengths[i] {
			return sum
		}
	}
	if lengths[len(lengths)-1] > checkSequence[len(lengths)-1] {
		return sum
	}

	sum = sum + Day12GetArrangementCount(description, checkSequence, questionMarkCount-1, checkMarkDiff-1)

	return sum*/

	text := string(description)
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
	text = string(description)

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
		summand := Day12GetArrangementCount(newDescription, checkSequence[1:], 0, 0)
		sum = sum + summand
	}
	_ = text

	return sum //*/

	// OLD
	/*
		if description[0] == '#' {
			if len(description) < checkSequence[0] {
				return 0
			}

			for i := 0; i < checkSequence[0]; i++ {
				if description[i] == '.' {
					return 0
				}
			}
			if len(description) > checkSequence[0] && description[checkSequence[0]] == '#' {
				return 0
			}

			if len(description) == checkSequence[0] {
				return Day12GetArrangementCount(description[checkSequence[0]:], checkSequence[1:], 0, 0)
			} else {
				return Day12GetArrangementCount(description[checkSequence[0]+1:], checkSequence[1:], 0, 0)
			}

		}

		sum := 0

		for len(description) > 0 && description[0] == '?' {
			if len(description) < checkSequence[0] {
				break
			}
			fits := true
			for i := 0; i < checkSequence[0]; i++ {
				if description[i] == '.' {
					fits = false
					break
				}
			}
			if !fits {
				break
			}
			if len(description) > checkSequence[0] && description[checkSequence[0]] == '#' {
				description = description[1:]
				continue
			}

			if len(description) == checkSequence[0] {
				sum = sum + Day12GetArrangementCount(description[checkSequence[0]:], checkSequence[1:], 0, 0)
			} else {
				sum = sum + Day12GetArrangementCount(description[checkSequence[0]+1:], checkSequence[1:], 0, 0)
			}

			description = description[1:]
			text = string(description)
		}

		start = 0
		for start < len(description) && description[start] == '?' {
			start++
		}

		_ = text

		return sum + Day12GetArrangementCount(description[start:], checkSequence, 0, 0) //*/
}

func Day12Task0(input string) string {
	springsReports := Day12ParseInput(input)

	sum := 0

	for i := range springsReports {
		summand := Day12GetArrangementCount([]rune(springsReports[i].Description), springsReports[i].CheckSequence, springsReports[i].QuestionMarkCount, springsReports[i].CheckMarkDiff)
		sum = sum + summand
	}

	return strconv.Itoa(sum)
}

func Day12Task1(input string) string {
	springsReports := Day12ParseInput(input)

	sum := 0

	for i := range springsReports {
		fmt.Println(i)
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

		sum = sum + Day12GetArrangementCount([]rune(description), checkSequence, springsReports[i].QuestionMarkCount*5+4, springsReports[i].CheckMarkDiff*5)
	}

	return strconv.Itoa(sum)
}
