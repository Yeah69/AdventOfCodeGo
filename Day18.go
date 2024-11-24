package main

import (
	"strconv"
	"strings"
)

type Day18PlanItem struct {
	Direction int
	Steps     int
}

func Day18ParseInput0(input string) []Day18PlanItem {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	plan := make([]Day18PlanItem, len(lines))

	for i := range lines {
		parts := strings.Split(lines[i], " ")
		steps, _ := strconv.Atoi(parts[1])
		direction := 0
		if parts[0][0] == 'D' {
			direction = 1
		}
		if parts[0][0] == 'L' {
			direction = 2
		}
		if parts[0][0] == 'U' {
			direction = 3
		}
		plan[i] = Day18PlanItem{direction, steps}
	}

	return plan
}

func Day18ParseInput1(input string) []Day18PlanItem {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	plan := make([]Day18PlanItem, len(lines))

	for i := range lines {
		color := strings.Split(lines[i], " ")[2]
		direction, _ := strconv.Atoi(color[len(color)-2 : len(color)-1])
		steps, _ := strconv.ParseInt(color[2:len(color)-2], 16, 64)
		plan[i] = Day18PlanItem{direction, int(steps)}
	}

	return plan
}

func Day18CommonTask(plan []Day18PlanItem) string {
	corners := make([]Pair[int, int], len(plan))
	currentPosition := Pair[int, int]{First: 0, Second: 0}

	stepsLength := 0

	for i := range plan {
		planItem := plan[i]
		steps := planItem.Steps
		switch planItem.Direction {
		case 0:
			currentPosition = Pair[int, int]{currentPosition.First + steps, currentPosition.Second}
		case 1:
			currentPosition = Pair[int, int]{currentPosition.First, currentPosition.Second + steps}
		case 2:
			currentPosition = Pair[int, int]{First: currentPosition.First - steps, Second: currentPosition.Second}
		case 3:
			currentPosition = Pair[int, int]{First: currentPosition.First, Second: currentPosition.Second - steps}
		}
		corners[((i + 1) % len(corners))] = currentPosition
		stepsLength = stepsLength + steps
	}

	area := 0

	for i := range corners {
		j := (i + 1) % len(corners)
		area = area + (corners[i].First * corners[j].Second) - (corners[j].First * corners[i].Second)
	}
	area = area / 2

	return strconv.Itoa(area + stepsLength/2 + 1)
}

func Day18Task0(input string) string {
	return Day18CommonTask(Day18ParseInput0(input))
}

func Day18Task1(input string) string {
	return Day18CommonTask(Day18ParseInput1(input))
}
