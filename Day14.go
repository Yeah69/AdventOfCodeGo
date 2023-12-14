package main

import (
	"slices"
	"strconv"
	"strings"
)

func Day14ParseInput(input string) [][]rune {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")
	stoneMap := make([][]rune, len(lines))
	for i := range lines {
		stoneMap[i] = []rune(lines[i])
	}
	return stoneMap
}

func Day14CalculateNorthLoad(stoneMap [][]rune) int {
	sum := 0
	for y := range stoneMap {
		load := len(stoneMap) - y
		for x := range stoneMap[y] {
			if stoneMap[y][x] == 'O' {
				sum = sum + load
			}
		}
	}
	return sum
}

func Day14ShiftNorth(stoneMap [][]rune) {
	for y := range stoneMap {
		for x := range stoneMap[y] {
			if stoneMap[y][x] == 'O' {
				iY := y
				for iY-1 >= 0 && stoneMap[iY-1][x] == '.' {
					iY--
				}
				if iY != y {
					stoneMap[y][x] = '.'
					stoneMap[iY][x] = 'O'
				}
			}
		}
	}
}

func Day14ShiftWest(stoneMap [][]rune) {
	for x := range stoneMap[0] {
		for y := range stoneMap {
			if stoneMap[y][x] == 'O' {
				iX := x
				for iX-1 >= 0 && stoneMap[y][iX-1] == '.' {
					iX--
				}
				if iX != x {
					stoneMap[y][x] = '.'
					stoneMap[y][iX] = 'O'
				}
			}
		}
	}
}

func Day14ShiftSouth(stoneMap [][]rune) {
	for y := len(stoneMap) - 1; y >= 0; y-- {
		for x := range stoneMap[y] {
			if stoneMap[y][x] == 'O' {
				iY := y
				for iY+1 < len(stoneMap) && stoneMap[iY+1][x] == '.' {
					iY++
				}
				if iY != y {
					stoneMap[y][x] = '.'
					stoneMap[iY][x] = 'O'
				}
			}
		}
	}
}

func Day14ShiftEast(stoneMap [][]rune) {
	for x := len(stoneMap[0]) - 1; x >= 0; x-- {
		for y := range stoneMap {
			if stoneMap[y][x] == 'O' {
				iX := x
				for iX+1 < len(stoneMap[0]) && stoneMap[y][iX+1] == '.' {
					iX++
				}
				if iX != x {
					stoneMap[y][x] = '.'
					stoneMap[y][iX] = 'O'
				}
			}
		}
	}
}

func Day14ShiftCycle(stoneMap [][]rune) int {
	Day14ShiftNorth(stoneMap)
	Day14ShiftWest(stoneMap)
	Day14ShiftSouth(stoneMap)
	Day14ShiftEast(stoneMap)
	return Day14CalculateNorthLoad(stoneMap)
}

func Day14DetectCycle(loads []int) int {
	if len(loads) < 100 {
		return -1
	}
	for i := 5; i <= len(loads)/2; i++ {
		first := loads[len(loads)-2*i : len(loads)-i]
		second := loads[len(loads)-i:]
		if slices.Equal(first, second) {
			return i
		}
	}
	return -1
}

func Day14Task0(input string) string {
	stoneMap := Day14ParseInput(input)
	Day14ShiftNorth(stoneMap)
	return strconv.Itoa(Day14CalculateNorthLoad(stoneMap))
}

func Day14Task1(input string) string {
	stoneMap := Day14ParseInput(input)
	loads := make([]int, 0)
	for i := 0; i < 1000000000; i++ {
		currentLoad := Day14ShiftCycle(stoneMap)
		loads = append(loads, currentLoad)
		if cycle := Day14DetectCycle(loads); cycle != -1 {
			initialCount := len(loads) - 2*cycle
			stepsBack := initialCount % cycle
			if stepsBack > 0 {
				return strconv.Itoa(loads[len(loads)-stepsBack])
			} else {
				return strconv.Itoa(loads[len(loads)-cycle])
			}
		}
	}
	return strconv.Itoa(Day14CalculateNorthLoad(stoneMap))
}
