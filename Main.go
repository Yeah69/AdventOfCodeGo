package main

import (
	"fmt"
	"log"
	"os"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

var taskFunctionsMap = map[string]Pair[task, task]{
	"01": {Day01Task0, Day01Task1},
	"02": {Day02Task0, Day02Task1},
	"03": {Day03Task0, Day03Task1},
	"04": {Day04Task0, Day04Task1},
	"05": {Day05Task0, Day05Task1},
	"06": {Day06Task0, Day06Task1},
	"07": {Day07Task0, Day07Task1},
	"08": {Day08Task0, Day08Task1},
	"09": {Day09Task0, Day09Task1},
	"10": {Day10Task0, Day10Task1},
	"11": {Day11Task0, Day11Task1},
	"12": {Day12Task0, Day12Task1},
	"13": {Day13Task0, Day13Task1},
	"14": {Day14Task0, Day14Task1},
	"15": {Day15Task0, Day15Task1},
	"16": {Day16Task0, Day16Task1},
	"17": {Day17Task0, Day17Task1},
	"18": {Day18Task0, Day18Task1},
	"19": {Day19Task0, Day19Task1},
	"20": {Day20Task0, Day20Task1},
	"21": {Day21Task0, Day21Task1},
	"22": {Day22Task0, Day22Task1},
	"23": {Day23Task0, Day23Task1},
	"24": {Day24Task0, Day24Task1},
	"25": {Day25Task0, Day25Task1},
}

func main() {
	day := "12"

	fileName := fmt.Sprintf("./resources/Day%s.txt", day)

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	p := taskFunctionsMap[day]

	executeWholeDay(day, input, p.First, p.Second)
}
