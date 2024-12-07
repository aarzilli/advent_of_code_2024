package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	part1 := 0
	for _, line := range lines {
		v := Getints(line, false)
		if enum(v[0], v[1], v[2:]) {
			pln(v)
			part1 += v[0]
		}
	}
	Sol(part1)

	part2 := 0
	for _, line := range lines {
		v := Getints(line, false)
		if enum2(v[0], v[1], v[2:]) {
			pln(v)
			part2 += v[0]
		}
	}
	Sol(part2)
}

func enum(out int, cur int, rest []int) bool {
	if len(rest) == 0 {
		return cur == out
	}
	if enum(out, cur+rest[0], rest[1:]) {
		return true
	}
	return enum(out, cur*rest[0], rest[1:])
}

func enum2(out int, cur int, rest []int) bool {
	if len(rest) == 0 {
		return cur == out
	}
	if enum2(out, cur+rest[0], rest[1:]) {
		return true
	}
	if enum2(out, cur*rest[0], rest[1:]) {
		return true
	}
	return enum2(out, concat(cur, rest[0]), rest[1:])
}

func concat(a, b int) int {
	return Atoi(fmt.Sprintf("%d%d", a, b))
}
