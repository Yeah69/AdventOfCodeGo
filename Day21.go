package main

import (
	"strconv"
	"strings"
)

type Day21Data struct {
	Start    Pair[int, int]
	RocksMap map[Pair[int, int]]bool
	Width    int
	Height   int
}

func Day21ParseInput(input string) Day21Data {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	start := Pair[int, int]{0, 0}
	rocksMap := make(map[Pair[int, int]]bool)

	for y := range lines {
		line := lines[y]
		for x := range line {
			if line[x] == '#' {
				rocksMap[Pair[int, int]{x, y}] = true
			} else if line[x] == 'S' {
				start = Pair[int, int]{x, y}
			}
		}
	}

	height := len(lines)
	width := len(lines[0])

	for y := -2; y < height+2; y++ {
		if y >= 0 && y < height {
			continue
		}

		projectedY := y - height
		if y < 0 {
			projectedY = height + y
		}
		for x := -2; x < width+2; x++ {
			if x >= 0 && x < width {
				continue
			}

			projectedX := x - width
			if x < 0 {
				projectedX = width + x
			}

			projectedPosition := Pair[int, int]{projectedX, projectedY}
			if _, ok := rocksMap[projectedPosition]; ok {
				rocksMap[Pair[int, int]{x, y}] = true
			}
		}
	}
	return Day21Data{start, rocksMap, len(lines[0]), len(lines)}
}

func Day21MakeTwoSteps0(
	rocksMap *map[Pair[int, int]]bool,
	queue *[]Pair[int, int],
	doneMap *map[Pair[int, int]]bool,
	width, height int) {
	newQueue := make([]Pair[int, int], 0)
	for len(*queue) > 0 {
		current := (*queue)[0]
		*queue = (*queue)[1:]

		testPoints := make([]Pair[int, int], 0)
		_ = testPoints

		isTop := true
		if is, ok := (*rocksMap)[Pair[int, int]{current.First, current.Second - 1}]; ok && is {
			isTop = false
		}
		isRight := true
		if is, ok := (*rocksMap)[Pair[int, int]{current.First + 1, current.Second}]; ok && is {
			isRight = false
		}
		isBottom := true
		if is, ok := (*rocksMap)[Pair[int, int]{current.First, current.Second + 1}]; ok && is {
			isBottom = false
		}
		isLeft := true
		if is, ok := (*rocksMap)[Pair[int, int]{current.First - 1, current.Second}]; ok && is {
			isLeft = false
		}

		tp := Pair[int, int]{current.First, current.Second - 2}
		if _, ok := (*rocksMap)[tp]; !ok && isTop {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First + 1, current.Second - 1}
		if _, ok := (*rocksMap)[tp]; !ok && (isTop || isRight) {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First + 2, current.Second}
		if _, ok := (*rocksMap)[tp]; !ok && isRight {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First + 1, current.Second + 1}
		if _, ok := (*rocksMap)[tp]; !ok && (isRight || isBottom) {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First, current.Second + 2}
		if _, ok := (*rocksMap)[tp]; !ok && isBottom {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First - 1, current.Second + 1}
		if _, ok := (*rocksMap)[tp]; !ok && (isBottom || isLeft) {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First - 2, current.Second}
		if _, ok := (*rocksMap)[tp]; !ok && isLeft {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First - 1, current.Second - 1}
		if _, ok := (*rocksMap)[tp]; !ok && (isLeft || isTop) {
			testPoints = append(testPoints, tp)
		}

		for i2 := range testPoints {
			testPoint := testPoints[i2]
			if testPoint.First >= 0 && testPoint.Second >= 0 && testPoint.First < width && testPoint.Second < height {
				if _, ok := (*rocksMap)[testPoint]; !ok {
					if _, doneOk := (*doneMap)[testPoint]; !doneOk {
						(*doneMap)[testPoint] = true
						newQueue = append(newQueue, testPoint)
					}
				}
			}
		}
	}
	*queue = newQueue
}

