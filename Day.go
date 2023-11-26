package main

import (
	"fmt"
	"time"
)

type task func(string) string

func executeWholeDay(label string, input string, taskZero task, taskOne task) {
	fmt.Println()
	fmt.Printf("Day %s", label)
	fmt.Println()
	fmt.Println()
	fmt.Println("Input =")
	fmt.Println(input)
	fmt.Println()

	executeTask("0", input, taskZero)
	executeTask("1", input, taskOne)
}

func executeTask(taskLabel string, input string, task task) {
	defer timeTrack(time.Now(), "(The task took %s)")
	fmt.Printf("Answer %s =", taskLabel)
	fmt.Println()
	fmt.Println(task(input))
	fmt.Println()
}

func timeTrack(start time.Time, format string) {
	elapsed := time.Since(start)
	fmt.Printf(format, elapsed)
	fmt.Println()
	fmt.Println()
}
