package main

import (
	"slices"
	"strconv"
	"strings"
)

func Day25ParseInput(input string) map[string][]string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	adjacents := make(map[string][]string)

	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, ": ")
		mainNodeName := parts[0]
		if _, ok := adjacents[mainNodeName]; !ok {
			adjacents[mainNodeName] = make([]string, 0)
		}
		adjacentNodeNames := strings.Split(parts[1], " ")

		for i2 := range adjacentNodeNames {
			adjacentNodeName := adjacentNodeNames[i2]
			if _, ok := adjacents[mainNodeName]; !ok {
				adjacents[mainNodeName] = make([]string, 0)
			}
			mainList, _ := adjacents[mainNodeName]
			mainList = append(mainList, adjacentNodeName)
			adjacents[mainNodeName] = mainList
			adjacentList := adjacents[adjacentNodeName]
			adjacentList = append(adjacentList, mainNodeName)
			adjacents[adjacentNodeName] = adjacentList
		}
	}
	return adjacents
}

func Day25GetMaxNode(nodes map[string][]string) string {
	maxNeighbors := -1
	maxNode := ""

	for s := range nodes {
		neighbors := nodes[s]
		if len(neighbors) > maxNeighbors {
			maxNeighbors = len(neighbors)
			maxNode = s
		}
	}

	return maxNode
}

func Day25GetExternalVerticesCount(nodes map[string][]string, group []string) int {
	count := 0
	for i := range group {
		member := group[i]
		vertices := nodes[member]

		for i2 := range vertices {
			targetNode := vertices[i2]
			if !slices.Contains(group, targetNode) {
				count += 1
			}
		}
	}

	return count
}

func Day25Task0(input string) string {
	nodes := Day25ParseInput(input)
	initialNode := Day25GetMaxNode(nodes)
	currentNeighbors := make([]string, 0)
	for i := range nodes[initialNode] {
		currentNeighbors = append(currentNeighbors, nodes[initialNode][i])
	}
	currentGroup := make([]string, 0)
	currentGroup = append(currentGroup, initialNode)

	for Day25GetExternalVerticesCount(nodes, currentGroup) > 3 {
		minimum := 1000000
		pickedNeighbor := ""
		pickedNeighborIndex := -1
		for i := range currentNeighbors {
			neighbor := currentNeighbors[i]
			neighborsNeighbors := nodes[neighbor]
			countNewNeighbors := 0
			for i2 := range neighborsNeighbors {
				neighborsNeighbor := neighborsNeighbors[i2]
				if !slices.Contains(currentGroup, neighborsNeighbor) {
					countNewNeighbors += 1
				}
			}
			temp := 2*countNewNeighbors - len(neighborsNeighbors)
			if temp < minimum {
				minimum = temp
				pickedNeighbor = neighbor
				pickedNeighborIndex = i
			}
		}

		currentNeighbors = append(currentNeighbors[:pickedNeighborIndex], currentNeighbors[pickedNeighborIndex+1:]...)

		currentGroup = append(currentGroup, pickedNeighbor)
		neighborsNeighbors := nodes[pickedNeighbor]
		for i2 := range neighborsNeighbors {
			neighborsNeighbor := neighborsNeighbors[i2]
			if !slices.Contains(currentNeighbors, neighborsNeighbor) && !slices.Contains(currentGroup, neighborsNeighbor) {
				currentNeighbors = append(currentNeighbors, neighborsNeighbor)
			}
		}
	}

	result := len(currentGroup) * (len(nodes) - len(currentGroup))

	return strconv.Itoa(result)
}

func Day25Task1(_ string) string {
	return NothingToDoHere
}
