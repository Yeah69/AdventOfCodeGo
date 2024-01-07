package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day20Module struct {
	Label string
	Type  int // 0: FlipFlop, 1: Conjunction, 2: broadcaster, 3: button
	Next  []string
}

func Day20ParseInput(input string) map[string]Day20Module {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	moduleMap := make(map[string]Day20Module)

	button := Day20Module{"button", 3, append(make([]string, 0), "broadcaster")}
	moduleMap["button"] = button

	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, " -> ")

		label := parts[0]
		if label != "broadcaster" {
			label = label[1:]
		}

		moduleType := 2
		if parts[0][0] == '%' {
			moduleType = 0
		} else if parts[0][0] == '&' {
			moduleType = 1
		}

		module := Day20Module{label, moduleType, strings.Split(parts[1], ", ")}
		moduleMap[label] = module
	}

	return moduleMap
}

func Day20Task0(input string) string {
	modulesMap := Day20ParseInput(input)

	flipFlopOnStateMap := make(map[string]bool)
	conjunctionMemoryMap := make(map[string]map[string]bool)

	for label, module := range modulesMap {
		if module.Type == 0 {
			flipFlopOnStateMap[label] = false
		}
		for i := range module.Next {
			if nextModule, ok := modulesMap[module.Next[i]]; ok && nextModule.Type == 1 {
				origin := module.Label
				target := nextModule.Label
				if memoryMap, memoryOk := conjunctionMemoryMap[target]; memoryOk {
					memoryMap[origin] = false
				} else {
					memoryMap0 := make(map[string]bool)
					memoryMap0[origin] = false
					conjunctionMemoryMap[target] = memoryMap0
				}
			}
		}
	}

	lowPulseCount := 0
	highPulseCount := 0

	for i := 0; i < 1000; i++ {
		queue := make([]Pair[Day20Module, bool], 0)
		queue = append(queue, Pair[Day20Module, bool]{modulesMap["button"], false})
		for len(queue) > 0 {
			currentPair := queue[0]
			queue = queue[1:]
			currentModule := currentPair.First
			currentPulse := currentPair.Second

			sendPulses := func(pulse bool) {
				origin := currentModule.Label
				for i2 := range currentModule.Next {
					target := currentModule.Next[i2]
					if pulse {
						highPulseCount++
					} else {
						lowPulseCount++
					}
					if module, ok := modulesMap[target]; ok {
						queue = append(queue, Pair[Day20Module, bool]{module, pulse})
						if module.Type == 1 {
							if memoryMap, memoryOk := conjunctionMemoryMap[target]; memoryOk {
								memoryMap[origin] = pulse
							} else {
								memoryMap0 := make(map[string]bool)
								memoryMap0[origin] = pulse
								conjunctionMemoryMap[target] = memoryMap0
							}
						}
					}
				}
			}

			if currentModule.Type == 0 {
				if !currentPulse {
					on, _ := flipFlopOnStateMap[currentModule.Label]
					sendingPulse := true
					if on {
						sendingPulse = false
					}
					flipFlopOnStateMap[currentModule.Label] = !on
					sendPulses(sendingPulse)
				}
			} else if currentModule.Type == 1 {
				sendingPulse := true
				if memoryMap, ok := conjunctionMemoryMap[currentModule.Label]; ok {
					for _, memorizedPulse := range memoryMap {
						if !memorizedPulse {
							sendingPulse = false
							break
						}
					}
				}
				sendPulses(!sendingPulse)
			} else if currentModule.Type == 2 {
				sendPulses(currentPulse)
			} else if currentModule.Type == 3 {
				sendPulses(false)
			}
		}
	}

	return strconv.Itoa(lowPulseCount * highPulseCount)
}

func Day20Task1(input string) string {
	modulesMap := Day20ParseInput(input)

	flipFlopOnStateMap := make(map[string]bool)
	conjunctionMemoryMap := make(map[string]map[string]bool)

	for label, module := range modulesMap {
		if module.Type == 0 {
			flipFlopOnStateMap[label] = false
		}
		for i := range module.Next {
			if nextModule, ok := modulesMap[module.Next[i]]; ok && nextModule.Type == 1 {
				origin := module.Label
				target := nextModule.Label
				if memoryMap, memoryOk := conjunctionMemoryMap[target]; memoryOk {
					memoryMap[origin] = false
				} else {
					memoryMap0 := make(map[string]bool)
					memoryMap0[origin] = false
					conjunctionMemoryMap[target] = memoryMap0
				}
			}
		}
	}

	lowPulseCount := 0
	highPulseCount := 0

	for i := 0; i < 100000; i++ {
		fmt.Printf("click := %d", i+1)
		fmt.Println()
		queue := make([]Pair[Day20Module, bool], 0)
		queue = append(queue, Pair[Day20Module, bool]{modulesMap["button"], false})
		for len(queue) > 0 {
			currentPair := queue[0]
			queue = queue[1:]
			currentModule := currentPair.First
			currentPulse := currentPair.Second

			sendPulses := func(pulse bool) {
				origin := currentModule.Label
				for i2 := range currentModule.Next {
					target := currentModule.Next[i2]
					if pulse {
						highPulseCount++
					} else {
						lowPulseCount++
					}
					if module, ok := modulesMap[target]; ok {
						queue = append(queue, Pair[Day20Module, bool]{module, pulse})
						if module.Type == 1 {
							if memoryMap, memoryOk := conjunctionMemoryMap[target]; memoryOk {
								memoryMap[origin] = pulse
							} else {
								memoryMap0 := make(map[string]bool)
								memoryMap0[origin] = pulse
								conjunctionMemoryMap[target] = memoryMap0
							}
						}
					}
				}
			}

			if currentModule.Type == 0 {
				if !currentPulse {
					on, _ := flipFlopOnStateMap[currentModule.Label]
					sendingPulse := true
					if on {
						sendingPulse = false
					}
					flipFlopOnStateMap[currentModule.Label] = !on
					sendPulses(sendingPulse)
				}
			} else if currentModule.Type == 1 {
				if currentModule.Label == "cs" {
					fmt.Println(conjunctionMemoryMap["cs"])
				}
				sendingPulse := true
				if memoryMap, ok := conjunctionMemoryMap[currentModule.Label]; ok {
					for _, memorizedPulse := range memoryMap {
						if !memorizedPulse {
							sendingPulse = false
							break
						}
					}
				}
				sendPulses(!sendingPulse)
			} else if currentModule.Type == 2 {
				sendPulses(currentPulse)
			} else if currentModule.Type == 3 {
				sendPulses(false)
			}
		}
	}

	return strconv.Itoa(lowPulseCount * highPulseCount)
}
