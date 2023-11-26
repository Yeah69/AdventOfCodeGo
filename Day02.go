package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day02Set struct {
	Red   int
	Green int
	Blue  int
}

type Day02Game struct {
	ID   int
	Sets []Day02Set
}

func Day02ParseInput(input string) []Day02Game {

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	games := make([]Day02Game, len(lines))
	for i := range lines {
		line := lines[i]
		gameSplit := strings.Split(line, ":")
		id, _ := strconv.ParseInt(strings.ReplaceAll(gameSplit[0], "Game ", ""), 10, 32)
		setsSplit := strings.Split(gameSplit[1], ";")
		sets := make([]Day02Set, len(setsSplit))
		for i2 := range setsSplit {
			setParts := strings.Split(setsSplit[i2], ",")
			red := 0
			green := 0
			blue := 0
			for i3 := range setParts {
				color := setParts[i3]
				if strings.Contains(color, "red") {
					numberAsText := strings.ReplaceAll(strings.ReplaceAll(color, "red", ""), " ", "")
					redTemp, _ := strconv.ParseInt(numberAsText, 10, 32)
					red = int(redTemp)
				}
				if strings.Contains(color, "green") {
					numberAsText := strings.ReplaceAll(strings.ReplaceAll(color, "green", ""), " ", "")
					greenTemp, _ := strconv.ParseInt(numberAsText, 10, 32)
					green = int(greenTemp)
				}
				if strings.Contains(color, "blue") {
					numberAsText := strings.ReplaceAll(strings.ReplaceAll(color, "blue", ""), " ", "")
					blueTemp, _ := strconv.ParseInt(numberAsText, 10, 32)
					blue = int(blueTemp)
				}
			}
			sets[i2] = Day02Set{Red: red, Green: green, Blue: blue}
		}
		games[i] = Day02Game{ID: int(id), Sets: sets}
	}
	return games
}

func Day02CommonTask(input string, getCount func(Day02Game) int) string {
	games := Day02ParseInput(input)

	sum := 0
	for i := range games {
		sum = sum + getCount(games[i])
	}

	return fmt.Sprintf("%d", sum)
}

func Day02Task0(input string) string {
	possibleSet := Day02Set{Red: 12, Green: 13, Blue: 14}

	return Day02CommonTask(input, func(game Day02Game) int {
		for i2 := range game.Sets {
			set := game.Sets[i2]
			if set.Red > possibleSet.Red || set.Green > possibleSet.Green || set.Blue > possibleSet.Blue {
				return 0
			}
		}
		return game.ID
	})
}

func Day02Task1(input string) string {
	return Day02CommonTask(input, func(game Day02Game) int {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for i2 := range game.Sets {
			set := game.Sets[i2]
			if set.Red > maxRed {
				maxRed = set.Red
			}
			if set.Green > maxGreen {
				maxGreen = set.Green
			}
			if set.Blue > maxBlue {
				maxBlue = set.Blue
			}
		}
		return maxRed * maxGreen * maxBlue
	})
}
