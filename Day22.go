package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day22Coords struct {
	X int
	Y int
	Z int
}

type Day22Brick struct {
	Parts []Day22Coords
}

func Day22ParseInput(input string) []Day22Brick {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	bricks := make([]Day22Brick, len(lines))

	for i := range lines {
		line := lines[i]
		coordTexts := strings.Split(line, "~")
		getCoord := func(coordText string) Day22Coords {
			parts := strings.Split(coordText, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			return Day22Coords{x, y, z}
		}
		bricks[i] = Day22Brick{Day22GetCoords(getCoord(coordTexts[0]), getCoord(coordTexts[1]))}
	}

	return bricks
}

func Day22LowestZ(brick Day22Brick) int {
	minZ := math.MaxInt
	for i := range brick.Parts {
		if brick.Parts[i].Z < minZ {
			minZ = brick.Parts[i].Z
		}
	}
	return minZ
}

func Day22GetCoords(start, end Day22Coords) []Day22Coords {
	coords := make([]Day22Coords, 0)
	for x := int(math.Min(float64(start.X), float64(end.X))); x <= int(math.Max(float64(start.X), float64(end.X))); x++ {
		for y := int(math.Min(float64(start.Y), float64(end.Y))); y <= int(math.Max(float64(start.Y), float64(end.Y))); y++ {
			for z := int(math.Min(float64(start.Z), float64(end.Z))); z <= int(math.Max(float64(start.Z), float64(end.Z))); z++ {
				coords = append(coords, Day22Coords{x, y, z})
			}
		}
	}
	return coords
}

func Day22Falling(bricks []Day22Brick) Pair[[]Day22Brick, int] {
	slices.SortFunc(bricks, func(a, b Day22Brick) int {
		if Day22LowestZ(a) < Day22LowestZ(b) {
			return -1
		} else if Day22LowestZ(a) > Day22LowestZ(b) {
			return 1
		} else {
			return 0
		}
	})

	moved := 0

	for currI := range bricks {
		currentBrick := bricks[currI]

		currentLowestZ := Day22LowestZ(currentBrick)
		diff := currentLowestZ - 1

		currentCoords := currentBrick.Parts
		for currCoordI := range currentCoords {
			currentCoord := currentCoords[currCoordI]

			previousBricks := bricks[:currI]
			for prevI := len(previousBricks) - 1; prevI >= 0; prevI-- {
				previousBrick := previousBricks[prevI]

				previousCoords := previousBrick.Parts
				for prevCoordI := range previousCoords {
					previousCoord := previousCoords[prevCoordI]

					if currentCoord.X == previousCoord.X && currentCoord.Y == previousCoord.Y {
						diffToPrevious := currentCoord.Z - previousCoord.Z - 1
						if diffToPrevious < diff {
							diff = diffToPrevious
						}
					}
				}
			}
		}
		if diff > 0 {
			moved++
		}
		c := make([]Day22Coords, len(currentCoords))
		for i3 := range currentCoords {
			c[i3] = Day22Coords{currentCoords[i3].X, currentCoords[i3].Y, currentCoords[i3].Z - diff}
		}
		bricks[currI] = Day22Brick{c}
	}

	slices.SortFunc(bricks, func(a, b Day22Brick) int {
		if Day22LowestZ(a) < Day22LowestZ(b) {
			return -1
		} else if Day22LowestZ(a) > Day22LowestZ(b) {
			return 1
		} else {
			return 0
		}
	})

	return Pair[[]Day22Brick, int]{bricks, moved}
}

func Day22Task0(input string) string {
	bricks := Day22ParseInput(input)

	pair := Day22Falling(bricks)

	disintegrate := 0
	bricks = pair.First

	for i := range bricks {
		newBricks := make([]Day22Brick, 0)
		for i2 := range bricks {
			if i2 != i {
				c := make([]Day22Coords, len(bricks[i2].Parts))
				for i3 := range bricks[i2].Parts {
					c[i3] = Day22Coords{bricks[i2].Parts[i3].X, bricks[i2].Parts[i3].Y, bricks[i2].Parts[i3].Z}
				}
				newBricks = append(newBricks, Day22Brick{c})
			}
		}
		newPair := Day22Falling(newBricks)
		if newPair.Second == 0 {
			disintegrate++
		}
	}

	return strconv.Itoa(disintegrate)

	/*bricks = pair.First

	supportsMap := make(map[int][]int)
	supportedByMap := make(map[int][]int)

	for currI := range bricks {
		currentBrick := bricks[currI]

		currentCoords := currentBrick.Parts
		for currCoordI := range currentCoords {
			currentCoord := currentCoords[currCoordI]

			for remI := currI + 1; remI < len(bricks); remI++ {
				remainingBrick := bricks[remI]

				remainingCoords := remainingBrick.Parts
				for remCoordI := range remainingCoords {
					remainingCoord := remainingCoords[remCoordI]

					if currentCoord.X == remainingCoord.X && currentCoord.Y == remainingCoord.Y && currentCoord.Z+1 == remainingCoord.Z {
						if list, ok := supportsMap[currI]; ok {
							supportsMap[currI] = append(list, remI)
						} else {
							newList := make([]int, 0)
							newList = append(newList, remI)
							supportsMap[currI] = newList
						}

						if list, ok := supportedByMap[remI]; ok {
							supportedByMap[remI] = append(list, currI)
						} else {
							newList := make([]int, 0)
							newList = append(newList, currI)
							supportedByMap[remI] = newList
						}

						break
					}
				}
			}
		}
	}

	disintegrate := 0

	for i := range bricks {
		if supported, ok := supportsMap[i]; ok {
			shouldDisintegrate := true
			for i2 := range supported {
				currI := supported[i2]
				if supportedBy, okSupported := supportedByMap[currI]; okSupported && len(supportedBy) <= 1 {
					shouldDisintegrate = false
					break
				}
			}
			if shouldDisintegrate {
				disintegrate++
			}
		} else {
			disintegrate++
		}
	}

	return strconv.Itoa(disintegrate)*/
}

func Day22Task1(input string) string {
	bricks := Day22ParseInput(input)

	pair := Day22Falling(bricks)

	movedPiecesSum := 0
	bricks = pair.First

	for i := range bricks {
		newBricks := make([]Day22Brick, 0)
		for i2 := range bricks {
			if i2 != i {
				c := make([]Day22Coords, len(bricks[i2].Parts))
				for i3 := range bricks[i2].Parts {
					c[i3] = Day22Coords{bricks[i2].Parts[i3].X, bricks[i2].Parts[i3].Y, bricks[i2].Parts[i3].Z}
				}
				newBricks = append(newBricks, Day22Brick{c})
			}
		}
		newPair := Day22Falling(newBricks)
		movedPiecesSum = movedPiecesSum + newPair.Second
	}

	return strconv.Itoa(movedPiecesSum)
}
