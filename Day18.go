package main

import (
	"math"
	"slices"
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

func Day18CommonTask3(plan []Day18PlanItem) string {
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
		if corners[i].First != corners[((i+1)%len(corners))].First {
			stepsLength = stepsLength + steps
		}
	}

	area := 0

	for j := range corners {
		i := j - 1
		if i == -1 {
			i = len(corners) - 1
		}
		area = area + corners[i].First*corners[j].Second - corners[j].First*corners[i].Second
	}

	return strconv.Itoa(area / 2)
}

func Day18CommonTask2(plan []Day18PlanItem) string {
	corners := make([]Pair[int, int], len(plan))
	currentPosition := Pair[int, int]{First: 0, Second: 0}

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
	}

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for i := range corners {
		pair := corners[i]
		if pair.First < minX {
			minX = pair.First
		}
		if pair.First > maxX {
			maxX = pair.First
		}
		if pair.Second < minY {
			minY = pair.Second
		}
		if pair.Second > maxY {
			maxY = pair.Second
		}
	}

	slices.SortFunc(corners, func(a, b Pair[int, int]) int {
		if a.Second < b.Second {
			return -1
		}
		if a.Second > b.Second {
			return 1
		}
		if a.First < b.First {
			return -1
		}
		if a.First > b.First {
			return 1
		}
		return 0
	})

	area := 0
	lines := make([]Pair[int, int], 0)
	linesSum := 0

	lastLine := math.MinInt

	for len(corners) > 0 {
		currentLine := corners[0].Second
		split := 1
		for split < len(corners) && corners[split].Second == currentLine {
			split++
		}

		area = area + linesSum*(currentLine-lastLine)

		onLine := corners[:split]
		corners = corners[split:]

		for i := 0; i < len(onLine); i = i + 2 {
			left := onLine[i].First
			right := onLine[i+1].First
			additionIndex := slices.IndexFunc(lines, func(pair Pair[int, int]) bool { return pair.First == right || pair.Second == left })
			subtractionIndex := slices.IndexFunc(lines, func(pair Pair[int, int]) bool { return pair.First == left || pair.Second == right })
			if additionIndex != -1 {
				leftMatchIndex := slices.IndexFunc(lines, func(pair Pair[int, int]) bool { return pair.Second == left })
				rightMatchIndex := slices.IndexFunc(lines, func(pair Pair[int, int]) bool { return pair.First == right })
				if leftMatchIndex != -1 && rightMatchIndex != -1 {
					lines[leftMatchIndex] = Pair[int, int]{lines[leftMatchIndex].First, lines[rightMatchIndex].Second}
					lines = slices.Delete(lines, rightMatchIndex, rightMatchIndex+1)
				} else if leftMatchIndex != -1 {
					lines[leftMatchIndex] = Pair[int, int]{lines[leftMatchIndex].First, right}
				} else {
					lines[rightMatchIndex] = Pair[int, int]{left, lines[rightMatchIndex].Second}
				}
			}
			if subtractionIndex != -1 {
				line := lines[subtractionIndex]
				if line.First == left && line.Second == right {
					lines = slices.Delete(lines, subtractionIndex, subtractionIndex+1)
					area = area + right - left + 1
				} else if line.First == left {
					lines[subtractionIndex] = Pair[int, int]{right, line.Second}
					area = area + right - left
				} else {
					lines[subtractionIndex] = Pair[int, int]{line.First, left}
					area = area + right - left
				}
			}
			if additionIndex == -1 && subtractionIndex == -1 {
				lines = append(lines, Pair[int, int]{left, right})
			}
		}

		linesSum = 0
		for i := range lines {
			linesSum = linesSum + lines[i].Second - lines[i].First + 1
		}

		lastLine = currentLine
	}

	return strconv.Itoa(area)
}