func Day21MakeTwoSteps1(
	rocksMap *map[Pair[int, int]]bool,
	queue *[]Pair[int, int],
	doneMap *map[Pair[int, int]]bool,
	width, height int) {

	hitsRock := func(first, second int) bool {
		for first < 0 {
			first += width
		}
		first = first % width
		for second < 0 {
			second += height
		}
		second = second % height
		hit, ok := (*rocksMap)[Pair[int, int]{first, second}]
		return ok && hit
	}

	newQueue := make([]Pair[int, int], 0)
	for len(*queue) > 0 {
		current := (*queue)[0]
		*queue = (*queue)[1:]

		testPoints := make([]Pair[int, int], 0)
		_ = testPoints

		isTop := true
		if hitsRock(current.First, current.Second-1) {
			isTop = false
		}
		isRight := true
		if hitsRock(current.First+1, current.Second) {
			isRight = false
		}
		isBottom := true
		if hitsRock(current.First, current.Second+1) {
			isBottom = false
		}
		isLeft := true
		if hitsRock(current.First-1, current.Second) {
			isLeft = false
		}

		tp := Pair[int, int]{current.First, current.Second - 2}
		if !hitsRock(tp.First, tp.Second) && isTop {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First + 1, current.Second - 1}
		if !hitsRock(tp.First, tp.Second) && (isTop || isRight) {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First + 2, current.Second}
		if !hitsRock(tp.First, tp.Second) && isRight {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First + 1, current.Second + 1}
		if !hitsRock(tp.First, tp.Second) && (isRight || isBottom) {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First, current.Second + 2}
		if !hitsRock(tp.First, tp.Second) && isBottom {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First - 1, current.Second + 1}
		if !hitsRock(tp.First, tp.Second) && (isBottom || isLeft) {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First - 2, current.Second}
		if !hitsRock(tp.First, tp.Second) && isLeft {
			testPoints = append(testPoints, tp)
		}
		tp = Pair[int, int]{current.First - 1, current.Second - 1}
		if !hitsRock(tp.First, tp.Second) && (isLeft || isTop) {
			testPoints = append(testPoints, tp)
		}

		for i2 := range testPoints {
			testPoint := testPoints[i2]
			if !hitsRock(testPoint.First, testPoint.Second) {
				if _, doneOk := (*doneMap)[testPoint]; !doneOk {
					(*doneMap)[testPoint] = true
					newQueue = append(newQueue, testPoint)
				}
			}
		}
	}
	*queue = newQueue
}

func Day21FixedSteps0(start Pair[int, int], rocksMap *map[Pair[int, int]]bool, width, height, steps int) int {
	queue := make([]Pair[int, int], 0)

	doneMap := make(map[Pair[int, int]]bool)

	queue = append(queue, start)

	doneMap[start] = true

	for i := 0; i < steps; i = i + 2 {
		Day21MakeTwoSteps0(rocksMap, &queue, &doneMap, width, height)
	}

	return len(doneMap)
}

func Day21FixedSteps1(start Pair[int, int], rocksMap *map[Pair[int, int]]bool, width, height, steps int) int {
	queue := make([]Pair[int, int], 0)

	doneMap := make(map[Pair[int, int]]bool)

	firstStep := make([]Pair[int, int], 0)
	if is, ok := (*rocksMap)[Pair[int, int]{start.First - 1, start.Second}]; !ok || !is {
		firstStep = append(firstStep, Pair[int, int]{start.First - 1, start.Second})
	}
	if is, ok := (*rocksMap)[Pair[int, int]{start.First + 1, start.Second}]; !ok || !is {
		firstStep = append(firstStep, Pair[int, int]{start.First + 1, start.Second})
	}
	if is, ok := (*rocksMap)[Pair[int, int]{start.First, start.Second - 1}]; !ok || !is {
		firstStep = append(firstStep, Pair[int, int]{start.First, start.Second - 1})
	}
	if is, ok := (*rocksMap)[Pair[int, int]{start.First, start.Second + 1}]; !ok || !is {
		firstStep = append(firstStep, Pair[int, int]{start.First, start.Second + 1})
	}

	for i := range firstStep {
		current := firstStep[i]
		queue = append(queue, current)
		doneMap[current] = true
	}

	for i := 1; i < steps; i = i + 2 {
		Day21MakeTwoSteps1(rocksMap, &queue, &doneMap, width, height)
	}

	return len(doneMap)
}

func Day21Task0(input string) string {
	data := Day21ParseInput(input)
	start := data.Start
	rocksMap := data.RocksMap
	width := data.Width
	height := data.Height

	return strconv.Itoa(Day21FixedSteps0(start, &rocksMap, width, height, 64))
}

func Day21Task1(input string) string {
	data := Day21ParseInput(input)
	start := data.Start
	rocksMap := data.RocksMap
	width := data.Width
	height := data.Height

	t0 := float64(65)
	t1 := float64(327)
	t2 := float64(589)

	A0 := float64(Day21FixedSteps1(start, &rocksMap, width, height, int(t0)))
	A1 := float64(Day21FixedSteps1(start, &rocksMap, width, height, int(t1)))
	A2 := float64(Day21FixedSteps1(start, &rocksMap, width, height, int(t2)))

	a := (A2 - (A1-A0)*(t2-t0)/(t1-t0) - A0) / (t2*t2 - (t1*t1-t0*t0)*(t2-t0)/(t1-t0) - t0*t0)
	b := (A1 - A0 - a*(t1*t1-t0*t0)) / (t1 - t0)
	c := A0 - a*t0*t0 - b*t0

	steps := float64(26501365)
	A := a*steps*steps + b*steps + c

	return strconv.Itoa(int(A))
}
