package main

import (
	"strconv"
	"strings"
)

type Day19Condition struct {
	Component  rune
	BiggerThan bool
	Value      int
	Then       string
}

type Day19ConditionChain struct {
	Chain     []Day19Condition
	Otherwise string
}

type Day19Part struct {
	X int
	M int
	A int
	S int
}

type Day19PartRange struct {
	X Pair[int, int]
	M Pair[int, int]
	A Pair[int, int]
	S Pair[int, int]
}

func Day19ParseInput(input string) Pair[map[string]Day19ConditionChain, []Day19Part] {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	bigParts := strings.Split(input, "\n\n")

	conditionChainLines := strings.Split(bigParts[0], "\n")
	conditionChains := make(map[string]Day19ConditionChain)
	for i := range conditionChainLines {
		outerParts := strings.Split(conditionChainLines[i][:len(conditionChainLines[i])-1], "{")

		parts := strings.Split(outerParts[1], ",")
		chain := make([]Day19Condition, len(parts)-1)
		for j := 0; j < len(parts)-1; j++ {
			innerParts := strings.Split(parts[j][2:], ":")
			component := rune(parts[j][0])
			biggerThan := parts[j][1] == '>'
			value, _ := strconv.Atoi(innerParts[0])
			then := innerParts[1]
			chain[j] = Day19Condition{component, biggerThan, value, then}
		}

		conditionChains[outerParts[0]] = Day19ConditionChain{chain, parts[len(parts)-1]}
	}

	partsLines := strings.Split(bigParts[1], "\n")
	parts := make([]Day19Part, len(partsLines))
	for i := range partsLines {
		partsParts := strings.Split(partsLines[i][1:len(partsLines[i])-1], ",")
		x, _ := strconv.Atoi(strings.Split(partsParts[0], "=")[1])
		m, _ := strconv.Atoi(strings.Split(partsParts[1], "=")[1])
		a, _ := strconv.Atoi(strings.Split(partsParts[2], "=")[1])
		s, _ := strconv.Atoi(strings.Split(partsParts[3], "=")[1])
		parts[i] = Day19Part{x, m, a, s}
	}
	return Pair[map[string]Day19ConditionChain, []Day19Part]{conditionChains, parts}
}

func Day19Task0(input string) string {
	pair := Day19ParseInput(input)
	conditionChains := pair.First
	parts := pair.Second

	sum := 0

	for i := range parts {
		part := parts[i]
		currentKey := "in"
		for {
			chain, _ := conditionChains[currentKey]
			found := false
			for i2 := range chain.Chain {
				condition := chain.Chain[i2]
				value := -1
				switch condition.Component {
				case 'x':
					value = part.X
				case 'm':
					value = part.M
				case 'a':
					value = part.A
				case 's':
					value = part.S
				}
				if condition.BiggerThan && value > condition.Value || !condition.BiggerThan && value < condition.Value {
					found = true
					currentKey = condition.Then
					break
				}
			}
			if !found {
				currentKey = chain.Otherwise
			}
			if currentKey == "A" {
				sum = sum + part.X + part.M + part.A + part.S
				break
			}
			if currentKey == "R" {
				break
			}
		}
	}

	return strconv.Itoa(sum)
}

func Day19GetAllPartRanges(
	currentPartRange Day19PartRange,
	currentKey string,
	conditionChains *map[string]Day19ConditionChain) int {
	getSplitRange := func(cur Pair[int, int], conditionIsBigger bool, conditionValue int) Pair[Pair[int, int], Pair[int, int]] {
		if conditionIsBigger {
			if cur.First > conditionValue {
				return Pair[Pair[int, int], Pair[int, int]]{cur, Pair[int, int]{-1, -1}}
			}
			if cur.Second <= conditionValue {
				return Pair[Pair[int, int], Pair[int, int]]{Pair[int, int]{-1, -1}, cur}
			}
			return Pair[Pair[int, int], Pair[int, int]]{Pair[int, int]{conditionValue + 1, cur.Second}, Pair[int, int]{cur.First, conditionValue}}
		} else {
			if cur.Second < conditionValue {
				return Pair[Pair[int, int], Pair[int, int]]{cur, Pair[int, int]{-1, -1}}
			}
			if cur.First >= conditionValue {
				return Pair[Pair[int, int], Pair[int, int]]{Pair[int, int]{-1, -1}, cur}
			}
			return Pair[Pair[int, int], Pair[int, int]]{Pair[int, int]{cur.First, conditionValue - 1}, Pair[int, int]{conditionValue, cur.Second}}
		}
	}

	if currentKey == "A" {
		return (currentPartRange.X.Second - currentPartRange.X.First + 1) *
			(currentPartRange.M.Second - currentPartRange.M.First + 1) *
			(currentPartRange.A.Second - currentPartRange.A.First + 1) *
			(currentPartRange.S.Second - currentPartRange.S.First + 1)
	}

	if currentKey == "R" {
		return 0
	}

	conditionChain, _ := (*conditionChains)[currentKey]
	sum := 0

	for i := range conditionChain.Chain {
		chainItem := conditionChain.Chain[i]
		currentRange := currentPartRange.X
		switch chainItem.Component {
		case 'm':
			currentRange = currentPartRange.M
		case 'a':
			currentRange = currentPartRange.A
		case 's':
			currentRange = currentPartRange.S
		}
		pairs := getSplitRange(currentRange, chainItem.BiggerThan, chainItem.Value)
		if pairs.First.First != -1 && pairs.First.Second != -1 {
			newPartRange := Day19PartRange{currentPartRange.X, currentPartRange.M, currentPartRange.A, currentPartRange.S}
			switch chainItem.Component {
			case 'x':
				newPartRange.X = pairs.First
			case 'm':
				newPartRange.M = pairs.First
			case 'a':
				newPartRange.A = pairs.First
			case 's':
				newPartRange.S = pairs.First
			}
			sum = sum + Day19GetAllPartRanges(newPartRange, chainItem.Then, conditionChains)
		}
		if pairs.Second.First == -1 || pairs.Second.Second == -1 {
			return sum
		}
		switch chainItem.Component {
		case 'x':
			currentPartRange.X = pairs.Second
		case 'm':
			currentPartRange.M = pairs.Second
		case 'a':
			currentPartRange.A = pairs.Second
		case 's':
			currentPartRange.S = pairs.Second
		}
	}
	return sum + Day19GetAllPartRanges(currentPartRange, conditionChain.Otherwise, conditionChains)
}

func Day19Task1(input string) string {
	pair := Day19ParseInput(input)
	conditionChains := pair.First
	initialRange := Day19PartRange{
		Pair[int, int]{1, 4000},
		Pair[int, int]{1, 4000},
		Pair[int, int]{1, 4000},
		Pair[int, int]{1, 4000}}

	sum := Day19GetAllPartRanges(initialRange, "in", &conditionChains)

	return strconv.Itoa(sum)
}
