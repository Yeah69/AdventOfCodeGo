package main

import (
	"strconv"
	"strings"
)

func Day17ParseInput(input string) [][]int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	textLines := strings.Split(input, "\n")

	lines := make([][]int, len(textLines))

	for i := range textLines {
		textLine := textLines[i]

		line := make([]int, len(textLine))

		for i2 := range textLine {
			line[i2] = int(textLine[i2] - '0')
		}

		lines[i] = line
	}

	return lines
}

type Day17State struct {
	X              int
	Y              int
	HorizontalNext bool
}

func Day17CommonTask(input string, minSteps, maxSteps int) string {
	cityMap := Day17ParseInput(input)

	queue := make([]Day17State, 0)

	firstHorizontalNext := Day17State{X: 0, Y: 0, HorizontalNext: true}
	firstVerticalNext := Day17State{X: 0, Y: 0, HorizontalNext: false}
	queue = append(queue, firstHorizontalNext)
	queue = append(queue, firstVerticalNext)

	cache := make(map[Day17State]int)
	cache[firstHorizontalNext] = 0
	cache[firstVerticalNext] = 0

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]
		currentHeatLoss := cache[currentState]
		nextDirection := !currentState.HorizontalNext

		nextSteps := func(negative bool) {
			intermediateSum := 0
			for i := 1; i <= maxSteps; i++ {
				actualI := i
				if negative {
					actualI = -i
				}
				nextX := currentState.X
				nextY := currentState.Y
				if currentState.HorizontalNext {
					nextX = nextX + actualI
				} else {
					nextY = nextY + actualI
				}
				if nextX >= 0 && nextY >= 0 && nextX < len(cityMap[0]) && nextY < len(cityMap) {
					intermediateSum = intermediateSum + cityMap[nextY][nextX]
					if i >= minSteps {
						nextState := Day17State{nextX, nextY, nextDirection}
						if value, ok := cache[nextState]; !ok || ok && value > currentHeatLoss+intermediateSum {
							cache[nextState] = currentHeatLoss + intermediateSum
							queue = append(queue, nextState)
						}
					}
				}
			}
		}

		nextSteps(false)
		nextSteps(true)
	}

	lastX := len(cityMap[0]) - 1
	lastY := len(cityMap) - 1

	horizontalValue, horizontalOk := cache[Day17State{lastX, lastY, false}]
	verticalValue, verticalOk := cache[Day17State{lastX, lastY, true}]

	if horizontalOk && verticalOk {
		if verticalValue < horizontalValue {
			return strconv.Itoa(verticalValue)
		} else {
			return strconv.Itoa(horizontalValue)
		}
	} else if horizontalOk {
		return strconv.Itoa(horizontalValue)
	} else if verticalOk {
		return strconv.Itoa(verticalValue)
	}
	return NoSolutionFound
}

func Day17Task0(input string) string {
	return Day17CommonTask(input, 1, 3)
}

func Day17Task1(input string) string {
	return Day17CommonTask(input, 4, 10)
}
