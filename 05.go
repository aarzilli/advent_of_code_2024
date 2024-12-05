package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"slices"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

var rules [][]int
var rulem = make(map[[2]int]int)

func verify(cand []int) bool {
	m := make(map[int]int)
	for i, x := range cand {
		m[x] = i
	}
	for _, rule := range rules {
		i, ok1 := m[rule[0]]
		j, ok2 := m[rule[1]]
		if !ok1 || !ok2 {
			continue
		}
		if i > j {
			return false
		}
	}
	return true
}

func reorder(cand []int) {
	slices.SortStableFunc(cand, func(a, b int) int {
		if rulem[[2]int{a, b}] != 0 {
			return -1
		}
		if rulem[[2]int{b, a}] != 0 {
			return 1
		}
		return 0
	})
}

func main() {
	lines := Input(os.Args[1], "\n\n", true)
	pf("len %d\n", len(lines))

	for _, x := range Spac(lines[0], "\n", -1) {
		rules = append(rules, Vatoi(Spac(x, "|", -1)))
	}

	pln(rules)

	for _, rule := range rules {
		rulem[[2]int{rule[0], rule[1]}] = -1
	}

	part1 := 0
	part2 := 0
	for _, line := range Spac(lines[1], "\n", -1) {
		cand := Vatoi(Spac(line, ",", -1))
		ok := verify(cand)
		pln(cand, ok)
		if ok {
			part1 += cand[len(cand)/2]
		} else {
			cand2 := make([]int, len(cand))
			copy(cand2, cand)
			reorder(cand2)
			pln(cand, "reorders to", cand2)
			part2 += cand2[len(cand2)/2]
		}
	}
	Sol(part1)
	Sol(part2)
}
