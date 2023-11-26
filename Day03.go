package main

import (
	"strconv"
	"strings"
	"unicode"
)

type Day03PartNumber struct {
	Number int
	Y      int
	XFirst int
	XLast  int
}

func Day03GetPartNumbers(lines []string) [][]Day03PartNumber {

	partNumberLines := make([][]Day03PartNumber, len(lines))

	for y := range lines {
		line := lines[y]
		partNumberLine := make([]Day03PartNumber, 0)
		for x := 0; x < len(line); x++ {
			if unicode.IsDigit(rune(line[x])) {
				end := x
				for end < len(line) && unicode.IsDigit(rune(line[end])) {
					end = end + 1
				}
				number, _ := strconv.ParseInt(line[x:end], 10, 32)

				newPartNumber := Day03PartNumber{Number: int(number), Y: y, XFirst: x, XLast: end - 1}
				partNumberLine = append(partNumberLine, newPartNumber)

				x = end
			}
		}
		partNumberLines[y] = partNumberLine
	}

	return partNumberLines
}

func Day03Task0(input string) string {

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	partNumberLines := Day03GetPartNumbers(lines)

	sum := 0

	for y := range partNumberLines {
		line := lines[y]
		partNumberLine := partNumberLines[y]
		for x := range partNumberLine {
			partNumber := partNumberLine[x]
			start := partNumber.XFirst - 1
			end := partNumber.XLast + 1

			shouldCount := false
			if start >= 0 && line[start] != '.' && !unicode.IsDigit(rune(line[start])) ||
				end < len(line) && line[end] != '.' && !unicode.IsDigit(rune(line[end])) {
				shouldCount = true
			}
			for i := start; i <= end; i++ {
				if i >= 0 && i < len(line) &&
					(y > 0 && lines[y-1][i] != '.' && !unicode.IsDigit(rune(lines[y-1][i])) ||
						y < len(lines)-1 && lines[y+1][i] != '.' && !unicode.IsDigit(rune(lines[y+1][i]))) {
					shouldCount = true
				}
			}

			if shouldCount {
				sum = sum + partNumber.Number
			}
		}
	}

	return strconv.Itoa(sum)
}

func Day03Task1(input string) string {

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	partNumberLines := Day03GetPartNumbers(lines)

	sum := 0

	for y := range lines {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			if line[x] == '*' {
				adjacentPartNumbers := make([]Day03PartNumber, 0)

				appendAdjacentPartNumbers := func(partNumberLine []Day03PartNumber, appendTo []Day03PartNumber) []Day03PartNumber {
					for i := range partNumberLine {
						partNumber := partNumberLine[i]
						if partNumber.XLast >= x-1 && partNumber.XFirst <= x+1 {
							appendTo = append(appendTo, partNumber)
						}
					}
					return appendTo
				}

				if y > 0 {
					adjacentPartNumbers = appendAdjacentPartNumbers(partNumberLines[y-1], adjacentPartNumbers)
				}
				adjacentPartNumbers = appendAdjacentPartNumbers(partNumberLines[y], adjacentPartNumbers)
				if y < len(partNumberLines) {
					adjacentPartNumbers = appendAdjacentPartNumbers(partNumberLines[y+1], adjacentPartNumbers)
				}

				if len(adjacentPartNumbers) == 2 {
					sum = sum + adjacentPartNumbers[0].Number*adjacentPartNumbers[1].Number
				}
			}
		}
	}

	return strconv.Itoa(sum)
}
