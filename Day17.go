package main

import "strings"

func Day17ParseInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

type Day17State struct {
	X              int
	Y              int
	HorizontalNext bool
}

func Day17Task0(input string) string {
	cityMap := Day17ParseInput(input)

	queue := make([]Day17State, 0)

	firstHorizontalNext := Day17State{X: 0, Y: 0, HorizontalNext: true}
	firstVerticalNext := Day17State{X: 0, Y: 0, HorizontalNext: false}
	queue = append(queue, firstHorizontalNext)
	queue = append(queue, firstVerticalNext)

	cache := make(map[Day17State]int)
	cache[firstHorizontalNext] = 0
	cache[firstVerticalNext] = 0

	for len(queue) > 0 {

	}
	_ = cityMap

	return NoSolutionFound
}

func Day17Task1(input string) string {
	return NoSolutionFound
}
