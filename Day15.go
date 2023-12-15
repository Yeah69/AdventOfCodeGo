package main

import (
	"slices"
	"strconv"
	"strings"
)

func Day15ParseInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "")
	input = strings.ReplaceAll(input, "\n", "")
	return strings.Split(input, ",")
}

func Day15Hash(in string) int {
	currentValue := 0
	for i2 := range in {
		currentValue = ((currentValue + int(in[i2])) * 17) % 256
	}
	return currentValue
}

func Day15Task0(input string) string {
	instructions := Day15ParseInput(input)
	sum := 0
	for i := range instructions {
		sum = sum + Day15Hash(instructions[i])
	}
	return strconv.Itoa(sum)
}

func Day15Task1(input string) string {
	instructions := Day15ParseInput(input)
	boxes := make([][]Pair[string, int], 256)
	for i := range boxes {
		boxes[i] = make([]Pair[string, int], 0)
	}
	for i := range instructions {
		instruction := instructions[i]
		data := Pair[string, int]{First: "", Second: -1}
		if strings.Contains(instruction, "=") {
			parts := strings.Split(instruction, "=")
			number, _ := strconv.Atoi(parts[1])
			data = Pair[string, int]{First: parts[0], Second: number}
		} else {
			data = Pair[string, int]{First: instruction[:len(instruction)-1], Second: -1}
		}
		boxIndex := Day15Hash(data.First)
		if data.Second == -1 {
			boxes[boxIndex] = slices.DeleteFunc(boxes[boxIndex], func(item Pair[string, int]) bool {
				return item.First == data.First
			})
		} else {
			index := slices.IndexFunc(boxes[boxIndex], func(item Pair[string, int]) bool {
				return item.First == data.First
			})
			if index == -1 {
				boxes[boxIndex] = append(boxes[boxIndex], data)
			} else {
				boxes[boxIndex][index] = data
			}
		}
	}
	sum := 0
	for i := range boxes {
		for i2 := range boxes[i] {
			sum = sum + (i+1)*(i2+1)*boxes[i][i2].Second
		}
	}
	return strconv.Itoa(sum)
}
