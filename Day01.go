package main

import (
	"strconv"
	"strings"
	"unicode"
)

type Day01LiteralDigit struct {
	DigitAsString      string
	DigitLiteral       string
	DigitLiteralLength int
}

var Day01LiteralDigits = [9]Day01LiteralDigit{
	{DigitAsString: "1", DigitLiteral: "one", DigitLiteralLength: 3},
	{DigitAsString: "2", DigitLiteral: "two", DigitLiteralLength: 3},
	{DigitAsString: "3", DigitLiteral: "three", DigitLiteralLength: 5},
	{DigitAsString: "4", DigitLiteral: "four", DigitLiteralLength: 4},
	{DigitAsString: "5", DigitLiteral: "five", DigitLiteralLength: 4},
	{DigitAsString: "6", DigitLiteral: "six", DigitLiteralLength: 3},
	{DigitAsString: "7", DigitLiteral: "seven", DigitLiteralLength: 5},
	{DigitAsString: "8", DigitLiteral: "eight", DigitLiteralLength: 5},
	{DigitAsString: "9", DigitLiteral: "nine", DigitLiteralLength: 4},
}

func Day01Task0GetDigit(line string, index int) string {
	char := line[index]
	if unicode.IsDigit(rune(char)) {
		ret := line[index : index+1]
		return ret
	}
	return ""
}

func Day01Task1GetDigit(line string, index int) string {
	digit := Day01Task0GetDigit(line, index)
	if digit != "" {
		return digit
	}
	length := len(line)
	remainingLength := length - index

	for i := range Day01LiteralDigits {
		literalDigit := Day01LiteralDigits[i]
		if remainingLength >= literalDigit.DigitLiteralLength && line[index:index+literalDigit.DigitLiteralLength] == literalDigit.DigitLiteral {
			return literalDigit.DigitAsString
		}
	}

	return ""
}

func Day01CommonTask(input string, getDigit func(string, int) string) string {
	getFirst := func(line string, getIndex func(int) int) string {
		for j := range line {
			char := getDigit(line, getIndex(j))
			if char != "" {
				return char
			}
		}
		return ""
	}

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	var sum int64 = 0
	for i := range lines {
		line := lines[i]
		length := len(line)
		first := getFirst(line, func(x int) int { return x })
		last := getFirst(line, func(x int) int { return length - x - 1 })
		num, _ := strconv.ParseInt(first+last, 10, 32)
		sum = sum + num
	}

	return strconv.FormatInt(sum, 10)
}

func Day01Task0(input string) string {
	return Day01CommonTask(input, Day01Task0GetDigit)
}

func Day01Task1(input string) string {
	return Day01CommonTask(input, Day01Task1GetDigit)
}
