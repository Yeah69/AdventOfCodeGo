package main

import (
	"strconv"
	"strings"
)

func Day16ParseInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func Day16CommonTask(caveMap []string, initDir, initX, initY int) int {
	makeStep := func(dir, currX, currY int) Pair[int, int] {
		switch dir {
		case 0:
			return Pair[int, int]{First: currX + 1, Second: currY}
		case 1:
			return Pair[int, int]{First: currX, Second: currY + 1}
		case 2:
			return Pair[int, int]{First: currX - 1, Second: currY}
		case 3:
			return Pair[int, int]{First: currX, Second: currY - 1}
		}
		return Pair[int, int]{First: currX, Second: currY}
	}

	queue := make([]Pair[int, Pair[int, int]], 0)
	queue = append(queue, Pair[int, Pair[int, int]]{First: initDir, Second: Pair[int, int]{First: initX, Second: initY}})

	processed := make(map[Pair[int, Pair[int, int]]]bool)
	energized := make(map[Pair[int, int]]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, ok := processed[current]; ok {
			continue
		}

		direction := current.First
		x := current.Second.First
		y := current.Second.Second

		if x < 0 || x >= len(caveMap[0]) || y < 0 || y >= len(caveMap) {
			continue
		}

		energized[Pair[int, int]{First: x, Second: y}] = true
		processed[Pair[int, Pair[int, int]]{First: direction, Second: Pair[int, int]{First: x, Second: y}}] = true

		switch caveMap[y][x] {
		case '\\':
			switch direction {
			case 0:
				direction = 1
			case 1:
				direction = 0
			case 2:
				direction = 3
			case 3:
				direction = 2
			}
			queue = append(queue, Pair[int, Pair[int, int]]{First: direction, Second: makeStep(direction, x, y)})
		case '/':
			switch direction {
			case 0:
				direction = 3
			case 1:
				direction = 2
			case 2:
				direction = 1
			case 3:
				direction = 0
			}
			queue = append(queue, Pair[int, Pair[int, int]]{First: direction, Second: makeStep(direction, x, y)})
		case '|':
			if direction == 0 || direction == 2 {
				queue = append(queue, Pair[int, Pair[int, int]]{First: 1, Second: makeStep(1, x, y)})
				queue = append(queue, Pair[int, Pair[int, int]]{First: 3, Second: makeStep(3, x, y)})
			} else {
				queue = append(queue, Pair[int, Pair[int, int]]{First: direction, Second: makeStep(direction, x, y)})
			}
		case '-':
			if direction == 1 || direction == 3 {
				queue = append(queue, Pair[int, Pair[int, int]]{First: 0, Second: makeStep(0, x, y)})
				queue = append(queue, Pair[int, Pair[int, int]]{First: 2, Second: makeStep(2, x, y)})
			} else {
				queue = append(queue, Pair[int, Pair[int, int]]{First: direction, Second: makeStep(direction, x, y)})
			}
		case '.':
			queue = append(queue, Pair[int, Pair[int, int]]{First: direction, Second: makeStep(direction, x, y)})
		}
	}

	return len(energized)
}

func Day16Task0(input string) string {
	return strconv.Itoa(Day16CommonTask(Day16ParseInput(input), 0, 0, 0))
}

func Day16Task1(input string) string {
	caveMap := Day16ParseInput(input)

	maximum := -1

	for y := range caveMap {
		left := Day16CommonTask(caveMap, 0, 0, y)
		if left > maximum {
			maximum = left
		}
		right := Day16CommonTask(caveMap, 2, len(caveMap[0])-1, y)
		if right > maximum {
			maximum = right
		}
	}

	for x := range caveMap[0] {
		up := Day16CommonTask(caveMap, 1, x, 0)
		if up > maximum {
			maximum = up
		}
		down := Day16CommonTask(caveMap, 3, x, len(caveMap)-1)
		if down > maximum {
			maximum = down
		}
	}

	return strconv.Itoa(maximum)
}
