package main

import (
	"strconv"
	"strings"
)

func Day09ParseInput(input string) [][]int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	sequences := make([][]int, len(lines))

	for i := range lines {
		line := lines[i]
		textNumbers := strings.Split(line, " ")
		numbers := make([]int, len(textNumbers))
		for i2 := range textNumbers {
			numbers[i2], _ = strconv.Atoi(textNumbers[i2])
		}
		sequences[i] = numbers
	}

	return sequences
}

func Day09GetNextNumber(pyramid [][]int, doLast bool) int {
	// Following is the All-Zeros recursion aborting condition
	allZero := true
	lastIndex := len(pyramid) - 1
	for i := range pyramid[lastIndex] {
		if pyramid[lastIndex][i] != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	// We're not there yet. So we generate the next sequence of differences and recurse into it
	nextDiffs := make([]int, len(pyramid[lastIndex])-1)

	for i := 1; i < len(pyramid[lastIndex]); i++ {
		nextDiffs[i-1] = pyramid[lastIndex][i] - pyramid[lastIndex][i-1]
	}

	pyramid = append(pyramid, nextDiffs)

	nextNumber := Day09GetNextNumber(pyramid, doLast)

	// Finally, decide which whether we need the next number in the beginning or end and calculate it accordingly
	if doLast {
		return pyramid[lastIndex][len(pyramid[lastIndex])-1] + nextNumber
	}

	return pyramid[lastIndex][0] - nextNumber
}

func Day09CommonTask(input string, doLast bool) string {
	sequences := Day09ParseInput(input)

	sum := 0

	for i := range sequences {
		// Start the recursion by interpreting the current sequence as the first and yet only line of the pyramid
		sequence := sequences[i]
		pyramid := make([][]int, 0)
		pyramid = append(pyramid, sequence)
		sum = sum + Day09GetNextNumber(pyramid, doLast)
	}

	return strconv.Itoa(sum)
}

func Day09Task0(input string) string { return Day09CommonTask(input, true) }

func Day09Task1(input string) string { return Day09CommonTask(input, false) }
