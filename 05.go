package main

import (
	. "aoc/util"
	"os"
	"slices"
)

var rulem = make(Set[[2]int])

func main() {
	lines := Input(os.Args[1], "\n\n", true)

	for _, x := range Spac(lines[0], "\n", -1) {
		v := Vatoi(Spac(x, "|", -1))
		rulem[[2]int{v[0], v[1]}] = true
	}

	part1 := 0
	part2 := 0
	for _, line := range Spac(lines[1], "\n", -1) {
		cand := Vatoi(Spac(line, ",", -1))
		cand2 := append([]int{}, cand...)
		slices.SortStableFunc(cand2, func(a, b int) int {
			if rulem[[2]int{a, b}] {
				return -1
			}
			if rulem[[2]int{b, a}] {
				return 1
			}
			return 0
		})
		if slices.Equal(cand, cand2) {
			part1 += cand[len(cand)/2]
		} else {
			part2 += cand2[len(cand2)/2]
		}
	}
	Sol(part1)
	Sol(part2)
}
