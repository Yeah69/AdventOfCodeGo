package main

import (
	"fmt"
	"os"
	"testing"
)

func ExecuteTest(label string, expected string, task task, t *testing.T) {
	fileName := fmt.Sprintf("./resources/Day%s.txt", label)
	file, _ := os.ReadFile(fileName)
	input := string(file)
	result := task(input)
	if result != expected {
		t.Fatalf(`Day %s failed. Expected = %s; Result = %s`, label, expected, result)
	}
}

func TestDay01Task0(t *testing.T) {
	ExecuteTest("01", "54561", Day01Task0, t)
}

func TestDay01Task1(t *testing.T) {
	ExecuteTest("01", "54076", Day01Task1, t)
}

func TestDay02Task0(t *testing.T) {
	ExecuteTest("02", "2317", Day02Task0, t)
}

func TestDay02Task1(t *testing.T) {
	ExecuteTest("02", "74804", Day02Task1, t)
}

func TestDay03Task0(t *testing.T) {
	ExecuteTest("03", "536576", Day03Task0, t)
}

func TestDay03Task1(t *testing.T) {
	ExecuteTest("03", "75741499", Day03Task1, t)
}

func TestDay04Task0(t *testing.T) {
	ExecuteTest("04", "23028", Day04Task0, t)
}

func TestDay04Task1(t *testing.T) {
	ExecuteTest("04", "9236992", Day04Task1, t)
}

func TestDay05Task0(t *testing.T) {
	ExecuteTest("05", "282277027", Day05Task0, t)
}

func TestDay05Task1(t *testing.T) {
	ExecuteTest("05", "11554135", Day05Task1, t)
}

func TestDay06Task0(t *testing.T) {
	ExecuteTest("06", "211904", Day06Task0, t)
}

func TestDay06Task1(t *testing.T) {
	ExecuteTest("06", "43364472", Day06Task1, t)
}

func TestDay07Task0(t *testing.T) {
	ExecuteTest("07", "253603890", Day07Task0, t)
}

func TestDay07Task1(t *testing.T) {
	ExecuteTest("07", "253630098", Day07Task1, t)
}

func TestDay08Task0(t *testing.T) {
	ExecuteTest("08", "16343", Day08Task0, t)
}

func TestDay08Task1(t *testing.T) {
	ExecuteTest("08", "15299095336639", Day08Task1, t)
}

func TestDay09Task0(t *testing.T) {
	ExecuteTest("09", "1955513104", Day09Task0, t)
}

func TestDay09Task1(t *testing.T) {
	ExecuteTest("09", "1131", Day09Task1, t)
}

func TestDay10Task0(t *testing.T) {
	ExecuteTest("10", "6800", Day10Task0, t)
}

func TestDay10Task1(t *testing.T) {
	ExecuteTest("10", "483", Day10Task1, t)
}

func TestDay11Task0(t *testing.T) {
	ExecuteTest("11", "10228230", Day11Task0, t)
}

func TestDay11Task1(t *testing.T) {
	ExecuteTest("11", "447073334102", Day11Task1, t)
}

func TestDay12Task0(t *testing.T) {
	ExecuteTest("12", "7361", Day12Task0, t)
}

func TestDay12Task1(t *testing.T) {
	ExecuteTest("12", "83317216247365", Day12Task1, t)
}

func TestDay13Task0(t *testing.T) {
	ExecuteTest("13", "34772", Day13Task0, t)
}

func TestDay13Task1(t *testing.T) {
	ExecuteTest("13", "35554", Day13Task1, t)
}

func TestDay14Task0(t *testing.T) {
	ExecuteTest("14", NoSolutionFound, Day14Task0, t)
}

func TestDay14Task1(t *testing.T) {
	ExecuteTest("14", NoSolutionFound, Day14Task1, t)
}

func TestDay15Task0(t *testing.T) {
	ExecuteTest("15", NoSolutionFound, Day15Task0, t)
}

func TestDay15Task1(t *testing.T) {
	ExecuteTest("15", NoSolutionFound, Day15Task1, t)
}

func TestDay16Task0(t *testing.T) {
	ExecuteTest("16", NoSolutionFound, Day16Task0, t)
}

func TestDay16Task1(t *testing.T) {
	ExecuteTest("16", NoSolutionFound, Day16Task1, t)
}

func TestDay17Task0(t *testing.T) {
	ExecuteTest("17", NoSolutionFound, Day17Task0, t)
}

func TestDay17Task1(t *testing.T) {
	ExecuteTest("17", NoSolutionFound, Day17Task1, t)
}

func TestDay18Task0(t *testing.T) {
	ExecuteTest("18", NoSolutionFound, Day18Task0, t)
}

func TestDay18Task1(t *testing.T) {
	ExecuteTest("18", NoSolutionFound, Day18Task1, t)
}

func TestDay19Task0(t *testing.T) {
	ExecuteTest("19", NoSolutionFound, Day19Task0, t)
}

func TestDay19Task1(t *testing.T) {
	ExecuteTest("19", NoSolutionFound, Day19Task1, t)
}

func TestDay20Task0(t *testing.T) {
	ExecuteTest("20", NoSolutionFound, Day20Task0, t)
}

func TestDay20Task1(t *testing.T) {
	ExecuteTest("20", NoSolutionFound, Day20Task1, t)
}

func TestDay21Task0(t *testing.T) {
	ExecuteTest("21", NoSolutionFound, Day21Task0, t)
}

func TestDay21Task1(t *testing.T) {
	ExecuteTest("21", NoSolutionFound, Day21Task1, t)
}

func TestDay22Task0(t *testing.T) {
	ExecuteTest("22", NoSolutionFound, Day22Task0, t)
}

func TestDay22Task1(t *testing.T) {
	ExecuteTest("22", NoSolutionFound, Day22Task1, t)
}

func TestDay23Task0(t *testing.T) {
	ExecuteTest("23", NoSolutionFound, Day23Task0, t)
}

func TestDay23Task1(t *testing.T) {
	ExecuteTest("23", NoSolutionFound, Day23Task1, t)
}

func TestDay24Task0(t *testing.T) {
	ExecuteTest("24", NoSolutionFound, Day24Task0, t)
}

func TestDay24Task1(t *testing.T) {
	ExecuteTest("24", NoSolutionFound, Day24Task1, t)
}

func TestDay25Task0(t *testing.T) {
	ExecuteTest("25", NoSolutionFound, Day25Task0, t)
}

func TestDay25Task1(t *testing.T) {
	ExecuteTest("25", NothingToDoHere, Day25Task1, t)
}
