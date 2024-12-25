package main

import (
	. "aoc/util"
	"os"
)

var keys, locks [][]int

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	Pf("len %d\n", len(groups))

	for _, group := range groups {
		lines := Noempty(Spac(group, "\n", -1))
		if lines[0] == "#####" {
			locks = append(locks, convert(lines[1:]))
		} else if lines[6] == "#####" {
			keys = append(keys, convert(lines[:6]))
		} else {
			panic("blah")
		}
	}

	part1 := 0
	for _, key := range keys {
		for _, lock := range locks {
			isok := true
			for i := range key {
				if key[i]+lock[i] > 5 {
					isok = false
					break
				}
			}
			if isok {
				part1++
			}
		}
	}
	Sol(part1)
}

func convert(ss []string) []int {
	r := make([]int, len(ss[0]))
	for i := range ss {
		for j := range ss[i] {
			switch ss[i][j] {
			case '#':
				r[j]++
			case '.':
				// ok
			default:
				panic("blah")
			}
		}
	}
	return r
}