func Day18CommonTask(plan []Day18PlanItem) string {
	corners := make([]Pair[int, int], len(plan))
	currentPosition := Pair[int, int]{First: 0, Second: 0}

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
	}

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for i := range corners {
		pair := corners[i]
		if pair.First < minX {
			minX = pair.First
		}
		if pair.First > maxX {
			maxX = pair.First
		}
		if pair.Second < minY {
			minY = pair.Second
		}
		if pair.Second > maxY {
			maxY = pair.Second
		}
	}

	wholeArea := (maxX - minX + 1) * (maxY - minY + 1)

	cutout := 0

	i := 0
	left := 0

	for corners[i].First < maxX {
		nextI := (i + 1) % len(corners)
		evenNextI := (nextI + 1) % len(corners)
		nextCorner := corners[nextI]
		evenNextCorner := corners[evenNextI]

		right := nextCorner.First // assume left turn
		if nextCorner.Second < evenNextCorner.Second {
			right++ // but was right turn
		}
		cutoutIncrement := (right - left) * (nextCorner.Second - minY)
		cutout = cutout + cutoutIncrement
		i = evenNextI
		left = right
	}

	i = (i - 1) % len(corners)
	left = corners[i].Second

	for corners[i].Second < maxY {
		nextI := (i + 1) % len(corners)
		evenNextI := (nextI + 1) % len(corners)
		nextCorner := corners[nextI]
		evenNextCorner := corners[evenNextI]

		right := nextCorner.Second // assume left turn
		if nextCorner.First > evenNextCorner.First {
			right++ // but was right turn
		}
		cutoutIncrement := (right - left) * (maxX - nextCorner.First)
		cutout = cutout + cutoutIncrement
		i = evenNextI
		left = right
	}

	i = (i - 1) % len(corners)
	left = corners[i].First

	for corners[i].First > minX {
		nextI := (i + 1) % len(corners)
		evenNextI := (nextI + 1) % len(corners)
		nextCorner := corners[nextI]
		evenNextCorner := corners[evenNextI]

		right := nextCorner.First // assume left turn
		if nextCorner.Second > evenNextCorner.Second {
			right-- // but was right turn
		}
		cutoutIncrement := (left - right) * (maxY - nextCorner.Second)
		cutout = cutout + cutoutIncrement
		i = evenNextI
		left = right
	}

	i = (i - 1) % len(corners)
	left = corners[i].Second

	for corners[i].Second > minY {
		nextI := (i + 1) % len(corners)
		evenNextI := (nextI + 1) % len(corners)
		nextCorner := corners[nextI]
		evenNextCorner := corners[evenNextI]

		right := nextCorner.Second // assume left turn
		if nextCorner.First < evenNextCorner.First {
			right-- // but was right turn
		}
		cutoutIncrement := (left - right) * (nextCorner.First - minX)
		cutout = cutout + cutoutIncrement
		i = evenNextI
		left = right
	}

	return strconv.Itoa(wholeArea - cutout)
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

func Day18Task0(input string) string {
	return Day18CommonTask3(Day18ParseInput0(input))
	plan := Day18ParseInput0(input)

	border := make(map[Pair[int, int]]bool)
	currentPosition := Pair[int, int]{First: 0, Second: 0}
	border[currentPosition] = true

	makeStepRight := func(cur Pair[int, int]) Pair[int, int] {
		return Pair[int, int]{First: cur.First + 1, Second: cur.Second}
	}
	makeStepDown := func(cur Pair[int, int]) Pair[int, int] {
		return Pair[int, int]{First: cur.First, Second: cur.Second + 1}
	}
	makeStepLeft := func(cur Pair[int, int]) Pair[int, int] {
		return Pair[int, int]{First: cur.First - 1, Second: cur.Second}
	}
	makeStepUp := func(cur Pair[int, int]) Pair[int, int] {
		return Pair[int, int]{First: cur.First, Second: cur.Second - 1}
	}

	for i := range plan {
		planItem := plan[i]
		makeStep := makeStepRight
		switch planItem.Direction {
		case 1:
			makeStep = makeStepDown
		case 2:
			makeStep = makeStepLeft
		case 3:
			makeStep = makeStepUp
		}
		tempPosition := currentPosition
		for i := 0; i < planItem.Steps; i++ {
			tempPosition = makeStep(tempPosition)
			border[tempPosition] = true
		}
		currentPosition = tempPosition
	}

	queue := make([]Pair[int, int], 0)
	inside := make(map[Pair[int, int]]bool)

	queue = append(queue, Pair[int, int]{1, 1})

	for len(queue) > 0 {
		current := queue[0]
		if _, ok := inside[current]; !ok {
			if _, okBorder := border[current]; !okBorder {
				inside[current] = true
				queue = append(queue, Pair[int, int]{current.First + 1, current.Second})
				queue = append(queue, Pair[int, int]{current.First, current.Second + 1})
				queue = append(queue, Pair[int, int]{current.First - 1, current.Second})
				queue = append(queue, Pair[int, int]{current.First, current.Second - 1})
			}
		}
		queue = queue[1:]
	}

	return strconv.Itoa(len(border) + len(inside))
}

func Day18Task1(input string) string {
	return Day18CommonTask3(Day18ParseInput1(input))
	plan := Day18ParseInput1(input)

	corners := make([]Pair[int, int], len(plan))
	currentPosition := Pair[int, int]{First: 0, Second: 0}

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
	}

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for i := range corners {
		pair := corners[i]
		if pair.First < minX {
			minX = pair.First
		}
		if pair.First > maxX {
			maxX = pair.First
		}
		if pair.Second < minY {
			minY = pair.Second
		}
		if pair.Second > maxY {
			maxY = pair.Second
		}
	}

	wholeArea := (maxX - minX + 1) * (maxY - minY + 1)

	cutout := 0

	i := 0
	left := 0

	for corners[i].First < maxX {
		nextI := (i + 1) % len(corners)
		evenNextI := (nextI + 1) % len(corners)
		nextCorner := corners[nextI]
		evenNextCorner := corners[evenNextI]
		right := nextCorner.First
		if evenNextCorner.Second > nextCorner.Second {
			right++
		}
		cutoutIncrement := (right - left) * (nextCorner.Second - minY)
		cutout = cutout + cutoutIncrement
		i = evenNextI
	}

	return strconv.Itoa(wholeArea - cutout)
}
