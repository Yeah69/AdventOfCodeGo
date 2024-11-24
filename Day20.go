package main

import (
	"fmt"
	"slices"
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

	rxTrigger := ""
	for current := range modulesMap {
		module := modulesMap[current]
		if slices.Contains(module.Next, "rx") {
			rxTrigger = current
		}
	}

	rxTriggerTriggers := make([]string, 0)
	for current := range modulesMap {
		module := modulesMap[current]
		if slices.Contains(module.Next, rxTrigger) {
			rxTriggerTriggers = append(rxTriggerTriggers, current)
		}
	}

	triggerTriggersCycleCounts := make(map[string]int)

	for i := range rxTriggerTriggers {
		triggerTriggersCycleCounts[rxTriggerTriggers[i]] = -1
	}

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

	i := 0
	for {
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
				if currentModule.Label == rxTrigger {
					if rxTriggerConjunction, ok := conjunctionMemoryMap[rxTrigger]; ok {
						for rxTriggerTrigger := range rxTriggerConjunction {
							rxTriggerTriggerPulse := rxTriggerConjunction[rxTriggerTrigger]
							if rxTriggerTriggerPulse && triggerTriggersCycleCounts[rxTriggerTrigger] == -1 {
								triggerTriggersCycleCounts[rxTriggerTrigger] = i + 1
							}
						}
					}
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
		i++
		/*if rxTriggerConjunction, ok := conjunctionMemoryMap[rxTrigger]; ok {
			for rxTriggerTrigger := range rxTriggerConjunction {
				rxTriggerTriggerPulse := rxTriggerConjunction[rxTriggerTrigger]
				if rxTriggerTriggerPulse && triggerTriggersCycleCounts[rxTriggerTrigger] == -1 {
					triggerTriggersCycleCounts[rxTriggerTrigger] = i
				}
			}
		}*/
		abort := true
		for s := range triggerTriggersCycleCounts {
			if triggerTriggersCycleCounts[s] == -1 {
				abort = false
			}
		}
		if abort {
			break
		}
	}

	ret := 1
	for s := range triggerTriggersCycleCounts {
		ret *= triggerTriggersCycleCounts[s]
	}

	connections := make([]string, 0)
	for current := range modulesMap {
		module := modulesMap[current]
		for i := range module.Next {
			next := module.Next[i]
			label := ""
			if module.Type == 0 {
				label = fmt.Sprintf(" [label=\"%s\"]", "%")
			} else if module.Type == 1 {
				label = fmt.Sprintf(" [label=\"%s\"]", "&")
			}
			connections = append(connections, fmt.Sprintf("    %s -> %s%s", current, next, label))
		}
	}

	_ = fmt.Sprintf("digraph {\n%s\n}", strings.Join(connections, "\n"))

	return strconv.Itoa(ret)
}
