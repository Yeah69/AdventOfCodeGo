package main

import (
	"fmt"
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
	return Day21Data{start, rocksMap, len(lines[0]), len(lines)}
}

func Day21Task0(input string) string {
	data := Day21ParseInput(input)
	start := data.Start
	rocksMap := data.RocksMap
	width := data.Width
	height := data.Height

	oddQueue := make([]Pair[int, int], 0)
	evenQueue := make([]Pair[int, int], 0)

	oddDoneMap := make(map[Pair[int, int]]bool)
	evenDoneMap := make(map[Pair[int, int]]bool)

	evenQueue = append(evenQueue, start)

	for i := 0; i < 65; i++ {
		makeStep := func(currentQueue, otherQueue *[]Pair[int, int], currentDoneMap, otherDoneMap *map[Pair[int, int]]bool) {
			for len(*currentQueue) > 0 {
				current := (*currentQueue)[0]
				*currentQueue = (*currentQueue)[1:]
				if _, doneOk := (*currentDoneMap)[current]; doneOk {
					continue
				}
				(*currentDoneMap)[current] = true
				testPoints := [4]Pair[int, int]{
					{current.First - 1, current.Second},
					{current.First + 1, current.Second},
					{current.First, current.Second - 1},
					{current.First, current.Second + 1},
				}
				for i2 := range testPoints {
					testPoint := testPoints[i2]
					if testPoint.First >= 0 && testPoint.Second >= 0 && testPoint.First < width && testPoint.Second < height {
						if _, ok := rocksMap[testPoint]; !ok {
							if _, doneOk := (*otherDoneMap)[testPoint]; !doneOk {
								*otherQueue = append(*otherQueue, testPoint)
							}
						}
					}
				}
			}
		}

		if i%2 == 0 {
			makeStep(&evenQueue, &oddQueue, &evenDoneMap, &oddDoneMap)
		} else {
			makeStep(&oddQueue, &evenQueue, &oddDoneMap, &evenDoneMap)
		}
		fmt.Println(i)
	}

	return strconv.Itoa(len(evenDoneMap))
}

func Day21Task1(input string) string {
	return NoSolutionFound
}
