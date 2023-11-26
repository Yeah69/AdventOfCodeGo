package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day11CommonTask(input string, voidSummand int) string {
	galaxyMap := func() []string {
		input = strings.ReplaceAll(input, "\r\n", "\n")
		return strings.Split(input, "\n")
	}()
	emptyColumns := func() []int {
		columns := make([]int, 0)
		for x := range galaxyMap[0] {
			isEmpty := true
			for y := range galaxyMap {
				if galaxyMap[y][x] == '#' {
					isEmpty = false
					break
				}
			}
			if isEmpty {
				columns = append(columns, x)
			}
		}
		return columns
	}()
	emptyRows := func() []int {
		rows := make([]int, 0)
		for y := range galaxyMap {
			if strings.Index(galaxyMap[y], "#") == -1 {
				rows = append(rows, y)
			}
		}
		return rows
	}()
	positions := func() []Pair[int, int] {
		positions := make([]Pair[int, int], 0)
		for y := range galaxyMap {
			for x := range galaxyMap[y] {
				if galaxyMap[y][x] == '#' {
					positions = append(positions, Pair[int, int]{First: x, Second: y})
				}
			}
		}
		return positions
	}()

	sum := 0

	for i := range positions {
		start := positions[i]
		for i2 := i + 1; i2 < len(positions); i2++ {
			end := positions[i2]
			sumPart := int(math.Abs(float64(start.First-end.First))) + int(math.Abs(float64(start.Second-end.Second)))
			addVoidSummand := func(empties []int, getter func(pair Pair[int, int]) int) {
				minI := int(math.Min(float64(getter(start)), float64(getter(end)))) + 1
				maxI := int(math.Max(float64(getter(start)), float64(getter(end)))) - 1
				for i3 := minI; i3 <= maxI; i3++ {
					if slices.Contains(empties, i3) {
						sumPart = sumPart + voidSummand
					}
				}
			}
			addVoidSummand(emptyRows, func(p Pair[int, int]) int { return p.Second })
			addVoidSummand(emptyColumns, func(p Pair[int, int]) int { return p.First })
			sum = sum + sumPart
		}
	}

	return strconv.Itoa(sum)
}

func Day11Task0(input string) string { return Day11CommonTask(input, 1) }

func Day11Task1(input string) string { return Day11CommonTask(input, 999999) }
