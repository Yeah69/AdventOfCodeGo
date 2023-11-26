package main

import (
	"slices"
	"strconv"
	"strings"
)

type Day07Game struct {
	Hand [5]int
	Bid  int
	Type int
}

func Day07ParseInput(input string, treatJokerAsSuch bool) []Day07Game {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	games := make([]Day07Game, len(lines))
	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, " ")

		handText := parts[0]
		hand := [5]int{0, 0, 0, 0, 0}

		if len(handText) == 5 {
			for i2 := range handText {
				char := handText[i2]
				number := 0
				switch char {
				case 'A':
					number = 14
				case 'K':
					number = 13
				case 'Q':
					number = 12
				case 'J':
					if treatJokerAsSuch {
						number = 1
					}
					if !treatJokerAsSuch {
						number = 11
					}
				case 'T':
					number = 10
				default:
					number = int(char - '0')
				}
				hand[i2] = number
			}
		}

		bid, _ := strconv.Atoi(parts[1])

		typeOfHand := Day07GetTypeOuter(hand)

		games[i] = Day07Game{Hand: hand, Bid: bid, Type: typeOfHand}
	}
	return games
}

func Day07GetTypeOuter(hand [5]int) int {
	index := -1
	for i := range hand {
		if hand[i] == 1 {
			index = i
			break
		}
	}

	if index == -1 {
		return Day07GetType(hand)
	}

	numbers := make([]int, 0)
	for i := range hand {
		if hand[i] != 1 && !slices.Contains(numbers, hand[i]) {
			numbers = append(numbers, hand[i])
		}
	}

	if len(numbers) == 0 {
		return 7 // no other numbers means the hand is all joker => treat like "Five of a Kind"
	}

	maximum := -1
	for i := range numbers {
		hand[index] = numbers[i]
		current := Day07GetTypeOuter(hand)
		if current > maximum {
			maximum = current
		}
	}
	return maximum
}

func Day07GetType(hand [5]int) int {
	// Count numbers of same kind and cache in map
	counts := make(map[int]int)
	for i := range hand {
		number := hand[i]
		count, hasKey := counts[number]
		if hasKey {
			counts[number] = count + 1
		}
		if !hasKey {
			counts[number] = 1
		}
	}

	// Decide which type of hand by cached counts
	switch len(counts) {
	case 1:
		return 7 // Five of a Kind
	case 2:
		for _, value := range counts {
			if value == 4 || value == 1 {
				return 6 // Four of a kind
			}
			return 5 // Full House
		}
	case 3:
		for _, value := range counts {
			if value == 3 {
				return 4 // Three of a kind
			}
			if value == 2 {
				return 3 // Two Pairs
			}
		}
		return 0 // Impossible
	case 4:
		return 2 // One Pair
	case 5:
		return 1 // High Card
	default:
		return 0 // Impossible
	}
	return 0
}

func Day07CommonTask(input string, treatJokerAsSuch bool) string {
	games := Day07ParseInput(input, treatJokerAsSuch)

	slices.SortFunc(games, func(a, b Day07Game) int {
		if a.Type < b.Type {
			return -1
		}
		if a.Type > b.Type {
			return 1
		}
		for i := range a.Hand {
			if a.Hand[i] < b.Hand[i] {
				return -1
			}
			if a.Hand[i] > b.Hand[i] {
				return 1
			}
		}
		return 0
	})

	sum := int64(0)

	for i := range games {
		game := games[i]
		sum = sum + int64(game.Bid*(i+1))
	}

	return strconv.FormatInt(sum, 10)
}

func Day07Task0(input string) string {
	return Day07CommonTask(input, false)
}

func Day07Task1(input string) string {
	return Day07CommonTask(input, true)
}
