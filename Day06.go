package main

import (
	"math"
	"strconv"
	"strings"
)

func Day06ParseInput0(input string) Pair[[]int, []int] {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	parts := strings.Split(input, "\n")

	getNumbers := func(text string) []int {
		parts := strings.Split(text, " ")
		numbers := make([]int, 0)
		for i := range parts {
			part := parts[i]
			if part != "" {
				number, _ := strconv.Atoi(part)
				numbers = append(numbers, number)
			}
		}
		return numbers
	}

	times := getNumbers(strings.ReplaceAll(parts[0], "Time:", ""))
	distances := getNumbers(strings.ReplaceAll(parts[1], "Distance:", ""))

	return Pair[[]int, []int]{times, distances}
}

func Day06ParseInput1(input string) Pair[int, int] {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, " ", "")
	parts := strings.Split(input, "\n")

	time, _ := strconv.Atoi(strings.ReplaceAll(parts[0], "Time:", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(parts[1], "Distance:", ""))

	return Pair[int, int]{time, distance}
}

// Day06CalculateWins charge := time/2 +- sqrt(time^2/4-distance) is the exact charge for the current record distance so any charge in between wins
func Day06CalculateWins(time int, distance int) int {
	middle := float64(time) / 2.0
	plusMinus := math.Sqrt(float64(time*time)/4.0 - float64(distance))
	minCharge := int(math.Max(math.Floor(middle-plusMinus), -1.0)) + 1
	maxCharge := int(math.Min(math.Ceil(middle+plusMinus), float64(time+1))) - 1
	return maxCharge - minCharge + 1
}

func Day06Task0(input string) string {
	pair := Day06ParseInput0(input)
	product := 1
	for i := range pair.First {
		product = product * Day06CalculateWins(pair.First[i], pair.Second[i])
	}
	return strconv.Itoa(product)
}

func Day06Task1(input string) string {
	pair := Day06ParseInput1(input)
	return strconv.Itoa(Day06CalculateWins(pair.First, pair.Second))
}
