package main

import (
	"strconv"
	"strings"
)

type Day24Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Day24Hailstone struct {
	Pos Day24Vec3
	Vel Day24Vec3
}

func Day24ParseInput(input string) []Day24Hailstone {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "@", ",")
	lines := strings.Split(input, "\n")

	hailstones := make([]Day24Hailstone, len(lines))

	for i := range lines {
		textParts := strings.Split(lines[i], ",")
		posX, _ := strconv.Atoi(textParts[0])
		posY, _ := strconv.Atoi(textParts[1])
		posZ, _ := strconv.Atoi(textParts[2])
		velX, _ := strconv.Atoi(textParts[3])
		velY, _ := strconv.Atoi(textParts[4])
		velZ, _ := strconv.Atoi(textParts[5])
		hailstones[i] = Day24Hailstone{
			Day24Vec3{float64(posX), float64(posY), float64(posZ)},
			Day24Vec3{float64(velX), float64(velY), float64(velZ)},
		}
	}

	return hailstones
}

func Day24Task0(input string) string {
	hailstones := Day24ParseInput(input)

	rMin := 200000000000000.0
	rMax := 400000000000000.0

	count := 0

	lineFunctions := make([]Pair[float64, float64], 0)

	for i := range hailstones {
		hailstone := hailstones[i]
		a := float64(hailstone.Vel.Y) / float64(hailstone.Vel.X)
		b := float64(hailstone.Pos.Y) - a*float64(hailstone.Pos.X)
		lineFunctions = append(lineFunctions, Pair[float64, float64]{a, b})
	}

	for a := 0; a < len(lineFunctions); a++ {
		lineA := lineFunctions[a]
		hailstoneA := hailstones[a]
		for b := a + 1; b < len(lineFunctions); b++ {
			lineB := lineFunctions[b]
			hailstoneB := hailstones[b]

			if lineA.First == lineB.First { // parallel
				continue
			}

			x := (lineB.Second - lineA.Second) / (lineA.First - lineB.First)
			y := lineA.First*x + lineA.Second

			stepsA := (x - float64(hailstoneA.Pos.X)) / float64(hailstoneA.Vel.X)
			stepsB := (x - float64(hailstoneB.Pos.X)) / float64(hailstoneB.Vel.X)

			if x >= rMin && x <= rMax && y >= rMin && y <= rMax && stepsA >= 0 && stepsB >= 0 {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func Day24Task1(input string) string {
	hailstones := Day24ParseInput(input)

	add := func(left Day24Vec3, right Day24Vec3) Day24Vec3 {
		return Day24Vec3{left.X + right.X, left.Y + right.Y, left.Z + right.Z}
	}

	minus := func(left Day24Vec3, right Day24Vec3) Day24Vec3 {
		return Day24Vec3{left.X - right.X, left.Y - right.Y, left.Z - right.Z}
	}

	cross := func(left Day24Vec3, right Day24Vec3) Day24Vec3 {
		return Day24Vec3{
			left.Y*right.Z - left.Z*right.Y,
			left.Z*right.X - left.X*right.Z,
			left.X*right.Y - left.Y*right.X}
	}

	dot := func(left Day24Vec3, right Day24Vec3) float64 {
		return left.X*right.X + left.Y*right.Y + left.Z*right.Z
	}

	scalar := func(left float64, right Day24Vec3) Day24Vec3 {
		return Day24Vec3{left * right.X, left * right.Y, left * right.Z}
	}

	scalarDiv := func(left Day24Vec3, right float64) Day24Vec3 {
		return Day24Vec3{left.X / right, left.Y / right, left.Z / right}
	}

	hailstone0 := hailstones[0]
	hailstone1 := hailstones[10]
	hailstone2 := hailstones[2]

	p1 := minus(hailstone1.Pos, hailstone0.Pos)
	v1 := minus(hailstone1.Vel, hailstone0.Vel)
	p2 := minus(hailstone2.Pos, hailstone0.Pos)
	v2 := minus(hailstone2.Vel, hailstone0.Vel)

	t1 := -dot(cross(p1, p2), v2) / dot(cross(v1, p2), v2)
	t2 := -dot(cross(p1, p2), v1) / dot(cross(p1, v2), v1)

	c1 := add(hailstone1.Pos, scalar(t1, hailstone1.Vel))
	c2 := add(hailstone2.Pos, scalar(t2, hailstone2.Vel))
	v := scalarDiv(minus(c2, c1), t2-t1)
	p := minus(c1, scalar(t1, v))

	return strconv.Itoa(int(p.X + p.Y + p.Z))
}
