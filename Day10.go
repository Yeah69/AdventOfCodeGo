package main

import (
	"strconv"
	"strings"
)

func Day10ParseInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func Day10IsNorth(pipeMap []string, x, y int) bool {
	char := pipeMap[y-1][x]
	return char == '|' || char == 'F' || char == '7'
}

func Day10IsSouth(pipeMap []string, x, y int) bool {
	char := pipeMap[y+1][x]
	return char == '|' || char == 'L' || char == 'J'
}

func Day10IsWest(pipeMap []string, x, y int) bool {
	char := pipeMap[y][x+1]
	return char == '-' || char == '7' || char == 'J'
}

func Day10IsEast(pipeMap []string, x, y int) bool {
	char := pipeMap[y][x-1]
	return char == '-' || char == 'L' || char == 'F'
}

func Day10GetPipeRoute(pipeMap []string, start Pair[int, int]) map[Pair[int, int]]bool {
	pipeRoute := make(map[Pair[int, int]]bool)

	currentOrigin := start
	current := Pair[int, int]{First: -1, Second: -1}

	if Day10IsNorth(pipeMap, start.First, start.Second) {
		current = Pair[int, int]{First: start.First, Second: start.Second - 1}
	} else if Day10IsSouth(pipeMap, start.First, start.Second) {
		current = Pair[int, int]{First: start.First, Second: start.Second + 1}
	} else if Day10IsWest(pipeMap, start.First, start.Second) {
		current = Pair[int, int]{First: start.First + 1, Second: start.Second}
	} else if Day10IsEast(pipeMap, start.First, start.Second) {
		current = Pair[int, int]{First: start.First - 1, Second: start.Second}
	}

	pipeRoute[current] = true

	for current != start {
		nextOrigin := current
		nextItem := func(first, second int) Pair[int, int] {
			ret := Pair[int, int]{First: first, Second: second}
			pipeRoute[ret] = true
			return ret
		}
		switch pipeMap[current.Second][current.First] {
		case '|':
			if currentOrigin.Second == current.Second-1 {
				current = nextItem(current.First, current.Second+1)
				break
			}
			current = nextItem(current.First, current.Second-1)
			break
		case '-':
			if currentOrigin.First == current.First-1 {
				current = nextItem(current.First+1, current.Second)
				break
			}
			current = nextItem(current.First-1, current.Second)
			break
		case 'L':
			if currentOrigin.Second == current.Second-1 {
				current = nextItem(current.First+1, current.Second)
				break
			}
			current = nextItem(current.First, current.Second-1)
			break
		case 'J':
			if currentOrigin.Second == current.Second-1 {
				current = nextItem(current.First-1, current.Second)
				break
			}
			current = nextItem(current.First, current.Second-1)
			break
		case '7':
			if currentOrigin.Second == current.Second+1 {
				current = nextItem(current.First-1, current.Second)
				break
			}
			current = nextItem(current.First, current.Second+1)
			break
		case 'F':
			if currentOrigin.Second == current.Second+1 {
				current = nextItem(current.First+1, current.Second)
				break
			}
			current = nextItem(current.First, current.Second+1)
		}
		currentOrigin = nextOrigin
	}

	return pipeRoute
}

func Day10Task0(input string) string {
	pipeMap := Day10ParseInput(input)

	start := func() Pair[int, int] {
		for y := range pipeMap {
			x := strings.Index(pipeMap[y], "S")
			if x != -1 {
				return Pair[int, int]{First: x, Second: y}
			}
		}
		return Pair[int, int]{First: -1, Second: -1}
	}()

	pipeRoute := Day10GetPipeRoute(pipeMap, start)

	return strconv.Itoa(len(pipeRoute) / 2)
}

func Day10Task1(input string) string {
	pipeMap := Day10ParseInput(input)

	start := func() Pair[int, int] {
		for y := range pipeMap {
			x := strings.Index(pipeMap[y], "S")
			if x != -1 {
				return Pair[int, int]{First: x, Second: y}
			}
		}
		return Pair[int, int]{First: -1, Second: -1}
	}()

	pipeRoute := Day10GetPipeRoute(pipeMap, start)

	isNorth := Day10IsNorth(pipeMap, start.First, start.Second)
	isSouth := Day10IsSouth(pipeMap, start.First, start.Second)
	isWest := Day10IsWest(pipeMap, start.First, start.Second)
	isEast := Day10IsEast(pipeMap, start.First, start.Second)

	if isNorth && isSouth {
		pipeMap[start.Second] = strings.ReplaceAll(pipeMap[start.Second], "S", "|")
	} else if isNorth && isWest {
		pipeMap[start.Second] = strings.ReplaceAll(pipeMap[start.Second], "S", "J")
	} else if isNorth && isEast {
		pipeMap[start.Second] = strings.ReplaceAll(pipeMap[start.Second], "S", "L")
	} else if isSouth && isWest {
		pipeMap[start.Second] = strings.ReplaceAll(pipeMap[start.Second], "S", "7")
	} else if isSouth && isEast {
		pipeMap[start.Second] = strings.ReplaceAll(pipeMap[start.Second], "S", "F")
	} else if isWest && isEast {
		pipeMap[start.Second] = strings.ReplaceAll(pipeMap[start.Second], "S", "-")
	}

	insideCounter := 0

	for y := range pipeMap {
		row := pipeMap[y]
		insideLayer := false
		lastEntryCorner := '.'
		for x := range row {
			if _, ok := pipeRoute[Pair[int, int]{First: x, Second: y}]; ok {
				switch pipeMap[y][x] {
				case '|':
					insideLayer = !insideLayer
				case 'L':
					insideLayer = !insideLayer
					lastEntryCorner = 'L'
				case 'F':
					insideLayer = !insideLayer
					lastEntryCorner = 'F'
				case 'J':
					if lastEntryCorner == 'L' {
						insideLayer = !insideLayer
					}
					lastEntryCorner = '.'
				case '7':
					if lastEntryCorner == 'F' {
						insideLayer = !insideLayer
					}
					lastEntryCorner = '.'
				case '-':
				}
			} else {
				if insideLayer {
					insideCounter++
				}
			}
		}
	}

	return strconv.Itoa(insideCounter)
}
